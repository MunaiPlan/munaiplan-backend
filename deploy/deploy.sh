#!/bin/sh
# Builds and (re)starts the stack from the given git ref of all three repositories.
# Usage (as root on the server): deploy.sh [git-ref]   default: recovery/phase-2
# DEPLOY_MODE (in /opt/munaiplan/.env): "tailscale" (default, private) or "public" (Caddy HTTPS).
# LOCAL_ONLY_REPOS: space-separated repos already present in /opt/munaiplan/src that must not be
# fetched from GitHub (e.g. cloned from a git bundle).
set -eu
ref=${1:-recovery/phase-2}
src=/opt/munaiplan/src
env=/opt/munaiplan/.env
[ -r "$env" ] || { echo "missing $env; run deploy/init-env.sh first" >&2; exit 1; }
mode=$(sed -n 's/^DEPLOY_MODE=//p' "$env"); mode=${mode:-tailscale}
local_only=$(sed -n 's/^LOCAL_ONLY_REPOS=//p' "$env")
# PREBUILT_IMAGES=1: images were built elsewhere and loaded with `docker load` (small servers
# cannot build TensorFlow/Node/Go images); only this repository is needed for compose files.
prebuilt=$(sed -n 's/^PREBUILT_IMAGES=//p' "$env")
repos="munaiplan-backend munaiplan-frontend munai-models"
[ "$prebuilt" = 1 ] && repos="munaiplan-backend"

for repo in $repos; do
  case " $local_only " in *" $repo "*)
    git -C "$src/$repo" checkout --quiet --detach "$ref"
    echo "$repo @ $(git -C "$src/$repo" rev-parse --short HEAD) (local)"; continue ;;
  esac
  if [ ! -d "$src/$repo/.git" ]; then
    git clone --quiet "https://github.com/MunaiPlan/$repo.git" "$src/$repo"
  fi
  git -C "$src/$repo" fetch --quiet --prune origin
  git -C "$src/$repo" checkout --quiet --detach "origin/$ref" 2>/dev/null || git -C "$src/$repo" checkout --quiet --detach "$ref"
  echo "$repo @ $(git -C "$src/$repo" rev-parse --short HEAD)"
done

files="-f $src/munaiplan-backend/deploy/compose.prod.yaml"
[ "$mode" = public ] && files="$files -f $src/munaiplan-backend/deploy/compose.public.yaml"
# shellcheck disable=SC2086
compose() { docker compose -p munaiplan --env-file "$env" $files "$@"; }
compose config --quiet
nobuild=""
if [ "$prebuilt" = 1 ]; then
  nobuild="--no-build"
  for image in munaiplan-api:prod munaiplan-frontend:prod munaiplan-model:prod; do
    docker image inspect "$image" >/dev/null 2>&1 || { echo "missing prebuilt image $image (docker load it first)" >&2; exit 1; }
  done
else
  compose build
fi
compose up -d $nobuild postgres
compose run --rm $nobuild migrate
compose up -d $nobuild app model frontend
[ "$mode" = public ] && compose up -d caddy
compose --profile tools run --rm $nobuild create-admin

echo "Waiting for services to become healthy (the model needs a few minutes)…"
i=0
while [ $i -lt 90 ]; do
  unhealthy=$(compose ps --format '{{.Service}} {{.Health}}' | awk '$2 != "healthy" && $2 != ""' | wc -l)
  [ "$unhealthy" -eq 0 ] && break
  i=$((i + 1)); sleep 5
done
compose ps

if [ "$mode" = public ]; then
  echo "Open: https://$(sed -n 's/^SITE_ADDRESS=//p' "$env")"
else
  # Publish the frontend to the tailnet over HTTPS (idempotent).
  tailscale serve --bg --https=443 http://127.0.0.1:8080 >/dev/null
  tailscale serve status
fi
