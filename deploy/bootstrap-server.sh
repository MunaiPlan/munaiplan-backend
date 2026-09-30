#!/bin/sh
# One-time preparation of a fresh Ubuntu 24.04 server (run as root).
# Installs Docker, Tailscale, a deny-by-default firewall, automatic security updates and swap.
set -eu

[ "$(id -u)" -eq 0 ] || { echo "run as root" >&2; exit 1; }
. /etc/os-release
[ "$ID" = ubuntu ] || { echo "expected Ubuntu, found $ID" >&2; exit 1; }

export DEBIAN_FRONTEND=noninteractive
apt-get update -q
apt-get upgrade -yq
apt-get install -yq ca-certificates curl git ufw unattended-upgrades docker.io docker-compose-v2
systemctl enable --now docker
dpkg-reconfigure -f noninteractive unattended-upgrades

# Tailscale from its signed apt repository.
curl -fsSL "https://pkgs.tailscale.com/stable/ubuntu/${VERSION_CODENAME}.noarmor.gpg" -o /usr/share/keyrings/tailscale-archive-keyring.gpg
curl -fsSL "https://pkgs.tailscale.com/stable/ubuntu/${VERSION_CODENAME}.tailscale-keyring.list" -o /etc/apt/sources.list.d/tailscale.list
apt-get update -q
apt-get install -yq tailscale
systemctl enable --now tailscaled

# Firewall: nothing public except SSH (restrict SSH to the tailnet later, see README).
# Docker-published ports bypass ufw, which is why the stack only binds 127.0.0.1.
ufw default deny incoming
ufw default allow outgoing
ufw allow OpenSSH
ufw allow in on tailscale0
ufw --force enable

# 2 GB swap: the ML image build and TensorFlow start-up peak above 4 GB RAM otherwise.
if ! swapon --show | grep -q /swapfile; then
  fallocate -l 2G /swapfile && chmod 600 /swapfile && mkswap /swapfile && swapon /swapfile
  echo '/swapfile none swap sw 0 0' >> /etc/fstab
fi

install -d -m 700 /opt/munaiplan /opt/munaiplan/backups
install -d -m 755 /opt/munaiplan/src
echo "Bootstrap complete. Next: tailscale up (log in with your Tailscale account), then deploy/deploy.sh."
