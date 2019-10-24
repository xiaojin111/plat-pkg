package configloader

import (
	"fmt"
	"net/http"

	"github.com/micro/cli"
	"github.com/micro/micro/plugin"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder/yaml"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-micro/config/source/etcd"
	"github.com/micro/go-micro/config/source/file"

	mlog "github.com/jinmukeji/go-pkg/log"
)

// Config 相关常量
const (
	DefaultConfigEnvPrefix  = "JM"
	DefaultConfigEtcdPrefix = "micro/config/jm"
)

type configLoaderPlugin struct {
	cfgFiles                                 cli.StringSlice
	cfgEnvPrefix, cfgEtcdAddr, cfgEtcdPrefix string
}

var (
	log *mlog.Logger = mlog.StandardLogger()
)

func (p *configLoaderPlugin) Flags() []cli.Flag {
	return []cli.Flag{
		// Config 相关
		cli.StringSliceFlag{
			Name:  "config_file",
			Usage: "Config file path",
			Value: &p.cfgFiles,
		},

		cli.StringFlag{
			Name:        "config_env_prefix",
			Usage:       "Config environment variables prefix",
			Value:       DefaultConfigEnvPrefix, // default value
			Destination: &p.cfgEnvPrefix,
		},

		cli.StringFlag{
			Name:        "config_etcd_address",
			Usage:       "Etcd config source address",
			Destination: &p.cfgEtcdAddr,
		},

		cli.StringFlag{
			Name:        "config_etcd_prefix",
			Usage:       "Etcd config K/V prefix",
			Value:       DefaultConfigEtcdPrefix, // default value
			Destination: &p.cfgEtcdPrefix,
		},
	}
}

func (p *configLoaderPlugin) Commands() []cli.Command {
	return nil
}

func (p *configLoaderPlugin) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		// 什么都不包装，透传
		return h
	}
}

func (p *configLoaderPlugin) Init(ctx *cli.Context) error {
	// 加载以下配置信息数据源，优先级依次从低到高：
	// 1. Etcd K/V 配置中心
	// 2. 配置文件，YAML格式
	// 3. 环境变量

	encoder := yaml.NewEncoder()

	// Load config from etcd
	if p.cfgEtcdAddr != "" {
		etcdSource := etcd.NewSource(
			// optionally specify etcd address;
			etcd.WithAddress(p.cfgEtcdAddr),
			// optionally specify prefix;
			etcd.WithPrefix(p.cfgEtcdPrefix),
			// optionally strip the provided prefix from the keys
			etcd.StripPrefix(true),
			source.WithEncoder(encoder),
		)

		if err := config.Load(etcdSource); err != nil {
			werr := fmt.Errorf("failed to load config from etcd at %s with prefix of [%s]: %w", p.cfgEtcdAddr, p.cfgEtcdPrefix, err)
			log.Error(werr)
			return werr
		}

		log.Infof("Loaded config from etcd at %s with prefix of [%s]", p.cfgEtcdAddr, p.cfgEtcdPrefix)
	}

	// Load config from files
	for _, f := range p.cfgFiles.Value() {
		fileSource := file.NewSource(
			file.WithPath(f),
			source.WithEncoder(encoder),
		)

		if err := config.Load(fileSource); err != nil {
			return fmt.Errorf("failed to load config file %s: %w", f, err)
		}

		log.Infof("Loaded config from file: %s", f)
	}

	// Load config from env
	envSource := env.NewSource(
		// optionally specify prefix
		env.WithStrippedPrefix(p.cfgEnvPrefix),
	)
	if err := config.Load(envSource); err != nil {
		return fmt.Errorf("failed to load config from environment variables: %w", err)
	}

	log.Infof("Loaded config from environment variables with prefix of [%s]", p.cfgEnvPrefix)

	return nil
}

func (p *configLoaderPlugin) String() string {
	return "config-loader"
}

func NewPlugin() plugin.Plugin {
	return NewConfigLoader()
}

func NewConfigLoader() plugin.Plugin {
	return &configLoaderPlugin{}
}
