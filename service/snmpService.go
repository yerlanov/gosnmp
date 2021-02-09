package service

import (
	"fmt"
	s "github.com/gosnmp/gosnmp"
	"github.com/hallidave/mibtool/smi"
	"strconv"
	"test/config"
	"test/model"
	"time"
)

func GetOperStatusService(login string) (model.OperStatus, error) {
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

		operStatus model.OperStatus
	)

	tkdOids, err := convertTextOidToOid(conf.SnmpMibDir, textTkdOids)
	if err != nil {
		return operStatus, err
	}
	responseTkd, err := snmpRequest(conf, "172.16.95.193", tkdOids)
	if err != nil {
		return operStatus, err
	}

	AguOids, err := convertTextOidToOid(conf.SnmpMibDir, textAguOids)
	if err != nil {
		return operStatus, err
	}

	responseAgu, err := snmpRequest(conf, "172.16.95.193", AguOids)
	if err != nil {
		return operStatus, err
	}

	operStatus = model.MapOperStatusToStruct(responseTkd, responseAgu)

	return operStatus, nil
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

func convertTextOidToOid(dir string, textOids []string) ([]string, error) {
	mib := smi.NewMIB(dir)
	mib.Debug = true
	var oids []string

	err := mib.LoadModules("IF-MIB", "SNMPv2-MIB", "EtherLike-MIB", "LLDP-MIB")
	if err != nil {
		return oids, err
	}

	mib.VisitSymbols(func(sym *smi.Symbol, oid smi.OID) {
		fmt.Printf("%-40s %s\n", sym, oid)
	})

	for _, v := range textOids {
		oid, err := mib.OID(v)
		if err != nil {
			return oids, err
		}
		oids = append(oids, oid.String())
	}
	return oids, nil
}
