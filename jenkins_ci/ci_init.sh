# 初始化项目
go version
# set -e
GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
CUR=`dirname $0`
cd ${CUR}/..
# go  mod download
make setup
