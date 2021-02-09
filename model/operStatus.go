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
}

func MapOperStatusToStruct(tkd []interface{}, agu []interface{}) OperStatus {
	return OperStatus{
		UpTimeTkd:        strconv.Itoa(int(tkd[0].(uint32))),
		ModelTkd:         util.ConvertOctetStringToString(tkd[1]),
		StatusClientPort: util.TranslateIfOperStatus(tkd[2]),
		ErrorClientPort:  strconv.Itoa(int(tkd[3].(uint))),
		MacAddressTkd:    util.ConvertDecimalToHexDecimal(tkd[4]),
		UpTimeAgu:        strconv.Itoa(int(agu[0].(uint32))),
		ModelAgu:         util.ConvertOctetStringToString(agu[1]),
		MacAddressAgu:    util.ConvertDecimalToHexDecimal(agu[2]),
	}
}
