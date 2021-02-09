package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port string `envconfig:"APP_PORT" required:"true"`
	SnmpConfig
	DatabaseConfig
}

type SnmpConfig struct {
	SnmpPort      string `envconfig:"SNMP_PORT"`
	SnmpCommunity string `envconfig:"SNMP_COMMUNITY"`
	SnmpMibDir    string `envconfig:"SNMP_MIBDIR"`
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
	err := envconfig.Process("APPs", &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
