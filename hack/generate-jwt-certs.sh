#!/usr/bin/env bash

set -e

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT_DIR="$(dirname "$CURRENT_DIR")"
JWT_CERTS_DIR=$ROOT_DIR/certs/jwt

mkdir -p $JWT_CERTS_DIR

# generate private key
openssl genrsa -out $JWT_CERTS_DIR/jwt-private.pem 2048  # 2048 is recommended, 4096 is a bit slower
# extatract public key from it
openssl rsa -in $JWT_CERTS_DIR/jwt-private.pem -pubout > $JWT_CERTS_DIR/jwt-public.pem
