#!/usr/bin/env bash
set -e
set -u
set -o pipefail

CUR=$(dirname $0)
cd ${CUR}

JWT_ALG="RS256"
JWT_AUD="com.jinmuhealth.partner.dayima.api.sys-intg" # 端点 SRV 名称
JWT_SUB="DaYiMa"                                      # 客户名称
JWT_ISS="dayima"                                      # APP ID
JWT_PRIVATE_KEY="dayima.pem"                          # Private Key
PASS_FILE="PASSWORD"

# UNIX time: NOW
JWT_IAT=$(date +'%s')
# UNIX time: NOW+600s
JWT_EXP=$(date -r ${JWT_IAT} -v+600S +'%s')

JWT=$(step crypto jwt sign \
    --alg=${JWT_ALG} \
    --iss=${JWT_ISS} \
    --sub=${JWT_SUB} \
    --aud=${JWT_AUD} \
    --iat=${JWT_IAT} \
    --exp=${JWT_EXP} \
    --key=${JWT_PRIVATE_KEY} \
    --password-file=${PASS_FILE})

echo $JWT

if command -v pbcopy &>/dev/null; then
    echo $JWT | pbcopy
    echo
    echo "JWT has been copied to clipboard."
fi
