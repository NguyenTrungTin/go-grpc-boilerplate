#!/usr/bin/env bash

set -e

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT_DIR="$(dirname "$CURRENT_DIR")"

MIGRATION_DIR=$ROOT_DIR/db/migrations

read -p "Enter database url (DNS - Data Source Name): " COCKROACHDB_URL

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

read -p "What command you want to run (up/down): " command

read -p "Enter ${type} version (just enter to set default - ${type} all): " version

migrate -path $MIGRATION_DIR -database $COCKROACHDB_URL $command $version
