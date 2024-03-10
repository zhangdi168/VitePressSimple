// Package gencore
// @Author: zhangdi
// @File: func
// @Version: 1.0.0
// @Date: 2023/8/30 18:19
package varname

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// ReplaceAndWriteTemplate 传入一个tmpl模板文件路径，一个占位符文本，一个真实值文本，输出文件路径四个参数。
// 实现读取tmp模板文件的值并将占位符本替换成传入的真实值，将最终替换好的文本写入到传入的输出文件路径中。
func ReplaceAndWriteTemplate(tmplFilePath string, outputFilePath string, placeholderMap map[string]string) error {

	//判断路径outputFilePath所在目录是否存在，不存在则先创建目录
	outputFileDir := filepath.Dir(outputFilePath)
	_, err := os.Stat(outputFileDir)
	if os.IsNotExist(err) {
		// 目录不存在，创建目录
		err := os.MkdirAll(outputFileDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	tmplContent, err := ioutil.ReadFile(tmplFilePath)
	if err != nil {
		return err
	}

	content := string(tmplContent)
	for placeholder, value := range placeholderMap {
		processedContent := strings.Replace(content, placeholder, value, -1)
		content = processedContent
	}

	err = ioutil.WriteFile(outputFilePath, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}

// GetDefaultRootDir 获取项目默认的根目录
func GetDefaultRootDir() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	rootDir := filepath.Join(dir, "../../..")
	return rootDir
}

// FileExists 判断一个文件是否存在
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// FormatGoImportStr 组装import字符串
func FormatGoImportStr(LibNames []string) string {
	if len(LibNames) <= 0 {
		return ""
	}
	if len(LibNames) == 1 {
		return fmt.Sprintf("import \"%s\"", LibNames[0])
	}

	importStr := fmt.Sprintf("import ( \n")
	for _, name := range LibNames {
		importStr += fmt.Sprintf("\"%s\" \n", name)
	}

	return importStr + ")"
}

// NameToLine 下划线命名
func NameToLine(camel string) string {
	var snake []rune
	for i, c := range camel {
		if unicode.IsUpper(c) {
			if i > 0 && (i+1 < len(camel) && unicode.IsLower(rune(camel[i+1])) || unicode.IsLower(rune(camel[i-1]))) {
				snake = append(snake, '_')
			}
			snake = append(snake, unicode.ToLower(c))
		} else {
			snake = append(snake, c)
		}
	}
	return string(snake)
}

// NameToCameSmall 小驼峰命名
func NameToCameSmall(s string) string {
	if strings.Contains(s, "_") { //如果原来是下划线命名法
		words := strings.Split(s, "_")
		for i := 0; i < len(words); i++ {
			words[i] = strings.ToLower(words[i])
			if i > 0 {
				words[i] = strings.Title(words[i])
			}
		}
		return strings.Join(words, "")
	}

	var words []string
	for i := 0; i < len(s); i++ {
		if i > 0 && (unicode.IsUpper(rune(s[i])) || !(unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i])))) {
			// 如果当前字符是大写字母或者非字母数字字符，则将前面的部分作为一个单词
			words = append(words, s[:i])
			s = s[i:]
			i = 0
		}
	}
	words = append(words, s)
	for i := 0; i < len(words); i++ {
		if i == 0 {
			words[i] = strings.ToLower(words[i])
		} else {
			words[i] = strings.Title(words[i])
		}
	}

	return strings.Join(words, "")
}

// NameToCameBig 大驼峰命名法
func NameToCameBig(line string) string {
	var camel []rune
	words := strings.Split(line, "_")

	for _, word := range words {
		firstChar := word[0]
		restChars := word[1:]

		camel = append(camel, unicode.ToUpper(rune(firstChar)))
		camel = append(camel, []rune(restChars)...)
	}

	return string(camel)
}

func ParseTagToMap(tagStr string) map[string]string {
	// 创建一个map来存储key-value对
	data := make(map[string]string)
	// 根据分号进行切割
	pairs := strings.Split(tagStr, ";")

	// 遍历切割后的字符串对
	for _, pair := range pairs {
		// 再根据冒号进行切割
		kv := strings.Split(pair, ":")
		if len(kv) != 2 {
			continue
		}
		// 存储到map中
		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])
		data[key] = value
	}

	// 打印map中的key-value对
	for k, v := range data {
		fmt.Println(k, ":", v)
	}
	return data
}

func GetMapVal(data map[string]string, k string) string {
	// 使用访问元素的语法获取给定key的值
	value, exists := data[k]
	if exists {
		return value
	}
	// 如果key不存在，可以根据需求返回一个默认值或者错误信息等
	return ""
}

// GetMapValDefault 获取同一个map.key 传入默认值
func GetMapValDefault(data map[string]string, k string, defaultValue string) string {
	// 使用访问元素的语法获取给定key的值
	value, exists := data[k]
	if exists {
		if value == "" {
			return defaultValue
		}
		return value
	}

	// 如果key不存在，可以根据需求返回一个默认值或者错误信息等
	return defaultValue
}
