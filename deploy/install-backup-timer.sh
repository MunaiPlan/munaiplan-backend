#!/bin/sh
# Installs a systemd timer that runs backup.sh daily at 02:30 server time.
set -eu
here=$(cd "$(dirname "$0")" && pwd)
cat > /etc/systemd/system/munaiplan-backup.service <<UNIT
[Unit]
Description=MunaiPlan database backup
[Service]
Type=oneshot
ExecStart=/bin/sh $here/backup.sh
UNIT
cat > /etc/systemd/system/munaiplan-backup.timer <<UNIT
[Unit]
Description=Daily MunaiPlan database backup
[Timer]
OnCalendar=*-*-* 02:30:00
Persistent=true
[Install]
WantedBy=timers.target
UNIT
systemctl daemon-reload
systemctl enable --now munaiplan-backup.timer
systemctl list-timers munaiplan-backup.timer --no-pager
