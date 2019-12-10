#!/bin/bash
set -e
set -u
set -o pipefail

REF=${1-master} # branch or tag; defaults to 'master' if parameter 1 not present
REMOTE="cert-origin" # just a name to identify the remote
REPO="https://github.com/jinmukeji/cert.git" # replace this with your repository URL
FOLDER="cert" # where to mount the subtree

CUR=`dirname $0`
cd ${CUR}

git remote add $REMOTE --no-tags $REPO
if [[ -d $FOLDER ]]; then # update the existing subtree
    git subtree pull $REMOTE $REF --prefix=$FOLDER --squash -m "Merging '$REF' into '$FOLDER'"
else # add the subtree
    git subtree add  $REMOTE $REF --prefix=$FOLDER --squash -m "Merging '$REF' into '$FOLDER'"
fi
git remote remove $REMOTE
