#关闭服务
CUR=`dirname $0`
cd ${CUR}/..
docker stop etcd || echo "etcd stop"
