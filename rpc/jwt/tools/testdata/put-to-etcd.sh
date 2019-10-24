#!/usr/bin/env bash
set -e
set -u
set -o pipefail

CUR=`dirname $0`
KEY_PREFIX="micro/config/jm/platform/app-key"

cd ${CUR}

for FILE in *.yml
do
    CFG_KEY=${FILE%.*}
    cat ${FILE} | etcdctl put ${KEY_PREFIX}/${CFG_KEY}
done
