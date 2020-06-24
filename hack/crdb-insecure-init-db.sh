#!/usr/bin/env bash

set -e

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT_DIR="$(dirname "$CURRENT_DIR")"

ENV_FILE=$ROOT_DIR/.env

source $ENV_FILE

cockroach sql --insecure --host="localhost:$DB_PORT" --execute="CREATE DATABASE IF NOT EXISTS $DB_NAME;"
