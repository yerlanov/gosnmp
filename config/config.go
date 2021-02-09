package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	SnmpConfig
	DatabaseConfig
}

type SnmpConfig struct {
	SnmpPort      string `envconfig:"PORT"`
	SnmpCommunity string `envconfig:"COMMUNITY"`
	SnmpMibDir    string `envconfig:"MIBDIR"`
}

type DatabaseConfig struct {
	DbHost     string `envconfig:"DB_HOST"`
	DbPort     string `envconfig:"DB_PORT"`
	DbUser     string `envconfig:"DB_USER"`
	DbPassword string `envconfig:"DB_PASSWORD"`
	DbDatabase string `envconfig:"DB_DATABASE"`
}

func New() Config {
	var conf Config
	err := envconfig.Process("APP", &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
