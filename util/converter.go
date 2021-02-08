package util

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
	}
	return ""
}
