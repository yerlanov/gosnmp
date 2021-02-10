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
		conf                     = config.New()
		operStatus               model.OperStatus
		responseAgu, responseTkd []interface{}
	)

	client, err := model.GetClientByLogin(login)
	if err != nil {
		return operStatus, exception.Error{ErrorType: exception.NotFound, ErrorMessage: err}
	}
	fmt.Println(client)

	pingAgu := pingSwitch(client.IpAgu)
	pingTkd := pingSwitch(client.IpTkd)
	fmt.Println(pingAgu)
	fmt.Println(pingTkd)

	if pingAgu == "UP" {
		textAguOids := []string{"SNMPv2-MIB::sysUpTime.0",
			"SNMPv2-MIB::sysDescr.0",
			"LLDP-MIB::lldpLocChassisId.0"}

		AguOids, err := convertTextOidToOid(conf.SnmpMibDir, textAguOids)
		if err.ErrorMessage != nil {
			return operStatus, err
		}

		responseAgu, err = snmpRequest(conf, client.IpAgu, AguOids)
		if err.ErrorMessage != nil {
			return operStatus, err
		}
	}

	if pingTkd == "UP" {
		textTkdOids := []string{"SNMPv2-MIB::sysUpTime.0",
			"SNMPv2-MIB::sysDescr.0",
			fmt.Sprintf("IF-MIB::ifOperStatus.%s", client.Port),
			fmt.Sprintf("EtherLike-MIB::dot3StatsFCSErrors.%s", client.Port),
			"LLDP-MIB::lldpLocChassisId.0"}

		tkdOids, err := convertTextOidToOid(conf.SnmpMibDir, textTkdOids)
		if err.ErrorMessage != nil {
			return operStatus, err
		}

		responseTkd, err = snmpRequest(conf, client.IpTkd, tkdOids)
		if err.ErrorMessage != nil {
			return operStatus, err
		}
	}

	operStatus = model.MapOperStatusToStruct(responseTkd, responseAgu, client)

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
	status, pingErr := ping.NewPinger(ip)
	if pingErr != nil {
		fmt.Println(pingErr)
		return "DOWN"
	}
	fmt.Println(status)
	return "UP"
}
