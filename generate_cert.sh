#! /usr/bin/env bash

# Vision Screening Upload Simulator
# Copyright (C) 2017 Andrew Allen
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as
# published by the Free Software Foundation, either version 3 of the
# License, or (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program.  If not, see <https://www.gnu.org/licenses/>.

if [ -f certs/good_key.pem ] ||
   [ -f certs/good_certificate.crt ] ||
   [ -f certs/bad_key.pem ] ||
   [ -f certs/bad_certificate.crt ]; then
  cat <<EOF
You have already generated certificates and they are available in the certs
directory.
EOF
  exit 1
fi

openssl req \
  -newkey rsa:2048 \
  -nodes \
  -keyout certs/good_key.pem \
  -x509 \
  -days 365 \
  -subj "/C=AA/ST=State/L=City/O=Testing certificate/CN=*" \
  -out certs/good_certificate.crt

openssl req \
  -newkey rsa:2048 \
  -nodes \
  -keyout certs/evil_key.pem \
  -x509 \
  -days 365 \
  -subj "/C=ZZ/ST=Evil/L=Evil/O=Evil corp/CN=*" \
  -out certs/evil_certificate.crt
