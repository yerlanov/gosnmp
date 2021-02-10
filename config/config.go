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
	SnmpPort      string `envconfig:"FTTB_SNMP_PORT" required:"true"`
	SnmpCommunity string `envconfig:"FTTB_SNMP_COMMUNITY" required:"true"`
	SnmpMibDir    string `envconfig:"FTTB_SNMP_MIBDIR" required:"true"`
}

type DatabaseConfig struct {
	DbHost     string `envconfig:"FTTB_DB_HOST" required:"true"`
	DbPort     string `envconfig:"FTTB_DB_PORT" required:"true"`
	DbUser     string `envconfig:"FTTB_DB_USER" required:"true"`
	DbPassword string `envconfig:"FTTB_DB_PASSWORD" required:"true"`
	DbDatabase string `envconfig:"FTTB_DB_DATABASE" required:"true"`
}

func New() Config {
	var conf Config
	err := envconfig.Process("APPs", &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
