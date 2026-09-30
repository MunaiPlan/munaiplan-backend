# Deploying MunaiPlan

There are two modes:
- **public** (preview): Caddy serves HTTPS on a public address, and access is protected by admin-created accounts.
- **tailscale** (private pilot): no public ports at all.

## Small servers (1 GB, e.g. DigitalOcean $6 droplet): prebuilt images

A 1 GB server can run the stack, which uses about 0.5 GB, but cannot build the images. Build them elsewhere for the server's architecture and load them there:

```sh
# On a build machine (from the workspace root):
docker buildx build --platform linux/amd64 -t munaiplan-api:prod --load munaiplan-backend
docker buildx build --platform linux/amd64 -t munaiplan-frontend:prod -f munaiplan-frontend/dockerfile --load munaiplan-frontend
docker buildx build --platform linux/amd64 -t munaiplan-model:prod --load munai-models
docker save munaiplan-api:prod munaiplan-frontend:prod munaiplan-model:prod | gzip | ssh root@<ip> 'gunzip | docker load'
# On the server: add PREBUILT_IMAGES=1 to /opt/munaiplan/.env, then run deploy.sh as usual.
```

`bootstrap-server.sh` adds 2 GB of swap for TensorFlow's start-up peak.

## Oracle Cloud Always Free (alternative)

- Use an **Ampere A1** instance: `VM.Standard.A1.Flex`, 2–4 OCPU and 12–24 GB RAM, Ubuntu 24.04. This stays within the Always Free limits and runs arm64 natively.
- **Log in as `ubuntu`, not root:** `ssh ubuntu@<ip>`. Run the scripts below with `sudo`.
- **Open ports 80 and 443 in two places:**
  - the VCN security list: Networking → VCN → Security Lists → Ingress, TCP 80 and 443 from `0.0.0.0/0`;
  - the host firewall, which `bootstrap-server.sh` handles automatically (Oracle images ship their own iptables rules).
- Let's Encrypt HTTPS works on `<ip-with-dashes>.sslip.io`.

## Public preview in short

```sh
DEPLOY_MODE=public sh deploy/bootstrap-server.sh
sh deploy/init-env.sh admin@your-company.kz 203-0-113-7.sslip.io   # <ip-with-dashes>.sslip.io needs no domain
sh deploy/deploy.sh recovery/phase-2                                 # prints https://203-0-113-7.sslip.io
sh deploy/install-backup-timer.sh
```

The Let's Encrypt certificate is issued on first start; ports 80 and 443 must be reachable. Anyone with the link sees the sign-in page. Only accounts created in **Администрирование** can sign in.

The sections below describe the private (Tailscale) mode.


One Ubuntu 24.04 server runs the same four containers as local development: Postgres, API, ML model and frontend. **No web port is public.** Tailscale publishes the frontend to your private tailnet over HTTPS, so only invited devices can open it.

Rehearsed locally on 30 September 2026:
- build → migrate (4 migrations) → four healthy services;
- `create-admin` run twice creates one admin;
- admin sign-in works; `/api/v1/status` reports the model ready;
- the public sign-up route returns 404, and `dev-seed` is refused in production;
- `pg_dump` → `pg_restore` into a fresh database gives identical table, user and migration counts.

## Requirements

- A VPS: **4 GB RAM, 2 vCPU, 60 GB SSD** or more, Ubuntu 24.04, SSH-key login. (A Kazakhstan provider was chosen so the data stays in Kazakhstan.)
- A free [Tailscale](https://tailscale.com) account. Every pilot user installs the Tailscale app and is invited to the tailnet.
- The three repositories on GitHub under `MunaiPlan`: `munaiplan-backend`, `munaiplan-frontend`, `munai-models`. If any is private, give the server a read-only deploy key.

## First deployment

```sh
# On the server, as root:
git clone https://github.com/MunaiPlan/munaiplan-backend.git /opt/munaiplan/src/munaiplan-backend
cd /opt/munaiplan/src/munaiplan-backend && git checkout recovery/phase-2
sh deploy/bootstrap-server.sh          # Docker, Tailscale, firewall, security updates, swap
tailscale up                           # open the printed link and log in to your Tailscale account
sh deploy/init-env.sh admin@your-company.kz   # generates /opt/munaiplan/.env (0600)
sh deploy/deploy.sh recovery/phase-2   # clones, builds, migrates, starts, creates the admin
sh deploy/install-backup-timer.sh      # nightly database backup at 02:30
```

- `deploy.sh` ends by printing the private HTTPS address, for example `https://munaiplan.tailXXXX.ts.net`.
- **Admin password:** read `ADMIN_PASSWORD` from `/opt/munaiplan/.env` yourself over SSH. Sign in and create organizations and users in **Администрирование**, then change nothing else in `.env`: the database password, salt and token secrets must stay stable.
- **Harden SSH** once Tailscale works: `ufw delete allow OpenSSH && ufw allow in on tailscale0 to any port 22`. After that, SSH works only over the tailnet (`ssh root@<tailscale-name>`).

## Updates

```sh
sh /opt/munaiplan/src/munaiplan-backend/deploy/deploy.sh <branch-or-tag>
```

The deploy is safe to repeat. Migrations apply once and existing data is kept. Take a backup first when a release contains new migrations: `sh deploy/backup.sh`.

## Backups and restore

- Dumps go to `/opt/munaiplan/backups/munaiplan-<UTC>.dump` and are kept for 14 days.
- **Copy them off the server regularly**, for example with `scp` from your machine. A backup stored on the same disk does not survive losing the server.

To restore:

```sh
cd /opt/munaiplan/src/munaiplan-backend
docker compose -p munaiplan --env-file /opt/munaiplan/.env -f deploy/compose.prod.yaml stop app frontend
docker exec -i munaiplan-postgres-1 pg_restore -U munaiplan -d munaiplan --clean --if-exists --no-owner < /opt/munaiplan/backups/<file>.dump
docker compose -p munaiplan --env-file /opt/munaiplan/.env -f deploy/compose.prod.yaml up -d app frontend
```

## Operations

```sh
C="docker compose -p munaiplan --env-file /opt/munaiplan/.env -f /opt/munaiplan/src/munaiplan-backend/deploy/compose.prod.yaml"
$C ps              # health
$C logs -f app     # API logs (also: model, frontend, postgres)
$C restart model
```

- The in-app status bar shows API and model health to every user.
- Access tokens last 8 hours; there is no refresh flow yet, so users sign in again after that.

## Before any wider rollout

- Predictions are **not validated**; see `docs/recovery/model-validation.md` in the workspace.
- Decide the pending units questions: yield in ksi, and missing coordinates.
- Add off-server backup copies, uptime monitoring and a documented rollback (redeploy the previous tag).
