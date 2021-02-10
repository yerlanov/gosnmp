package service

import (
	"fmt"
	"github.com/go-ping/ping"
	s "github.com/gosnmp/gosnmp"
	"github.com/hallidave/mibtool/smi"
	"strconv"
	"test/config"
	"test/exception"
	"test/model"
	"time"
)

func GetOperStatusService(login string) (model.OperStatus, exception.Error) {
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

		operStatus               model.OperStatus
		responseAgu, responseTkd []interface{}
	)

	client, err := model.GetClientByLogin("0013328282")
	if err != nil {
		return operStatus, exception.Error{ErrorType: exception.NOTFOUND, ErrorMessage: err}
	}

	fmt.Println(client)

	ipTkd := "172.16.68.200"
	ipAgu := "172.16.95.193"

	pingAgu := pingSwitch(ipTkd)
	pingTkd := pingSwitch(ipAgu)

	if pingAgu == "UP" {
		AguOids, err := convertTextOidToOid(conf.SnmpMibDir, textAguOids)
		if err.ErrorMessage != nil {
			return operStatus, err
		}

		responseAgu, err = snmpRequest(conf, ipAgu, AguOids)
		if err.ErrorMessage != nil {
			return operStatus, err
		}
	}

	if pingTkd == "UP" {
		tkdOids, err := convertTextOidToOid(conf.SnmpMibDir, textTkdOids)
		if err.ErrorMessage != nil {
			return operStatus, err
		}

		responseTkd, err = snmpRequest(conf, ipTkd, tkdOids)
		if err.ErrorMessage != nil {
			return operStatus, err
		}
	}

	operStatus = model.MapOperStatusToStruct(responseTkd, responseAgu)

	return operStatus, exception.Error{ErrorType: "", ErrorMessage: nil}
}

func snmpRequest(conf config.Config, ip string, oids []string) ([]interface{}, exception.Error) {
	var (
		response []interface{}
	)

	port, err := strconv.ParseUint(conf.SnmpPort, 10, 16)
	if err != nil {
		return response, exception.Error{ErrorType: exception.ParseError, ErrorMessage: err}
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
		return response, exception.Error{ErrorType: exception.SnmpConnectError, ErrorMessage: err}
	}

	defer params.Conn.Close()

	result, err := params.Get(oids)
	if err != nil {
		return response, exception.Error{ErrorType: exception.SnmpGetError, ErrorMessage: err}
	}

	for _, variable := range result.Variables {
		response = append(response, variable.Value)
	}
	return response, exception.Error{ErrorType: "", ErrorMessage: nil}
}

func convertTextOidToOid(dir string, textOids []string) ([]string, exception.Error) {
	mib := smi.NewMIB(dir)
	mib.Debug = true
	var oids []string

	err := mib.LoadModules("IF-MIB", "SNMPv2-MIB", "EtherLike-MIB", "LLDP-MIB")
	if err != nil {
		return oids, exception.Error{ErrorType: exception.OidConvertError, ErrorMessage: err}
	}

	mib.VisitSymbols(func(sym *smi.Symbol, oid smi.OID) {
		fmt.Printf("%-40s %s\n", sym, oid)
	})

	for _, v := range textOids {
		oid, err := mib.OID(v)
		if err != nil {
			return oids, exception.Error{ErrorType: exception.OidConvertError, ErrorMessage: err}
		}
		oids = append(oids, oid.String())
	}
	return oids, exception.Error{ErrorType: "", ErrorMessage: nil}
}

func pingSwitch(ip string) string {
	_, pingErr := ping.NewPinger(ip)
	if pingErr != nil {
		return "DOWN"
	}
	return "UP"
}
