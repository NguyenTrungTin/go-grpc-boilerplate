#!/usr/bin/env bash

set -e

protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative *.proto
