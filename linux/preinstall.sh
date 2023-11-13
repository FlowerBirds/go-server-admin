#!/usr/bin/env sh
set -eu

[ -f /etc/profile.d/bash-record.sh ] && rm -f /etc/profile.d/bash-record.sh
[ -d /var/lib/server-admin/static ] && rm -rf /var/lib/server-admin/static
echo clear old bash and static file
