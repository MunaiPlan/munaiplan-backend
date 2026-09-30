#!/bin/sh
# One-command release on the server: fetch sources, back up the database, build images (when
# PREBUILT_IMAGES=1) and deploy. Safe to repeat.
# Usage (as root on the server): update.sh [git-ref]   default: recovery/phase-2
#
# Every repository is updated from its `origin` remote. munai-models is private and is fetched
# with a read-only deploy key (ssh host alias github-munai-models, see deploy/README.md).
set -eu
main() {
  ref=${1:-recovery/phase-2}
  src=/opt/munaiplan/src
  here=$src/munaiplan-backend/deploy
  env=/opt/munaiplan/.env
  for repo in munaiplan-backend munaiplan-frontend munai-models; do
    git -C "$src/$repo" fetch --quiet --prune origin
    git -C "$src/$repo" checkout --quiet --detach "origin/$ref" 2>/dev/null \
      || git -C "$src/$repo" checkout --quiet --detach "$ref"
    echo "$repo @ $(git -C "$src/$repo" log -1 --format='%h %s')"
  done
  sh "$here/backup.sh"
  if [ "$(sed -n 's/^PREBUILT_IMAGES=//p' "$env")" = 1 ]; then
    sh "$here/build-images.sh"
  fi
  sh "$here/deploy.sh" "$ref"
}
main "$@"
