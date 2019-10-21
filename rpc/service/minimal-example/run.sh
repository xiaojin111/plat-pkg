#!/usr/bin/env bash
set -e
set -u
set -o pipefail

CUR=`dirname $0`

cd ${CUR}

SERVER_ADDR=0.0.0.0:9090
CONSUL_ADDR=localhost:8500

# 启动服务
go run . \
    --log_level=DEBUG \
    --register_interval=5 \
    --register_ttl=10 \
    --server_address=${SERVER_ADDR} \
    --client_pool_size=100 \
    --config_consul_address=localhost:8500 \
    #  --enable_jwt \
