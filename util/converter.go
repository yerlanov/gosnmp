package util

import (
	"fmt"
)

func TranslateIfOperStatus(in interface{}) string {
	status := in.(int)
	switch status {
	case 1:
		return "Up"
	case 2:
		return "Down"
	case 3:
		return "Testing"
	case 4:
		return "Unknown"
	case 5:
		return "Dormant"
	case 6:
		return "NotPresent"
	case 7:
		return "LowerLayerDown"
	default:
		return ""
	}
}

func ConvertDecimalToHexDecimal(in interface{}) string {
	var (
		macs     []string
		mac      string
		decimals = in.([]uint8)
	)

	for i, _ := range decimals {
		if decimals[i] == 0 {
			macs = append(macs, "00")
		} else {
			macs = append(macs, fmt.Sprintf("%X", decimals[i]))
		}
	}
	if len(macs) == 6 {
		mac = fmt.Sprintf("%s:%s:%s:%s:%s:%s", macs[0], macs[1], macs[2], macs[3], macs[4], macs[5])
	}
	return mac
}

func ConvertOctetStringToString(in interface{}) string {
	bytes := in.([]byte)
	return string(bytes)
}
