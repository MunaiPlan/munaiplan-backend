# Deploying MunaiPlan

This is the single runbook for releases, CI/CD, first installs and operations. The server facts for the live preview, such as its address and accounts, are in the workspace handoff (`docs/recovery/handoff/START_HERE.md` §6).

## The live preview

| | |
|---|---|
| URL | **https://munaiplan.com**; `www.munaiplan.com` and `https://159-89-10-247.sslip.io` serve the same app |
| Server | DigitalOcean $6 droplet: 1 vCPU, 1 GB RAM plus 2 GB swap, FRA1, Ubuntu 24.04, `root@159.89.10.247` |
| SSH | Key `~/.ssh/munaiplan_do` on the owner's Mac |
| DNS | GoDaddy: `A @ → 159.89.10.247`, `CNAME www → munaiplan.com` |
| Mode | `DEPLOY_MODE=public` (Caddy with Let's Encrypt), `PREBUILT_IMAGES=1` (images built on the server one at a time) |
| Branch | `recovery/phase-2` in all three repositories |

## How a release flows

```
push to GitHub ──► CI (GitHub Actions: tests, lint, image build)
      │
      └─► release: remote-update.sh (workstation)  or  Actions → Deploy (manual)
                 └─► server: update.sh
                        1. fetch backend, frontend, models (origin/<ref>)
                        2. backup.sh        (pg_dump into /opt/munaiplan/backups)
                        3. build-images.sh  (sequential docker builds, about 10 min)
                        4. deploy.sh        (migrate, start, create-admin, wait for health)
                 └─► smoke checks: /health 200, SPA deep link 200, anonymous API 401, sign-up 404
```

A release is manual on purpose. CI runs on every push, but nothing deploys automatically: this is a single server with real data, and every release takes a backup first.

### Release from a workstation (recommended)

From the workspace root, after pushing your commits:
```sh
sh munaiplan-backend/deploy/remote-update.sh                 # branch recovery/phase-2
sh munaiplan-backend/deploy/remote-update.sh v0.3.0          # or a tag present in all three repositories
```

The script:
- refuses to run if the backend or frontend `ref` is not pushed;
- checks that `munai-models` is pushed too;
- runs `update.sh` on the server and prints the smoke-check results.

It takes about 12 minutes, most of it the model image. The site stays up during the build. There is a short restart only when the containers are replaced.

### Release from GitHub Actions

`munaiplan-backend` → Actions → **Deploy** → Run workflow (input: `ref`). It needs three repository secrets, which the owner adds:
- `DEPLOY_HOST`;
- `DEPLOY_SSH_KEY`: a **dedicated** key; add its public half to the server's `/root/.ssh/authorized_keys`;
- `DEPLOY_KNOWN_HOSTS`: the output of `ssh-keyscan -t ed25519 159.89.10.247`.

It also needs a GitHub environment named `preview`, where a required reviewer can be added.

The deploy covers all three repositories, `munai-models` included.

### CI

| Repository | Workflow | Checks |
|---|---|---|
| munaiplan-backend | `.github/workflows/ci.yml` | `go vet`, `go test` (real-report tests are skipped without `WELLPLAN_REPORTS_DIR`), deploy scripts parse, production image builds |
| munaiplan-frontend | `.github/workflows/ci.yml` | `npm ci`, lint (0 warnings), build, unit tests, production image builds |
| munai-models | `.github/workflows/ci.yml` | Runtime image builds, then the unit tests run inside it. The repository is **private** |

**Local note:** Go 1.21.1 test binaries do not start on recent macOS (`dyld: missing LC_UUID`). Run the tests in Linux instead:
```sh
docker run --rm -v "$PWD":/src -w /src golang:1.21.1-alpine3.18 go test ./...
```

### Rollback

Redeploy the previous commit or tag with `remote-update.sh <previous-ref>`. Migrations only move forward. If a release contained a migration, restore the backup that `update.sh` took just before it (see "Backups and restore" below).

## Scripts

| Script | Runs on | Purpose |
|---|---|---|
| `remote-update.sh [ref]` | workstation | Push check (all three repositories), `update.sh`, smoke checks |
| `update.sh [ref]` | server | Fetch, back up, build, deploy |
| `build-images.sh` | server | Sequential image builds; log in `/opt/munaiplan/build.log` |
| `deploy.sh [ref]` | server | Migrate, start, create the admin, wait for health. Builds images itself when `PREBUILT_IMAGES` is not 1 |
| `bootstrap-server.sh` | server, once | Docker, `docker-buildx`, firewall, unattended upgrades, 2 GB swap; Tailscale only in private mode |
| `init-env.sh EMAIL [SITE]` | server, once | Generates `/opt/munaiplan/.env` (mode 0600); a site address selects public mode |
| `backup.sh`, `install-backup-timer.sh` | server | `pg_dump`, and its nightly systemd timer at 02:30 with 14-day retention |

All server scripts keep their body in a `main` function, because they check out new versions of themselves while running.

**Compose files:**
- `compose.prod.yaml`: Postgres, API, model and frontend (bound to `127.0.0.1:8080`), plus `migrate` and `create-admin` jobs;
- `compose.public.yaml`: Caddy on ports 80 and 443; see `Caddyfile`.

## `/opt/munaiplan/.env` (server only, never printed or committed)

| Key | Notes |
|---|---|
| `DB_PASSWORD`, `USER_ACCESS_TOKEN_SECRET`, `USER_REFRESH_TOKEN_SECRET` | Generated. Changing the token secrets signs everyone out |
| `PASSWORD_SALT` | **Never change it.** Every stored password hash depends on it. The live server uses the same salt as the owner's local stack, because local data was copied there |
| `ADMIN_EMAIL`, `ADMIN_PASSWORD` | The bootstrap admin, created by `create-admin` if missing. The owner reads the password over SSH themselves |
| `DEPLOY_MODE`, `SITE_ADDRESS` | `public`; a comma-separated list of hostnames, the first being the canonical one |
| `PREBUILT_IMAGES` | `1` on small servers: `update.sh` builds the images and `deploy.sh` does not |
| `LOCAL_ONLY_REPOS` | Optional. Repositories `deploy.sh` must not fetch |

## Custom domain

1. At the registrar (GoDaddy: My Products → DNS), set `A @ → <ip>` and `CNAME www → <domain>`.
2. List every hostname: `SITE_ADDRESS=munaiplan.com, www.munaiplan.com, 159-89-10-247.sslip.io`.
3. Run `deploy.sh` again. Caddy obtains and renews a certificate for each hostname.

## First install on a new server

For a 1 GB server, keep `PREBUILT_IMAGES=1`. With 4 GB or more you can leave it out, and `deploy.sh` builds with compose.

```sh
# From the workspace root (IP=a.b.c.d, SITE=a-b-c-d.sslip.io or a domain pointing at the server):
SSH="ssh -i ~/.ssh/munaiplan_do root@$IP"
$SSH 'git clone -q -b recovery/phase-2 https://github.com/MunaiPlan/munaiplan-backend.git /opt/munaiplan/src/munaiplan-backend'
$SSH 'DEPLOY_MODE=public sh /opt/munaiplan/src/munaiplan-backend/deploy/bootstrap-server.sh'
$SSH 'git clone -q -b recovery/phase-2 https://github.com/MunaiPlan/munaiplan-frontend.git /opt/munaiplan/src/munaiplan-frontend'
# munai-models is private: give the server a read-only deploy key (the live server's is
# /root/.ssh/munai_models_deploy, registered in GitHub → munai-models → Settings → Deploy keys).
$SSH 'ssh-keygen -q -t ed25519 -N "" -f /root/.ssh/munai_models_deploy && cat /root/.ssh/munai_models_deploy.pub'   # add as a read-only deploy key
$SSH 'ssh-keyscan -t ed25519 github.com >> /root/.ssh/known_hosts && printf "Host github-munai-models\n  HostName github.com\n  User git\n  IdentityFile /root/.ssh/munai_models_deploy\n  IdentitiesOnly yes\n" >> /root/.ssh/config'
$SSH 'git clone -q -b recovery/phase-2 git@github-munai-models:MunaiPlan/munai-models.git /opt/munaiplan/src/munai-models'
$SSH "sh /opt/munaiplan/src/munaiplan-backend/deploy/init-env.sh admin@your-company.kz $SITE && echo PREBUILT_IMAGES=1 >> /opt/munaiplan/.env"
$SSH 'sh /opt/munaiplan/src/munaiplan-backend/deploy/update.sh && sh /opt/munaiplan/src/munaiplan-backend/deploy/install-backup-timer.sh'
```

**Private pilot instead of public:**
- Omit the site address in `init-env.sh`. Bootstrap installs Tailscale; run `tailscale up`.
- `deploy.sh` publishes the frontend to the tailnet only.
- Harden SSH afterwards: `ufw delete allow OpenSSH && ufw allow in on tailscale0 to any port 22`.

**Oracle Cloud Always Free (Ampere A1):**
- Log in as `ubuntu` and use `sudo`.
- Open ports 80 and 443 in the VCN security list. The bootstrap script handles Oracle's host iptables.

## Operations

```sh
C="docker compose -p munaiplan --env-file /opt/munaiplan/.env -f /opt/munaiplan/src/munaiplan-backend/deploy/compose.prod.yaml -f /opt/munaiplan/src/munaiplan-backend/deploy/compose.public.yaml"
$C ps                 # health of every service
$C logs -f app        # also: model, frontend, caddy, postgres
$C restart model
free -m; docker stats --no-stream
```

- The in-app status bar shows API and model health.
- Access tokens last 8 hours; there is no refresh flow yet.

## Backups and restore

- `update.sh` and the nightly timer write `/opt/munaiplan/backups/munaiplan-<UTC>.dump`, kept for 14 days.
- **They stay on the same disk.** Copy them off regularly, for example: `scp -i ~/.ssh/munaiplan_do 'root@159.89.10.247:/opt/munaiplan/backups/*.dump' ~/munaiplan-backups/`.

To restore:
```sh
$C stop app frontend
docker exec -i munaiplan-postgres-1 pg_restore -U munaiplan -d munaiplan --clean --if-exists --no-owner < /opt/munaiplan/backups/<file>.dump
$C up -d app frontend
```

## Known limits

- Predictions are **not validated**; see `docs/recovery/model-validation.md`.
- There is no uptime monitoring and no off-server backup copy yet.
- The server is in Frankfurt; the owner accepted this for the preview on 30 September 2026.
