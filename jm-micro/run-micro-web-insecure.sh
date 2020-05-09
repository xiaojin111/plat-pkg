#!/usr/bin/env bash
set -e
set -u
set -o pipefail

CUR=`dirname $0`

cd ${CUR}

NAMESPACE="com.jinmuhealth.examples"
SERVER_ADDR="0.0.0.0:8082"
SERVER_NAME="${NAMESPACE}.web.jm_micro"
ETCD_ADDR="localhost:2379"

go run . \
    --log_level=DEBUG \
    --server_name=${SERVER_NAME} \
    web \
    --address=${SERVER_ADDR} \
    --namespace=${NAMESPACE} \


# go run . \
#     --log_level=DEBUG \
#     --server_name=com.jinmuhealth.platform.api.jm_micro \
#     api \
#     --namespace="com.jinmuhealth.platform" \
#     --handler=rpc \
#     --enable_rpc

# go run . \
#     --log_level=DEBUG \
#     --server_name=com.jinmuhealth.platform.api.jm_micro \
#     --api_namespace="com.jinmuhealth.platform" \
#     api \
#     --namespace="com.jinmuhealth.platform" \
#     --handler=api
