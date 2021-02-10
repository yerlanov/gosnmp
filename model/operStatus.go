package model

import (
	"strconv"
	"test/util"
)

type OperStatus struct {
	UpTimeTkd        string `json:"uptime_tkd"`
	UpTimeAgu        string `json:"uptime_agu"`
	ModelTkd         string `json:"model_tkd"`
	ModelAgu         string `json:"model_agu"`
	StatusClientPort string `json:"status_client_port"`
	ErrorClientPort  string `json:"error_client_port"`
	MacAddressTkd    string `json:"mac_address_tkd"`
	MacAddressAgu    string `json:"mac_address_agu"`
	SwitchStatus
	Client
}

type SwitchStatus struct {
	StatusAgu string `json:"status_agu"`
	StatusTkd string `json:"status_tkd"`
}

func MapOperStatusToStruct(tkd []interface{}, agu []interface{}, client Client) OperStatus {
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
		Client:           client,
	}
}
