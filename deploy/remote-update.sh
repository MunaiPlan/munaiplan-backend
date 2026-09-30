#!/bin/sh
# Release from a developer machine: push your commits first, then run from the workspace root:
#   sh munaiplan-backend/deploy/remote-update.sh [git-ref]
# Checks that all three repositories are pushed, runs update.sh on the server and smoke-checks
# the public site.
# Environment: DEPLOY_HOST (default 159.89.10.247), DEPLOY_KEY (default ~/.ssh/munaiplan_do),
# SITE (default munaiplan.com).
set -eu
main() {
  ref=${1:-recovery/phase-2}
  host=${DEPLOY_HOST:-159.89.10.247}
  key=${DEPLOY_KEY:-$HOME/.ssh/munaiplan_do}
  site=${SITE:-munaiplan.com}
  ssh_() { ssh -i "$key" -o BatchMode=yes "root@$host" "$@"; }

  for repo in munaiplan-backend munaiplan-frontend munai-models; do
    git -C "$repo" fetch --quiet origin
    [ "$(git -C "$repo" rev-parse "$ref")" = "$(git -C "$repo" rev-parse "origin/$ref")" ] \
      || { echo "$repo: $ref is not pushed (or is behind origin); push first" >&2; exit 1; }
  done
  # Refresh the deploy scripts themselves first, so the server runs this release's update.sh.
  ssh_ "git -C /opt/munaiplan/src/munaiplan-backend fetch --quiet origin && git -C /opt/munaiplan/src/munaiplan-backend checkout --quiet --detach 'origin/$ref' 2>/dev/null || git -C /opt/munaiplan/src/munaiplan-backend checkout --quiet --detach '$ref'"
  ssh_ "sh /opt/munaiplan/src/munaiplan-backend/deploy/update.sh '$ref'"

  fail=0
  check() { code=$(curl -s -o /dev/null -w '%{http_code}' "$2" || true); echo "$1 $code"; [ "$code" = "$3" ] || fail=1; }
  check "health          " "https://$site/health" 200
  check "spa deep link   " "https://$site/cases/smoke" 200
  check "user manual     " "https://$site/docs" 200
  check "anonymous api   " "https://$site/api/v1/companies/" 401
  check "public sign-up  " "https://$site/api/v1/auth/sign-up" 404
  [ $fail = 0 ] && echo "Release OK: https://$site" || { echo "smoke checks failed" >&2; exit 1; }
}
main "$@"
