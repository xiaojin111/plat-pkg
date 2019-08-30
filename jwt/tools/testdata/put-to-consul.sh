#!/usr/bin/env bash
set -e
set -u
set -o pipefail

CUR=`dirname $0`
KEY_PREFIX="micro/config/jm/app-key"

cd ${CUR}

for FILE in *.yml
do
    CFG_KEY=${FILE%.*}
    cat ${FILE} | consul kv put ${KEY_PREFIX}/${CFG_KEY} -
done
