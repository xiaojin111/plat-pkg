#!/usr/bin/env bash
set -e
set -u
set -o pipefail

CUR=`dirname $0`
CONFIG_FILE="${CUR}/errors.yml"
OUT_FILE="${CUR}/error.gen.go"
PKG_NAME="errors"

cd ${CUR}

go run ./gen/gen.go -c ${CONFIG_FILE} -o ${OUT_FILE} -p ${PKG_NAME}
goimports -w ${OUT_FILE} 
