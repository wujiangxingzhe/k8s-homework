#!/usr/bin/env bash

GO_PROJEDT_DIR=$1
BIN_FILE=$2
DOCKER_CONTEXT=$3

cd "${GO_PROJEDT_DIR}"
test -f "${BIN_FILE}" && rm -rf "${BIN_FILE}"
go build -o "${BIN_FILE}" main.go

\cp -rpf "$BIN_FILE" "$DOCKER_CONTEXT"