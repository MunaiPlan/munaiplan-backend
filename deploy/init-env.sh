#!/bin/sh
# Creates /opt/munaiplan/.env with generated secrets. Never overwrites an existing file.
# Usage: init-env.sh admin@your-company.kz
set -eu
file=/opt/munaiplan/.env
email=${1:?usage: init-env.sh ADMIN_EMAIL}
if [ -e "$file" ]; then echo "$file exists; not changed."; exit 0; fi
umask 077
hex() { openssl rand -hex "$1"; }
# PASSWORD_SALT + ";" + password must stay within bcrypt's 72 bytes, so both are 32 hex chars.
cat > "$file" <<ENV
DB_PASSWORD=$(hex 32)
PASSWORD_SALT=$(hex 16)
USER_ACCESS_TOKEN_SECRET=$(hex 32)
USER_REFRESH_TOKEN_SECRET=$(hex 32)
ADMIN_EMAIL=$email
ADMIN_PASSWORD=$(hex 12)
ENV
echo "Created $file (mode 600). Read ADMIN_PASSWORD there on the server; do not copy it into chat or tickets."
