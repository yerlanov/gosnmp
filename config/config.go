package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port string `envconfig:"FTTB_APP_PORT" required:"true"`
	SnmpConfig
	DatabaseConfig
}

type SnmpConfig struct {
	SnmpPort      string `envconfig:"FTTB_SNMP_PORT"`
	SnmpCommunity string `envconfig:"FTTB_SNMP_COMMUNITY"`
	SnmpMibDir    string `envconfig:"FTTB_SNMP_MIBDIR"`
}

type DatabaseConfig struct {
	DbHost     string `envconfig:"FTTB_DB_HOST"`
	DbPort     string `envconfig:"FTTB_DB_PORT"`
	DbUser     string `envconfig:"FTTB_DB_USER"`
	DbPassword string `envconfig:"FTTB_DB_PASSWORD"`
	DbDatabase string `envconfig:"FTTB_DB_DATABASE"`
}

func New() Config {
	var conf Config
	err := envconfig.Process("APPs", &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
