package main

import (
	"fmt"
	s "github.com/gosnmp/gosnmp"
	"github.com/hallidave/mibtool/smi"
	"github.com/kelseyhightower/envconfig"
	"log"
	"strconv"
	"test/util"
	"time"
)

type Conf struct {
	Port      string `envconfig:"PORT"`
	Community string `envconfig:"COMMUNITY"`
	MibDir    string `envconfig:"MIBDIR"`
	IpAddress string `envconfig:"IP"`
}

type Client struct {
	IpAddress string
	Port      string
}

type SnmpResponse struct {
	SysUpTime    string
	SysDescr     string
	ifOperStatus string
}

func main() {
	var conf Conf
	err := envconfig.Process("APP", &conf)
	if err != nil {
		panic(err)
	}

	conf.getSnmp()
}

func (c *Conf) getSnmp() {
	port, _ := strconv.ParseUint(c.Port, 10, 16)

	params := &s.GoSNMP{
		Target:    c.IpAddress,
		Port:      uint16(port),
		Community: c.Community,
		Version:   s.Version2c,
		Timeout:   time.Duration(2) * time.Second,
	}

	err := params.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}

	defer params.Conn.Close()

	mibs := []string{"SNMPv2-MIB::sysUpTime.0",
		"SNMPv2-MIB::sysDescr.0",
		"IF-MIB::ifOperStatus.9",
		"EtherLike-MIB::dot3StatsFCSErrors.9"}
	oids := c.convertMibToOid(mibs)
	result, err := params.Get(oids)
	if err != nil {
		log.Fatalf("error getting results: %v", err)
	}

	for _, variable := range result.Variables {
		switch variable.Type {
		case s.OctetString:
			bytes := variable.Value.([]byte)
			fmt.Println(string(bytes))
		case s.TimeTicks:
			fmt.Println(variable.Value)
		case s.Integer:
			res := util.TranslateIfOperStatus(variable.Value)
			fmt.Println(res)
		case s.Counter32:
			fmt.Println(variable.Value)
		default:
			fmt.Println(variable.Value)
		}
	}
}

func (c *Conf) convertMibToOid(mibs []string) []string {
	mib := smi.NewMIB(c.MibDir)
	mib.Debug = true
	var oids []string

	err := mib.LoadModules("IF-MIB", "SNMPv2-MIB", "EtherLike-MIB")
	if err != nil {
		log.Fatal(err)
	}

	mib.VisitSymbols(func(sym *smi.Symbol, oid smi.OID) {
		fmt.Printf("%-40s %s\n", sym, oid)
	})

	for _, v := range mibs {
		oid, err := mib.OID(v)
		if err != nil {
			log.Println(err)
		}
		oids = append(oids, oid.String())
	}
	return oids
}
