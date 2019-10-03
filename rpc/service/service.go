package service

import (
	"fmt"
	"os"
	"time"

	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"

	wsvc "github.com/micro/go-plugins/wrapper/service"

	wcid "github.com/jinmukeji/plat-pkg/rpc/wrapper/cid"
	wfm "github.com/jinmukeji/plat-pkg/rpc/wrapper/formatmeta"
	wjwt "github.com/jinmukeji/plat-pkg/rpc/wrapper/jwt"
	wlog "github.com/jinmukeji/plat-pkg/rpc/wrapper/log"
	wme "github.com/jinmukeji/plat-pkg/rpc/wrapper/microerr"

	cm "github.com/jinmukeji/plat-pkg/rpc/ctxmeta"
)

func CreateService(opts *Options) micro.Service {
	// jmSvc:= newJMService(opts)

	// 设置 service，并且加载配置信息
	svc := newService(opts)
	err := setupService(svc, opts)
	die(err)

	// 设置 server
	srv := svc.Server()
	err = setupServer(srv, opts)
	die(err)

	return svc
}

const (
	// DefaultRegisterTTL specifies how long a registration should exist in
	// discovery after which it expires and is removed
	DefaultRegisterTTL = 30 * time.Second

	// DefaultRegisterInterval is the time at which a service should re-register
	// to preserve it’s registration in service discovery.
	DefaultRegisterInterval = 15 * time.Second
)

func newService(opts *Options) micro.Service {
	versionMeta := opts.ServiceMetadata()

	// Create a new service. Optionally include some options here.
	svc := micro.NewService(
		// Service Basic Info
		micro.Name(opts.FQDN()),
		micro.Version(opts.ProductVersion),

		// Fault Tolerance - Heartbeating
		// 	 See also: https://micro.mu/docs/fault-tolerance.html#heartbeating
		micro.RegisterTTL(DefaultRegisterTTL),
		micro.RegisterInterval(DefaultRegisterInterval),

		// Setup metadata
		micro.Metadata(versionMeta),
	)

	svc.Options().Cmd.App().Description = fmt.Sprintf("fqdn: %s", opts.FQDN())

	return svc
}

func setupService(svc micro.Service, opts *Options) error {
	// 设置启动参数
	flags := defaultFlags()
	if len(opts.Flags) > 0 {
		flags = append(flags, opts.Flags...)
	}

	svc.Init(
		// Setup runtime flags
		micro.Flags(flags...),

		micro.Action(func(c *cli.Context) {
			if opts.CliPreAction != nil {
				opts.CliPreAction(c)
			}

			if c.Bool("version") {
				printFullVersionInfo(opts)
				os.Exit(0)
			}

			setupLogger(log, opts.Name)
			// 启动阶段打印版本号
			// 由于内部使用到了 logger，需要在 logger 被设置后调用
			logVersionInfo(opts)

			// 加载 config
			err := loadServiceConfig()
			die(err)

			if opts.CliPostAction != nil {
				opts.CliPostAction(c)
			}
		}),
	)

	// Setup wrappers
	setupHandlerWrappers(svc, opts)

	return nil
}

func defaultFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.BoolFlag{
			Name:  "version",
			Usage: "Show version information",
		},
	}

	flags = append(flags, logCliFlags()...)
	flags = append(flags, configCliFlags()...)
	flags = append(flags, jwtFlags()...)

	return flags
}

// JWT 相关
var (
	jwtOption = wjwt.DefaultOptions()
	enableJwt = false
)

func jwtFlags() []cli.Flag {
	return []cli.Flag{
		// JWT 相关
		cli.BoolFlag{
			Name:        "enable_jwt",
			Usage:       "Enable JWT validation",
			EnvVar:      "ENABLE_JWT",
			Destination: &enableJwt,
		},
		cli.StringFlag{
			Name:        "jwt_key",
			Usage:       "JWT HTTP header key",
			EnvVar:      "JWT_KEY",
			Value:       cm.MetaJwtKey,
			Destination: &(jwtOption.HeaderKey),
		},
		cli.StringFlag{
			Name:        "jwt_config_path",
			Usage:       "Micro config path for JWT",
			EnvVar:      "JWT_CONFIG_PATH",
			Value:       wjwt.DefaultMicroConfigPath,
			Destination: &(jwtOption.MicroConfigPath),
		},
		cli.DurationFlag{
			Name:        "jwt_max_exp_interval",
			Usage:       "JWT max expiration interval",
			EnvVar:      "JWT_MAX_EXP_INTERVAL",
			Value:       wjwt.DefaultMaxExpInterval,
			Destination: &(jwtOption.MaxExpInterval),
		},
	}
}

func setupHandlerWrappers(svc micro.Service, opts *Options) {
	// 设置 Server Handler Wrappers
	srvWrappers := []server.HandlerWrapper{}

	// 自定义 pre
	if len(opts.PreServerHandlerWrappers) > 0 {
		srvWrappers = append(srvWrappers, opts.PreServerHandlerWrappers...)
	}

	srvWrappers = append(srvWrappers,
		// 默认的的 wrappers
		wsvc.NewHandlerWrapper(svc),
		wfm.FormatMetadataWrapper,
		wcid.CidWrapper,
		wme.MicroErrWrapper,
		wlog.LogWrapper,
	)

	if enableJwt {
		srvWrappers = append(srvWrappers, wjwt.NewHandlerWrapper(jwtOption))
	}

	// 自定义 post
	if len(opts.PostServerHandlerWrappers) > 0 {
		srvWrappers = append(srvWrappers, opts.PostServerHandlerWrappers...)
	}

	svc.Init(micro.WrapHandler(srvWrappers...))

	// 设置 Client Wrappers
	clientWrappers := []client.Wrapper{}
	if len(opts.PreClientWrappers) > 0 {
		clientWrappers = append(clientWrappers, opts.PreClientWrappers...)
	}

	clientWrappers = append(clientWrappers,
		// 默认的的 wrappers
		wsvc.NewClientWrapper(svc),
	)
	if len(opts.PostClientWrappers) > 0 {
		clientWrappers = append(clientWrappers, opts.PostClientWrappers...)
	}

	svc.Init(micro.WrapClient(clientWrappers...))
}

func setupServer(srv server.Server, opts *Options) error {

	err := srv.Init(
		// Graceful shutdown of a service using the server.Wait option
		// The server deregisters the service and waits for handlers to finish executing before exiting.
		server.Wait(nil),
	)
	if err != nil {
		return err
	}

	if opts.RegisterServer != nil {
		err = opts.RegisterServer(srv)
		if err != nil {
			return err
		}
	}

	return nil
}

func die(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
