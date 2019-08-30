#!/usr/bin/env bash
set -e
set -u
set -o pipefail

CUR=`dirname $0`

cd ${CUR}

SERVER_IP=0.0.0.0
SERVER_PORT=8080

LOG_LEVEL=DEBUG go run . \
    --register_interval=5 \
    --register_ttl=10 \
    --client_pool_size=100 \
    api \
    --address=${SERVER_IP}:${SERVER_PORT} \
	--handler=rpc \
	--enable_rpc \
	--namespace=com.jinmuhealth.platform.srv
