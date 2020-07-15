#!/usr/bin/env bash

set -e

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT_DIR="$(dirname "$CURRENT_DIR")"

ENV_FILE=$ROOT_DIR/.env

source $ENV_FILE

read -p "What migration type you want to create? (migrate/seed/sample): " type

if [ "${type}" == "migrate" ] || [ "${type}" == "migrates" ]; then
	MIGRATION_DIR=$ROOT_DIR/db/migrations
elif  [ "${type}" == "seed" ] ||  [ "${type}" == "seeds" ]; then
	MIGRATION_DIR=$ROOT_DIR/db/seeds
elif [ "${type}" == "sample" ] || [ "${type}" == "samples" ]; then
	MIGRATION_DIR=$ROOT_DIR/db/samples
else
	echo "Please enter valid migration type to continue!"
	exit 1;
fi

read -p "Enter ${type} name: " name

migrate create -ext sql -dir $MIGRATION_DIR -seq $name
