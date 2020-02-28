#!/usr/bin/env bash
set -e
set -u
set -o pipefail

CUR=`dirname $0`

cd ${CUR}

SERVER_IP=0.0.0.0
SERVER_PORT=8082
ETCD_ADDR=localhost:2379

go run . \
    --registry=etcd \
    --log_level=DEBUG \
    --server_name=com.jinmuhealth.platform.web \
    web \
    --address=${SERVER_IP}:${SERVER_PORT} \
    --namespace=com.jinmuhealth.platform.srv \

