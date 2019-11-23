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
    --enable_tls \
    --tls_cert_file=./cert/server/localhost/localhost.crt \
    --tls_key_file=./cert/server/localhost/localhost.key \
    web \
    --address=${SERVER_IP}:${SERVER_PORT} \
    --namespace=com.jinmuhealth.platform.srv \

