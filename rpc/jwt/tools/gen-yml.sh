#!/usr/bin/env bash
set -e
set -u
set -o pipefail

CUR=`dirname $0`

cd ${CUR}

export APP_ID=${1:-app}
OUT_DIR=${2:-.}
export DISABLED="false"
export FINGERPRINT=$(head -n 1 ${OUT_DIR}/${APP_ID}.fp.txt)
export PUBLIC_KEY_FILE="${OUT_DIR}/${APP_ID}.pub"
OUT_FILE=${OUT_DIR}/${APP_ID}.yml

rm -f temp.yml
( echo "cat <<EOF > ${OUT_FILE}";
  cat template.yml;
  echo "EOF";
) >temp.yml
. temp.yml
rm -f temp.yml
cat ${OUT_FILE}
