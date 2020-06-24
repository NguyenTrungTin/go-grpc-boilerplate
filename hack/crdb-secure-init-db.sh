#!/usr/bin/env bash

set -e

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT_DIR="$(dirname "$CURRENT_DIR")"

ENV_FILE=$ROOT_DIR/.env

source $ENV_FILE

echo -n "Enter your CockroachDB root certs: "
# shellcheck disable=SC2162
read CRDB_ROOT_CERTS

cockroach sql --certs-dir="$CRDB_ROOT_CERTS" --host="localhost:$DB_PORT" --execute="CREATE DATABASE IF NOT EXISTS $DB_NAME; \
																CREATE USER IF NOT EXISTS $DB_USER WITH PASSWORD $DB_PASSWORD; \
																GRANT ALL ON DATABASE $DB_NAME TO $DB_USER;"
