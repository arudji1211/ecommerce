package config

import (
	"flag"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

const (
	ENV_LOCAL      = "local"
	ENV_PRODUCTION = "production"
)

var (
	Load       structure
	searchPath = []string{
		".",
		"./config",
	}
)

type (
	structure struct {
		Server     server     `mapstructure:"server"`
		DataSource dataSource `mapstructure:"dataSource"`
	}

	server struct {
		Name string `mapstructure:"name"`
		Env  string `mapstructure:"env"`
		Http struct {
			port uint `mapstructure:"port"`
		} `mapstrcture:"http"`
	}

	dataSource struct {
		Migrate  bool `mapstructure:"migrate"`
		Postgres struct {
			Master PostgresConfig `mapstructure:"master"`
		}
	}
)

func init() {
	var configName = flag.String("configName", "local", "config name for service, default local")
	v := viper.New()
	if err := initialiseFileAndEnv(v, *configName); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warning("No config file found on search path")
		} else {
			log.Fatalf("Error occurred during loading config: %s", err.Error())
		}
		panic(err)
	}

	if err := v.Unmarshal(&Load); err != nil {
		log.Fatal("cannot unmarshal config")
		panic(err)
	}
}

func initialiseFileAndEnv(v *viper.Viper, configName string) error {
	v.SetConfigName(configName)
	for _, path := range searchPath {
		v.AddConfigPath(path)
	}
	return v.ReadInConfig()
}
