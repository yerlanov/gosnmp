package service

import (
	"fmt"
	s "github.com/gosnmp/gosnmp"
	"github.com/hallidave/mibtool/smi"
	"log"
	"strconv"
	"test/config"
	"test/model"
	"time"
)

func GetOperStatus() {
	var (
		conf        = config.New()
		textTkdOids = []string{"SNMPv2-MIB::sysUpTime.0",
			"SNMPv2-MIB::sysDescr.0",
			"IF-MIB::ifOperStatus.9",
			"EtherLike-MIB::dot3StatsFCSErrors.9",
			"LLDP-MIB::lldpLocChassisId.0"}
		textAguOids = []string{"SNMPv2-MIB::sysUpTime.0",
			"SNMPv2-MIB::sysDescr.0",
			"LLDP-MIB::lldpLocChassisId.0"}
	)
	tkdOids := convertMibToOid(conf.SnmpMibDir, textTkdOids)
	responseTkd, err := snmpRequest(conf, "172.16.95.193", tkdOids)
	if err != nil {
		fmt.Println(err)
		return
	}

	AguOids := convertMibToOid(conf.SnmpMibDir, textAguOids)
	responseAgu, err := snmpRequest(conf, "172.16.95.193", AguOids)
	if err != nil {
		fmt.Println(err)
		return
	}

	operStatus := model.MapOperStatusToStruct(responseTkd, responseAgu)

	fmt.Println(operStatus)
}

func snmpRequest(conf config.Config, ip string, oids []string) ([]interface{}, error) {
	var (
		response []interface{}
	)

	port, err := strconv.ParseUint(conf.SnmpPort, 10, 16)
	if err != nil {
		return response, err
	}

	params := &s.GoSNMP{
		Target:    ip,
		Port:      uint16(port),
		Community: conf.SnmpCommunity,
		Version:   s.Version2c,
		Timeout:   time.Duration(2) * time.Second,
	}

	err = params.Connect()
	if err != nil {
		return response, fmt.Errorf("Connect() err: %v", err)
	}

	defer params.Conn.Close()

	result, err := params.Get(oids)
	if err != nil {
		return response, fmt.Errorf("Connect() err: %v", err)
	}

	for _, variable := range result.Variables {
		response = append(response, variable.Value)
	}
	return response, nil
}

func convertMibToOid(dir string, textOids []string) []string {
	mib := smi.NewMIB(dir)
	mib.Debug = true
	var oids []string

	err := mib.LoadModules("IF-MIB", "SNMPv2-MIB", "EtherLike-MIB", "LLDP-MIB")
	if err != nil {
		log.Fatal(err)
	}

	mib.VisitSymbols(func(sym *smi.Symbol, oid smi.OID) {
		fmt.Printf("%-40s %s\n", sym, oid)
	})

	for _, v := range textOids {
		oid, err := mib.OID(v)
		if err != nil {
			log.Println(err)
		}
		oids = append(oids, oid.String())
	}
	return oids
}
