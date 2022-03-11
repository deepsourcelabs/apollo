package utils

import (
	"log"
	"path/filepath"
	"time"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

type Config struct {
	Redis   redisConfig   `koanf:"redis"`
	Server  serverConfig  `koanf:"server"`
	Logging loggingConfig `koanf:"logging"`
	Timeout timeoutConfig `koanf:"timeout"`
}

type redisConfig struct {
	Addr     string `koanf:"addr"`
	Password string `koanf:"password"`
	DB       int    `koanf:"db"`
}

type serverConfig struct {
	Addr string `koanf:"addr"`
}

type loggingConfig struct {
	File string `koanf:"file"`
}

type timeoutConfig struct {
	Value int    `koanf:"value"`
	Type  string `koanf:"type"`
}

// GetConfig returns a populated Config.
func GetConfig(filePath string) *Config {
	// create a new koanf instance
	k := koanf.New(".")

	// use a YAML parser
	parser := yaml.Parser()

	// get absolute path from filePath
	configPath, err := filepath.Abs(filePath)
	if err != nil {
		log.Fatalln("failed to read config path")
	}

	// load configs using a file provider
	err = k.Load(file.Provider(configPath), parser)
	if err != nil {
		log.Fatalln("failed to read config file:", err)
	}

	// unmarshal redis configs
	rc := &redisConfig{}
	err = k.Unmarshal("redis", rc)
	if err != nil {
		log.Fatalln("failed to unmarshal redis config:", err)
	}

	// unmarshal server configs
	sc := &serverConfig{}
	err = k.Unmarshal("server", sc)
	if err != nil {
		log.Fatalln("failed to unmarshal server config:", err)
	}

	// unmarshal logging configs
	lc := &loggingConfig{}
	err = k.Unmarshal("logging", lc)
	if err != nil {
		log.Fatalln("failed to unmarshal logging config:", err)
	}

	// unmarshal timeout configs
	tc := &timeoutConfig{}
	err = k.Unmarshal("timeout", tc)
	if err != nil {
		log.Fatalln("failed to unmarshal timeout config:", err)
	}

	// populate config
	conf := &Config{
		Redis:   *rc,
		Server:  *sc,
		Logging: *lc,
		Timeout: *tc,
	}

	return conf
}

// GetTimeout returns a time duration using configs.
func GetTimeout(conf *Config) time.Duration {
	switch conf.Timeout.Type {
	case "minute":
		return time.Minute * time.Duration(conf.Timeout.Value)
	case "second":
		return time.Second * time.Duration(conf.Timeout.Value)
	case "microsecond":
		return time.Microsecond * time.Duration(conf.Timeout.Value)
	}

	return time.Minute
}
