package utils

func GetStringFromMap(m map[string]interface{}, key string) (string, bool) {
	if m == nil {
		return "", false
	}
	value, ok := m[key]
	if !ok {
		return "", false // 返回空字符串和 false 表示 key 不存在于 map 中
	}
	strValue, isString := value.(string)
	if !isString {
		return "", false // 返回空字符串和 false 表示值不是字符串类型
	}

	return strValue, true // 返回字符串值和 true 表示获取成功
}
