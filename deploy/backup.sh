#!/bin/sh
# Nightly logical backup of the database; keeps 14 days. Installed as a systemd timer by
# install-backup-timer.sh. Restore: see deploy/README.md.
set -eu
dir=/opt/munaiplan/backups
stamp=$(date -u +%Y%m%dT%H%M%SZ)
container=$(docker compose -p munaiplan ps -q postgres)
[ -n "$container" ] || { echo "postgres is not running" >&2; exit 1; }
umask 077
docker exec "$container" pg_dump -U munaiplan -d munaiplan -Fc > "$dir/munaiplan-$stamp.dump.partial"
mv "$dir/munaiplan-$stamp.dump.partial" "$dir/munaiplan-$stamp.dump"
find "$dir" -name 'munaiplan-*.dump' -mtime +14 -delete
echo "backup $dir/munaiplan-$stamp.dump ($(du -h "$dir/munaiplan-$stamp.dump" | cut -f1))"
