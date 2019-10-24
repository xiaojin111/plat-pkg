package service

import (
	"fmt"

	"github.com/micro/cli"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder/yaml"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-micro/config/source/etcd"
	"github.com/micro/go-micro/config/source/file"
)

// Config 相关常量
const (
	DefaultConfigEnvPrefix  = "JM"
	DefaultConfigEtcdPrefix = "/micro/config/jm"
)

// config 相关变量
var (
	cfgFiles                                 cli.StringSlice
	cfgEnvPrefix, cfgEtcdAddr, cfgEtcdPrefix string
)

func configCliFlags() []cli.Flag {
	return []cli.Flag{
		// Config 相关
		cli.StringSliceFlag{
			Name:  "config_file",
			Usage: "Config file path",
			Value: &cfgFiles,
		},

		cli.StringFlag{
			Name:        "config_env_prefix",
			Usage:       "Config environment variables prefix",
			Value:       DefaultConfigEnvPrefix, // default value
			Destination: &cfgEnvPrefix,
		},

		cli.StringFlag{
			Name:        "config_etcd_address",
			Usage:       "Etcd config source address",
			Destination: &cfgEtcdAddr,
		},

		cli.StringFlag{
			Name:        "config_etcd_prefix",
			Usage:       "Etcd config K/V prefix",
			Value:       DefaultConfigEtcdPrefix, // default value
			Destination: &cfgEtcdPrefix,
		},
	}
}

func loadServiceConfig() error {
	// 加载以下配置信息数据源，优先级依次从低到高：
	// 1. Etcd K/V 配置中心
	// 2. 配置文件，YAML格式
	// 3. 环境变量

	encoder := yaml.NewEncoder()

	// Load config from etcd
	if cfgEtcdAddr != "" {
		etcdSource := etcd.NewSource(
			// optionally specify etcd address;
			etcd.WithAddress(cfgEtcdAddr),
			// optionally specify prefix;
			etcd.WithPrefix(cfgEtcdPrefix),
			// optionally strip the provided prefix from the keys
			etcd.StripPrefix(true),
			source.WithEncoder(encoder),
		)

		if err := config.Load(etcdSource); err != nil {
			return fmt.Errorf("failed to load config from etcd at %s with prefix of [%s]: %w", cfgEtcdAddr, cfgEtcdPrefix, err)
		}

		log.Infof("Loaded config from etcd at %s with prefix of [%s]", cfgEtcdAddr, cfgEtcdPrefix)
	}

	// Load config from files
	for _, f := range cfgFiles.Value() {
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
		env.WithStrippedPrefix(cfgEnvPrefix),
	)
	if err := config.Load(envSource); err != nil {
		return fmt.Errorf("failed to load config from environment variables: %w", err)
	}

	log.Infof("Loaded config from environment variables with prefix of [%s]", cfgEnvPrefix)

	return nil
}
