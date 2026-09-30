#!/bin/sh
# Builds and (re)starts the stack from the given git ref of all three repositories.
# Usage (as root on the server): deploy.sh [git-ref]   default: recovery/phase-2
set -eu
ref=${1:-recovery/phase-2}
src=/opt/munaiplan/src
env=/opt/munaiplan/.env
[ -r "$env" ] || { echo "missing $env; run deploy/init-env.sh first" >&2; exit 1; }

for repo in munaiplan-backend munaiplan-frontend munai-models; do
  if [ ! -d "$src/$repo/.git" ]; then
    git clone --quiet "https://github.com/MunaiPlan/$repo.git" "$src/$repo"
  fi
  git -C "$src/$repo" fetch --quiet --prune origin
  git -C "$src/$repo" checkout --quiet --detach "origin/$ref" 2>/dev/null || git -C "$src/$repo" checkout --quiet --detach "$ref"
  echo "$repo @ $(git -C "$src/$repo" rev-parse --short HEAD)"
done

compose() { docker compose -p munaiplan --env-file "$env" -f "$src/munaiplan-backend/deploy/compose.prod.yaml" "$@"; }
compose config --quiet
compose build
compose up -d postgres
compose run --rm migrate
compose up -d app model frontend
compose --profile tools run --rm create-admin

echo "Waiting for services to become healthy (the model needs a few minutes)…"
i=0
while [ $i -lt 90 ]; do
  unhealthy=$(compose ps --format '{{.Service}} {{.Health}}' | awk '$2 != "healthy" && $2 != ""' | wc -l)
  [ "$unhealthy" -eq 0 ] && break
  i=$((i + 1)); sleep 5
done
compose ps

# Publish the frontend to the tailnet over HTTPS (idempotent).
tailscale serve --bg --https=443 http://127.0.0.1:8080 >/dev/null
tailscale serve status
