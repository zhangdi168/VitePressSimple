// Package utils
// @Author: zhangdi
// @File: var
// @Version: 1.0.0
// @Date: 2023/5/29 18:32
package utils

import "strings"

// VarLineToCamelCase 将下划线命名转成大驼峰命名
func VarLineToCamelCase(s string) string {
	var result string

	// 分隔字符串并遍历每个部分
	parts := strings.Split(s, "_")
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}

		// 将部分的第一个字符大写，其余字符保持不变
		firstChar := strings.ToUpper(part[0:1])
		restChars := part[1:]

		// 拼接所有部分
		result += firstChar + restChars
	}

	return result
}
