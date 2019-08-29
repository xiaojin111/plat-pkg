#!/usr/bin/env bash
set -e
set -u
set -o pipefail

CUR=`dirname $0`

cd ${CUR}

APP_ID=${1:-app} 
OUT_DIR=${2:-.}

PRIVATE_KEY_FILE=${OUT_DIR}/${APP_ID}.pem
PUBLIC_KEY_FILE=${OUT_DIR}/${APP_ID}.pub
FINGERPRINT_FILE=${OUT_DIR}/${APP_ID}.fp.txt

# Generate private key
openssl genrsa -out ${PRIVATE_KEY_FILE} 2048

# Generate public key
openssl rsa -in ${PRIVATE_KEY_FILE} -outform PEM -pubout -out ${PUBLIC_KEY_FILE}

# Generate fingerprint
openssl rsa -in ${PRIVATE_KEY_FILE} -pubout -outform DER | \
    openssl sha1 -c | \
    awk '{print $2}' \
    > ${FINGERPRINT_FILE}
