# plat-pkg ![License](https://img.shields.io/github/license/jinmukeji/plat-pkg) [![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/gitee.com/jt-heath/plat-pkg/v2?tab=doc) [![Go Report Card](https://goreportcard.com/badge/jinmukeji/plat-pkg)](https://goreportcard.com/report/gitee.com/jt-heath/plat-pkg) 

金姆微服务平台基础工具包。

## License

plat-pkg 使用 [LGPLv3 协议](./LICENSE)。



## TODO

**架构**

- [ ] 重构配置信息管理方案

**部署**

- [ ] 支持 k8s 部署

**gRPC**

- [ ] 支持 TLS 通讯（服务与服务之间，jm-micro web 与服务之间）
- [ ] TLS 自签发证书
- [ ] 支持 gRPC Reflection API

**日志**

- [x] 集成 logrus
- [ ] 支持 grom 自定义 logger

**数据**

- [ ] 重构 MySQL Client 包，简化处理机制
- [ ] 分布式事物支持 (Saga or Event Soucing ?)

**jm-micro**

- [ ] `jm-micro api` 的使用方法
- [ ] `jm-micro proxy` 的使用方法
- [ ] `jm-micro tunnel` 的使用方法
- [ ] 其它 `jm-micro` 指令使用示例

**proto**

- [ ] 使用 buf 工具管理生命周期

**开发与测试工具**

- [ ] gRPC 测试工具与方法
- [ ] 尝试引入 Makego 管理 Go 代码
- [ ] 完善示例代码场景
  - [ ] 最简 API 调用
  - [ ] 服务间 API 调用
  - [ ] 使用配置中心信息
  - [ ] 数据访问示例
  - [ ] Broker 示例
  - [ ] 分布式场景示例