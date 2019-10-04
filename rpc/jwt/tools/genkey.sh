#!/usr/bin/env bash
set -e
set -u
set -o pipefail

CUR=`dirname $0`

cd ${CUR}

APP_ID=${1:-app} 
OUT_DIR=${2:-.}

PASS_FILE=${OUT_DIR}/"PASSWORD"

PRIVATE_KEY_FILE=${OUT_DIR}/${APP_ID}.pem
PUBLIC_KEY_FILE=${OUT_DIR}/${APP_ID}.pub
FINGERPRINT_FILE=${OUT_DIR}/${APP_ID}.fp.txt

#####################
# 使用 OpenSSL 方式
#####################
# 1 - Check OpenSSL version
# Do NOT use osx version LibreSSL
# openssl version | grep "OpenSSL"

# 2 - Generate private key
# openssl genrsa -out ${PRIVATE_KEY_FILE} 2048

# 3 - Generate public key
# openssl rsa -in ${PRIVATE_KEY_FILE} -outform PEM -pubout -out ${PUBLIC_KEY_FILE}

# 4 - Generate fingerprint
# openssl rsa -in ${PRIVATE_KEY_FILE} -pubout -outform DER | \
#     openssl sha1 -c | \
#     awk '{print $2}' \
#     > ${FINGERPRINT_FILE}


#####################
# 使用 step 方式
#####################

# 1. 生成秘钥对
step crypto keypair ${PUBLIC_KEY_FILE} ${PRIVATE_KEY_FILE} \
	--kty=RSA \
	--password-file=${PASS_FILE} # 密码文件
	
# 2. 根据公钥（Public Key）生成 SHA1 指纹
step crypto hash digest ${PUBLIC_KEY_FILE} --alg=SHA1 \
    | awk '{print $1}' \
    > ${FINGERPRINT_FILE}

# Generate yaml config
./gen-yml.sh ${APP_ID} ${OUT_DIR}
