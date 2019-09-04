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
    --config_consul_address=localhost:8500 \
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

