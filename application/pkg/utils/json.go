package utils

import "encoding/json"

func DataToJsonStr(data any) string {
	str, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(str)
}
