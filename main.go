package main

import (
	"test/config"
	"test/service"
)

type SnmpResponse struct {
	SysUpTime    string
	SysDescr     string
	ifOperStatus string
}

func main() {
	_ = config.New()

	service.GetOperStatus()
}
