package tsparser

import (
	"fmt"
	"strings"
)

// ReverseString 反转字符串
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// GetTsDataContent 提取第一个{和}之间的内容，返回包括startChar和endChar的内容 如传入[和]返回[2121,12]
// startChar 开始字符如 { [
// endChar 结束字符如 } ]
func GetTsDataContent(s string, startChar string, endChar string) (string, error) {
	start := strings.IndexAny(s, startChar)            //第一个{位置
	end := strings.IndexAny(ReverseString(s), endChar) //逆序字符串后 第一个}位置
	if start == -1 {
		return "", fmt.Errorf("no  found {" + startChar)
	}
	if end == -1 {
		return "", fmt.Errorf("no  found" + endChar)
	}
	end = len(s) - end
	if start > end {
		return "", fmt.Errorf("start > end")
	}
	if end+1 < len(s) {
		end += 1
	}

	//截取{和}之间的内容
	return s[start:end], nil
}
