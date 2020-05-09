#!/usr/bin/env bash
set -e
set -u
set -o pipefail

CUR=`dirname $0`

cd ${CUR}

NAMESPACE="com.jinmuhealth.examples"
SERVER_ADDR="0.0.0.0:8080"
SERVER_NAME="${NAMESPACE}.api.jm_micro"
ETCD_ADDR="localhost:2379"

# =============================================
# 启用 micro api 的仅 rpc handler + api handler 方式
# 本方式下可以访问 /rpc 调用API
# 也可以同时访问 /<micro_api_service_name>//[rpc_service]/[rpc_method]
# 参考：
#	https://micro.mu/docs/api.html#api-handler
#	https://micro.mu/docs/api.html#rpc-handler
#	https://micro.mu/docs/api.html#rpc-resolver
# =============================================
go run . \
    --log_level=DEBUG \
    --server_name=${SERVER_NAME} \
    --metadata=X-Err-Style=MicroDetailed \
    api \
    --address=${SERVER_ADDR} \
    --handler=api \
    --enable_rpc \
    --namespace=${NAMESPACE} \
    --enable_tcp_healthcheck
    --tcp_healthcheck_addr=:9901

# =============================================
# 启用 micro api 的仅 rpc handler 方式
# 本方式下仅能访问 /rpc 调用API
# 参考：
#	https://micro.mu/docs/api.html#rpc-handler
# =============================================
# go run . \
#     --log_level=DEBUG \
#     --server_name=${SERVER_NAME} \
#     --metadata=X-Err-Style=MicroDetailed \
#     api \
#     --address=${SERVER_ADDR} \
#     --handler=rpc \
#     --enable_rpc \
#     --namespace=${NAMESPACE} \
#     --enable_tcp_healthcheck
#     --tcp_healthcheck_addr=:9901

# go run . \
#     --log_level=DEBUG \
#     --register_interval=5 \
#     --register_ttl=10 \
#     --client_pool_size=100 \
#     --server_name=${SERVER_NAME} \
#     --metadata=X-Err-Style=MicroDetailed \
#     --config_etcd_address=${ETCD_ADDR} \
#     api \
#     --address=${SERVER_ADDR} \
#     --handler=rpc \
#     --enable_rpc \
#     --namespace=${NAMESPACE} \
#     --enable_tcp_healthcheck
#     --tcp_healthcheck_addr=:9901
#     # --enable_jwt \
#     # --jwt_max_exp_interval=600s \
#     # --rpc_whitelist=com.jinmuhealth.platform.srv.template-service
