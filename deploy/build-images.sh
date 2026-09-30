#!/bin/sh
# Builds the three production images on the server, one at a time (a 1 GB server runs out of
# memory building them in parallel). Used with PREBUILT_IMAGES=1. Log: /opt/munaiplan/build.log
# Usage (as root on the server): build-images.sh
set -eu
main() {
  src=/opt/munaiplan/src
  log=/opt/munaiplan/build.log
  : > "$log"
  build() {
    echo "== $1 $(date -u +%T)" | tee -a "$log"
    if ! docker build -t "$1" -f "$2/$3" "$2" >> "$log" 2>&1; then
      echo "build of $1 failed; see $log" >&2
      tail -20 "$log" >&2
      exit 1
    fi
  }
  build munaiplan-api:prod "$src/munaiplan-backend" Dockerfile
  build munaiplan-frontend:prod "$src/munaiplan-frontend" dockerfile
  build munaiplan-model:prod "$src/munai-models" Dockerfile   # about 8 min on 1 vCPU
  echo "== images built $(date -u +%T)" | tee -a "$log"
  docker image prune -f >/dev/null   # drop the untagged previous images
}
main "$@"
