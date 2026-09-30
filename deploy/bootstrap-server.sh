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
apt-get install -yq ca-certificates curl git unattended-upgrades docker.io docker-compose-v2 docker-buildx
systemctl enable --now docker
dpkg-reconfigure -f noninteractive unattended-upgrades

# Tailscale from its signed apt repository (skipped for a public preview: DEPLOY_MODE=public).
if [ "${DEPLOY_MODE:-tailscale}" != public ]; then
curl -fsSL "https://pkgs.tailscale.com/stable/ubuntu/${VERSION_CODENAME}.noarmor.gpg" -o /usr/share/keyrings/tailscale-archive-keyring.gpg
curl -fsSL "https://pkgs.tailscale.com/stable/ubuntu/${VERSION_CODENAME}.tailscale-keyring.list" -o /etc/apt/sources.list.d/tailscale.list
apt-get update -q
apt-get install -yq tailscale
systemctl enable --now tailscaled
fi

# Firewall: nothing public except SSH (and 80/443 in public mode).
# Docker-published ports bypass host firewalls, which is why the stack only binds 127.0.0.1.
if [ -f /etc/iptables/rules.v4 ] && grep -q 'REJECT' /etc/iptables/rules.v4; then
  # Oracle Cloud Ubuntu images ship their own iptables rules (ufw conflicts with them): open
  # the web ports ahead of the final REJECT and persist. Also allow 80/443 in the VCN security list.
  if [ "${DEPLOY_MODE:-tailscale}" = public ]; then
    for port in 443 80; do
      iptables -C INPUT -p tcp --dport "$port" -m state --state NEW -j ACCEPT 2>/dev/null ||
        iptables -I INPUT 6 -p tcp --dport "$port" -m state --state NEW -j ACCEPT
    done
  else
    iptables -C INPUT -i tailscale0 -j ACCEPT 2>/dev/null || iptables -I INPUT 6 -i tailscale0 -j ACCEPT
  fi
  apt-get install -yq iptables-persistent
  netfilter-persistent save
else
  apt-get install -yq ufw
  ufw default deny incoming
  ufw default allow outgoing
  ufw allow OpenSSH
  if [ "${DEPLOY_MODE:-tailscale}" = public ]; then
    ufw allow 80/tcp
    ufw allow 443/tcp
  else
    ufw allow in on tailscale0
  fi
  ufw --force enable
fi

# 2 GB swap: the ML image build and TensorFlow start-up peak above 4 GB RAM otherwise.
if ! swapon --show | grep -q /swapfile; then
  fallocate -l 2G /swapfile && chmod 600 /swapfile && mkswap /swapfile && swapon /swapfile
  echo '/swapfile none swap sw 0 0' >> /etc/fstab
fi

install -d -m 700 /opt/munaiplan /opt/munaiplan/backups
install -d -m 755 /opt/munaiplan/src
echo "Bootstrap complete (mode ${DEPLOY_MODE:-tailscale}). Next: init-env.sh, then deploy.sh."
