package model

import (
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
}

func MapOperStatusToStruct(tkd []interface{}, agu []interface{}) OperStatus {
	return OperStatus{
		UpTimeTkd:        tkd[0].(string),
		ModelTkd:         util.ConvertOctetStringToString(tkd[1]),
		StatusClientPort: util.TranslateIfOperStatus(tkd[2]),
		ErrorClientPort:  tkd[3].(string),
		MacAddressTkd:    util.ConvertDecimalToHexDecimal(tkd[4]),
		UpTimeAgu:        agu[0].(string),
		ModelAgu:         agu[1].(string),
		MacAddressAgu:    util.ConvertDecimalToHexDecimal(agu[3]),
	}
}
