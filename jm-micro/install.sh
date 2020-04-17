#!/usr/bin/env bash
set -e
set -u
set -o pipefail

CUR=`dirname $0`
cd ${CUR}
go install
