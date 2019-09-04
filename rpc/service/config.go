package service

import (
	"fmt"

	"github.com/micro/cli"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder/yaml"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/consul"
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-micro/config/source/file"
)

// Config 相关常量
const (
	DefaultConfigEnvPrefix    = "JM"
	DefaultConfigConsulPrefix = "/micro/config/jm"
)

// config 相关变量
var (
	cfgFiles                                     cli.StringSlice
	cfgEnvPrefix, cfgConsulAddr, cfgConsulPrefix string
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
			Name:        "config_consul_address",
			Usage:       "Consul config source address",
			Destination: &cfgConsulAddr,
		},

		cli.StringFlag{
			Name:        "config_consul_prefix",
			Usage:       "Consul config K/V prefix",
			Value:       DefaultConfigConsulPrefix, // default value
			Destination: &cfgConsulPrefix,
		},
	}
}

func loadServiceConfig() error {
	// 加载以下配置信息数据源，优先级依次从低到高：
	// 1. Consul K/V 配置中心
	// 2. 配置文件，YAML格式
	// 3. 环境变量

	encoder := yaml.NewEncoder()

	// Load config from consul
	if cfgConsulAddr != "" {
		consulSource := consul.NewSource(
			// optionally specify consul address;
			consul.WithAddress(cfgConsulAddr),
			// optionally specify prefix;
			consul.WithPrefix(cfgConsulPrefix),
			// optionally strip the provided prefix from the keys
			consul.StripPrefix(true),
			source.WithEncoder(encoder),
		)

		if err := config.Load(consulSource); err != nil {
			return fmt.Errorf("failed to load config from consul at %s with prefix of [%s]: %s", cfgConsulAddr, cfgConsulPrefix, err)
		}

		log.Infof("Loaded config from consul at %s with prefix of [%s]", cfgConsulAddr, cfgConsulPrefix)
	}

	// Load config from files
	for _, f := range cfgFiles.Value() {
		fileSource := file.NewSource(
			file.WithPath(f),
			source.WithEncoder(encoder),
		)

		if err := config.Load(fileSource); err != nil {
			return fmt.Errorf("failed to load config file %s: %s", f, err)
		}

		log.Infof("Loaded config from file: %s", f)
	}

	// Load config from env
	envSource := env.NewSource(
		// optionally specify prefix
		env.WithStrippedPrefix(cfgEnvPrefix),
	)
	if err := config.Load(envSource); err != nil {
		return fmt.Errorf("failed to load config from environment variables: %s", err)
	}

	log.Infof("Loaded config from environment variables with prefix of [%s]", cfgEnvPrefix)

	return nil
}
