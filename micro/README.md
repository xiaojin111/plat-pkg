基于 go-micro v2.4.0 说明

## 创建及初始化服务

```go
// New Service
service := micro.NewService(
  micro.Name("com.foo.service.hello"),
  micro.Version("latest"),
)
```

`micro.Option` 说明：

1. **micro.Name(n string) Option** ， 指定服务名称。命名规则一般是“$namespace.$type.$name”。其中namespace 代表项目的名称空间， type代表服务类型（例如gRPC 和 web），一般会把gRPC service类型缩写成 srv。服务实例运行后， 此名称将自动注册到 Registry， 成为服务发现的依据。默认为“go.micro.server”。 **注**：因此此项必须要指定， 否则所有节点使用相同的默认名称，会导致调用混乱
2. **micro.Version(v string) Option**，指定服务版本。默认为启动时间格式化的字符串。恰当地选择版本号再配合相应的Selector， 可以实现优雅的轮转升级、灰度发布、A/B 测试等功能。
3. **micro.Address(addr string) Option**，指定gRPC 服务地址。 默认为随机端口。由于客户端是通过注册中心来定位服务， 所以随机端口并不影响使用。 但实践中经常是指定固定端口号的， 这会有利于运维管理和安全控制
4. **micro.RegisterTTL(t time.Duration) Option**，指定服务注册信息在注册中心的有效期。 默认为一分种
5. **micro.RegisterInterval(t time.Duration) Option**，指定服务主动向注册中心报告健康状态的时间间隔， 默认为30秒。 这两个注册中心相关的Option结合起来用，可以避免因服务意外宕机而未通知注册中心，产生“无效注册信息”
6. **micro.WrapHandler(w …server.HandlerWrapper) Option**，包装服务Handler， 概念上类似于 [Gin Middleware](https://github.com/gin-gonic/gin#using-middleware)， 集中控制Handler行为。可包装多层，执行顺序由外到内。
7. **micro.WrapSubscriber(w …server.SubscriberWrapper) Option**，与WrapHandler相似，不同之处在于它用来包装异步消费处理中的“订阅者”。
8. **micro.WrapCall(w …client.CallWrapper) Option**，包装客户端发起的每一次方法调用。
9. **micro.WrapClient(w …client.Wrapper) Option**，包装客户端，可包装多层， 执行顺序由内到外。
10. **micro.BeforeStart(fn func() error) Option**，设置服务启动前回调函数，可设置多个。
11. **micro.BeforeStop(fn func() error) Option**，设置服务关闭前回调函数，可设置多个。
12. **micro.AfterStart(fn func() error) Option**，设置服务启动后回调函数，可设置多个。
13. **micro.AfterStop(fn func() error) Option**，设置服务关闭后回调函数，可设置多个。
14. **micro.Action(a func(\*cli.Context)) Option**，处理命令行参数。 支持子命令及控制标记。 详情请见 [micro/cli](https://github.com/micro/cli)
15. **micro.Flags(flags …cli.Flag) Option**，快捷支持命令行控制标记， 详情请见 [micro/cli](https://github.com/micro/cli)
16. **micro.Cmd(c cmd.Cmd) Option**， 指定命令行处理对象。 默认由 [newCmd](https://github.com/micro/go-micro/blob/v2.4.0/config/cmd/cmd.go#L383)生成，此对象包含了一系列默认的环境变量、命令行参数支持。 可以看作是多个内置cli.Flag的集合。**注**： go-micro 框架对命令行处理的设计方案有利有弊。 利是提供大量默认选项，可以节省开发者时间。 弊是此设计对用户程序的**有强烈的侵入性**： 框架要求开发者必须以 micro/cli 统一要求的方式来处理命令行参数。如若不然， 程序会报错无法运行。 例如，我们运行 `./hello-service --foo=bar` 就会报出“**Incorrect Usage. flag provided but not defined: -foo=bar**”的错误。 好在有这个Option，可以弥补这种强侵入性带来的弊端。假如一个现存项目想引入Micro ，而它已经有自己的参数处理机制， 那么就需要使用此Option覆盖默认行为（同时丢掉一些默认的参数处理能力）。 关于命令行参数， 本文后面部分有进一步解释。
17. **micro.Metadata(md map[string]string) Option**，指定服务元数据。 元数据时常被用来为服务标记与分组， 实现特定的负载策略等
18. **micro.Transport(t transport.Transport) Option**，指定传输协议， 默认为http协议
19. **micro.Selector(s selector.Selector) Option** ，指定节点选择器， 实现不同负载策略。默认为随机Selector
20. **micro.Registry(r registry.Registry) Option**，指定用于服务发现的注册机制， 默认为基于 mDNS 的注册机制
21. **micro.Server(s server.Server) Option**， 指定自定义Server， 用于默认Server不满足业务要求的情况。默认为rpcServer
22. **micro.HandleSignal(b bool) Option**， 是否允许服务自动响应 TERM, INT, QUIT 等信号。默认为true
23. **micro.Context(ctx context.Context) Option**，指定服务初始Context，默认为context.BackGround()，可用于控制服务生存期及其它
24. **micro.Client(c client.Client) Option**，指定对外调用的客户端。 默认为rpcClient
25. **micro.Broker(b broker.Broker) Option**， 指定用于 发布/订阅 消息通讯的Broker。默认为http broker
26. **micro.Profile(p profile.Profile) Option**，指定Profile对象，用于性能调优
27. **micro.Tracer(t trace.Tracer) Option**，指定Trace对象，方便性能跟踪
28. **micro.Auth(a auth.Auth) Option**，指定Auth对象， 用于自定义认证。 （根据官方Slack的说法，API尚不稳定，不建议v2.4.0中使用）
29. **micro.Config(c config.Config) Option**，批定Config对象， 用于自定义配置。



常用命令

```sh
# 列出可用服务
micro list services

# 查看服务的详细情况
micro get service com.foo.service.hello

# 发起调用
micro call com.foo.service.hello Hello.Call '{"name": "Bill"}'
```

交互式访问

```sh
micro cli
```

