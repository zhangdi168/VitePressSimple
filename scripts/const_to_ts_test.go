// Package script
// @Author: zhangdi
// @File: const_to_ts
// @Version: 1.0.0
// @Date: 2023/8/31 13:55
package scripts

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"wailstemplate/application/pkg/utils"
)

// 常量生成前端常量
func TestCore(t *testing.T) {
	GenFrontConstFile()
}

// GenFrontConstFile 同步生成ts常量文件
func GenFrontConstFile() {
	err := filepath.Walk(GetGoDirConst(), GenTsConstHandler)
	if err != nil {
		panic(err)
	}
}

// GetGoDirConst api常量定义文件，该文件路径固定
func GetGoDirConst() string {
	return filepath.Join(ProjectDir(), "application", "constant")
}

// GetFrontConstDir go常量默认生成路径
func GetFrontConstDir() string {
	return filepath.Join(FrontDir(), "src", "constant")
}

// GetFrontApisConst api常量 前段存储路径
func GetFrontApisConst() string {
	return filepath.Join(FrontDir(), "src", "apis")
}

// GenTsConstHandler 常量生成
func GenTsConstHandler(path string, info os.FileInfo, err error) error {
	tsDir := GetFrontConstDir()
	goDir := GetGoDirConst()
	if strings.Contains(path, "apis") {
		tsDir = GetFrontApisConst()
	}
	if err != nil {
		return err
	}
	if !strings.HasSuffix(info.Name(), ".go") {
		return nil // 忽略非 .go 文件
	}
	if utils.ArrContains(skinGenConstants, filepath.Base(path)) {
		fmt.Println("跳过：" + filepath.Base(path))
		return nil
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var constants []string
	var Values string
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "const ") { // 处理 const 声明
			fields := strings.Fields(line)
			if len(fields) < 3 {
				continue
			}
			name := fields[1]
			value := strings.Join(fields[2:], " ")
			//value = strings.ReplaceAll(value, `"`, ``)
			value = strings.ReplaceAll(value, `=`, ``)
			value = strings.TrimSpace(value)
			Values += value + ","
			constants = append(constants, fmt.Sprintf("export const %s = %s\n", name, value))
		}
	}
	fileName := filepath.Base(path)
	fileName = strings.ReplaceAll(fileName, ".go", "")
	Values = strings.TrimRight(Values, ",")
	//导出
	//export  const BoolArray = [BoolYes, BoolNo]; // 使用数组字面量
	//export const BoolObj= {BoolNo,BoolYes};
	constants = append(constants, fmt.Sprintf("export  const %sArray = [%s];\n", utils.VarLineToCamelCase(fileName), Values))
	//	constants = append(constants, fmt.Sprintf("export  const %sObj = {%s};\n", utils.VarLineToCamelCase(fileName), Names))

	outFile := path[len(goDir)+1:len(path)-3] + ".ts"
	outPath := filepath.Join(tsDir, outFile)
	dir := filepath.Dir(outPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) { // 如果目录不存在，则创建目录
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	outContent := strings.Join(constants, "\n")
	if err := ioutil.WriteFile(outPath, []byte(outContent), 0644); err != nil {
		return err
	}
	fmt.Printf("generate constant file %s\n", outPath)
	return nil
}
