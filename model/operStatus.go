package model

import (
	"strconv"
	"test/util"
)

type OperStatus struct {
	UpTimeTkd        string
	UpTimeAgu        string
	ModelTkd         string
	ModelAgu         string
	StatusClientPort string
	ErrorClientPort  string
	MacAddressTkd    string
	MacAddressAgu    string
	SwitchStatus
}

type SwitchStatus struct {
	StatusAgu string
	StatusTkd string
}

func MapOperStatusToStruct(tkd []interface{}, agu []interface{}) OperStatus {
	var (
		upTimeTkd        string
		upTimeAgu        string
		modelTkd         string
		modelAgu         string
		statusClientPort string
		errorClientPort  string
		macAddressTkd    string
		macAddressAgu    string
		statusTkd        = "DOWN"
		statusAgu        = "DOWN"
	)

	if len(agu) == 3 {
		upTimeAgu = strconv.Itoa(int(agu[0].(uint32)))
		modelAgu = util.ConvertOctetStringToString(agu[1])
		macAddressAgu = util.ConvertDecimalToHexDecimal(agu[2])
		statusAgu = "UP"
	}

	if len(tkd) == 5 {
		upTimeTkd = strconv.Itoa(int(tkd[0].(uint32)))
		modelTkd = util.ConvertOctetStringToString(tkd[1])
		statusClientPort = util.TranslateIfOperStatus(tkd[2])
		errorClientPort = strconv.Itoa(int(tkd[3].(uint)))
		macAddressTkd = util.ConvertDecimalToHexDecimal(tkd[4])
		statusTkd = "UP"
	}

	switchStatus := SwitchStatus{
		StatusAgu: statusAgu,
		StatusTkd: statusTkd,
	}

	return OperStatus{
		UpTimeTkd:        upTimeTkd,
		ModelTkd:         modelTkd,
		StatusClientPort: statusClientPort,
		ErrorClientPort:  errorClientPort,
		MacAddressTkd:    macAddressTkd,
		UpTimeAgu:        upTimeAgu,
		ModelAgu:         modelAgu,
		MacAddressAgu:    macAddressAgu,
		SwitchStatus:     switchStatus,
	}
}
