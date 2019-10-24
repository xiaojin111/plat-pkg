# jm-micro

**jm-micro** 命令行工具是对 Go Micro 命令行工具进行定制化的版本，其中包含了针对金姆平台用途相关的插件定制。

## 1 安装与使用

安装 **jm-micro** 命令行工具:

```sh
go get -u github.com/jinmukeji/plat-pkg/jm-micro
```

使用：

```sh
# 查看命令
jm-micro

# e.g. 启动 micro api
jm-micro api \
	--handler=rpc \
	--enable_rpc \
	--namespace=com.jinmuhealth.platform.srv
	
# 常用启动参数
```



## 2 开发与调试

```sh
# 运行(无插件)
go run main.go

# 运行(包含插件)
go run main.go plugin.go
# or
go run .

# 带参数启动方式
go run . api \
	--handler=rpc \
	--enable_rpc \
	--namespace=com.jinmuhealth.platform.srv
	
# 常用启动参数
go run . \
    --log_level=DEBUG \
    --register_interval=5 \
    --register_ttl=10 \
    --client_pool_size=100 \
    --server_name=com.jinmuhealth.platform.api \
    --config_etcd_address=localhost:2379 \
    api \
    --address=0.0.0.0:8080 \
    --handler=rpc \
    --enable_rpc \
    --namespace=com.jinmuhealth.platform.srv \
    --enable_jwt
    
# 查看参数信息
go run . -h

# 查看 api 子命令参数信息
go run . api -h
```

注意事项：

- 使用 `--enable_jwt` 参数之前，需要指定好配置文件加载相关参数，并确保配置中包含 JWT 所需的配置信息

## 3 TLS 调试

启动 TLS 模式，需要指定一下几个启动参数：

- **`--enable_tls`**: 启用 TLS
- **`--tls_cert_file`**: TLS 证书文件。要求证书 SAN 中包含请求地址（DNS或IP地址）。
- **`--tls_key_file`**: TLS 秘钥文件。
- **`--tls_client_ca_file`**: 客户端证书CA的根证书文件。如果启用客户端证书，则指定本参数，可以用来做双向的验证。

```sh
go run . \
    --log_level=DEBUG \
    --register_interval=5 \
    --register_ttl=10 \
    --client_pool_size=100 \
    --server_name=com.jinmuhealth.platform.api \
    --metadata=X-Err-Style=MicroDetailed \
    --config_etcd_address=localhost:2379 \
    --enable_tls \
    --tls_cert_file=./cert/localhost.crt \
    --tls_key_file=./cert/localhost.key \
    --tls_client_ca_file=./cert/root_ca.crt \
    api \
    --address=${SERVER_IP}:${SERVER_PORT} \
    --handler=rpc \
    --enable_rpc \
    --namespace=com.jinmuhealth.platform.srv \
```

如果使用 Postman 进行调试，则需要在 **Settings** 中配置证书信任：

![Trust](cert/postman.png)

1. 启用 **CA Certificates**，并设置根证书文件 `cert/root_ca.crt`。信任后，Postman请求时将不会再报告警告。
2. 设置 **Client Certificates**，添加一个客户端证书配置：
   - **Host:** `127.0.0.1:8080`
   - **CRT file:** `cert/jm-app.crt`
   - **KEY file:** `cert/jm-app.key`

