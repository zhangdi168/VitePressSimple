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
// startString 开始查找字符串，先匹配到改字符串，然后对该字符串后面的内容进行查找，传入空字符串则全局匹配
// GetTsDataContent(s, "[", "]", "export")表示export字符串之前的字符串不参与计算，从export字符串开始，匹配到{和}之间的内容
func GetTsDataContent(fullContent string, startMatchChar string, endMatchChar string, startString string) (string, error) {
	startStrIndex := 0 //开始处理的起始字符
	if startString != "" {
		startStrIndex = strings.Index(fullContent, startString)
		if startStrIndex == -1 {
			startStrIndex = 0
		}
	}

	content := fullContent[startStrIndex:]                     //取出要处理的字符串
	start := strings.Index(content, startMatchChar)            //第一个{位置
	end := strings.Index(ReverseString(content), endMatchChar) //逆序字符串后 第一个}位置
	if start == -1 {
		return "", fmt.Errorf("no  found " + startMatchChar)
	}
	if end == -1 {
		return "", fmt.Errorf("no  found" + endMatchChar)
	}
	end = len(content) - end
	if start > end {
		return "", fmt.Errorf("start > end")
	}

	//截取{和}之间的内容
	return content[start:end], nil
}
