package vpsimpler

import (
	"fmt"
	"path/filepath"
	"strings"
	"wailstemplate/application/constant/keys"
	"wailstemplate/application/pkg/cfg"
	"wailstemplate/application/pkg/filehelper"
	"wailstemplate/application/pkg/mylog"
	"wailstemplate/application/vitepress/tsparser"
)

var ConfigNames = []string{"config.mts", "config.ts", "config.js", "index.js", "index.ts", "index.mts"}
var defaultFileName = "config.mts"

// VpConfig 操作配置文件相关的类
type VpConfig struct {
}

func NewVpConfig() *VpConfig {
	return &VpConfig{}
}

// GetVpConfigData  获取vitepress配置文件内容
func (s *VpConfig) GetVpConfigData() string {
	path, err := s.GetConfigPath()
	if err != nil {
		mylog.Error("获取vitepress配置文件路径异常" + err.Error())
		return ""
	}
	content, err := filehelper.ReadContent(path)
	if err != nil {
		mylog.Error("获取vitepress配置文件内容异常" + err.Error())
		return ""
	}

	dataContent, err := tsparser.GetTsDataContent(content, "{", "}", "export")
	if err != nil {
		mylog.Error("解析vitepress配置文件内容异常" + err.Error())
		return dataContent
	}
	return dataContent
}
func (s *VpConfig) filterContentString(content string) string {
	//删除掉导入的花括号
	content = strings.ReplaceAll(content, "{ defineConfig }", "")
	content = strings.ReplaceAll(content, "{defineConfig}", "")
	content = strings.ReplaceAll(content, "{ defineConfig}", "")
	content = strings.ReplaceAll(content, "{defineConfig }", "")
	return content
}

// SaveConfig 保存配置文件
func (s *VpConfig) SaveConfig(configContentJson string) string {
	configPath, err := s.GetConfigPath()
	if err != nil {
		return err.Error()
	}
	err = s.RenameConfigName(configPath)
	if err != nil {
		return "重命名配置文件config.mts异常：" + err.Error()
	}
	newContent := "import { defineConfig } from \"vitepress\";\n\texport default defineConfig(" + configContentJson + ");\n"
	err = filehelper.WriteContent(s.getDefaultConfigPath(), newContent)
	if err != nil {
		return err.Error()
	}
	return ""
}

// BackupOriginConfig 备份原来的配置文件
//func (s *VpConfig) BackupOriginConfig() {
//	filehelper.CopyFolder()
//}

// RenameConfigName 重命名配置文件的后缀名称 默认使用config.mts后缀
func (s *VpConfig) RenameConfigName(currPath string) error {
	//名字一样则不修改
	if filepath.Base(currPath) == defaultFileName {
		return nil
	}
	//名字不一样则修改
	err := filehelper.RenamePath(currPath, s.getDefaultConfigPath())
	if err != nil {
		return err
	}
	return nil
}

func (s *VpConfig) getDefaultConfigPath() string {
	baseDir := cfg.GetString(keys.ConfigKeyProjectDir)
	return filepath.Join(baseDir, ".vitepress", defaultFileName)
}

// ReplaceDefaultConfigContent 获取默认的属性信息
func (s *VpConfig) ReplaceDefaultConfigContent(title string, description string) string {
	template := "{\n  title: \"{title}\",\n  description: \"{description}\",\n  srcDir: \"./docs\",\n  themeConfig: {\n    \n    nav: [\n      {\n        text: \"Home\",\n        link: \"/\"\n      },\n      {\n        text: \"About\",\n        link: \"/about\"\n      },\n      {\n        text: \"Contact\",\n        link: \"/contact\"\n      }\n    ],\n    \n    sidebar: [\n      {\n        text: \"Examples\",\n        items: [\n          { text: \"Markdown Examples\", link: \"/markdown-examples\" },\n          { text: \"Runtime API Examples\", link: \"/api-examples\" }\n        ]\n      }\n    ],\n   \n    socialLinks: [\n      { icon: \"github\", link: \"https://github.com/vuejs/vitepress\" }\n    ]\n  }\n}"
	newContent := strings.ReplaceAll(template, "{title}", title)
	newContent = strings.ReplaceAll(newContent, "{description}", description)
	return newContent
}

// ExistsVitepressDir 判断是否存在.vitepress目录
func (s *VpConfig) ExistsVitepressDir(dir string) bool {
	path := filepath.Join(dir, ".vitepress")
	return filehelper.FileExists(path)
}
func (s *VpConfig) GetConfigPath() (string, error) {
	baseDir := cfg.GetString(keys.ConfigKeyProjectDir)
	if baseDir == "" {
		return "", fmt.Errorf("未配置当前项目路径，请先配置")
	}
	for _, name := range ConfigNames {
		configPath := filepath.Join(baseDir, ".vitepress", name)
		if filehelper.FileExists(configPath) {
			return configPath, nil
		}
	}
	return "", fmt.Errorf("配置文件不存在config.js/config.ts/config.mts")
	//path := filepath.Join(baseDir, ".vitepress", defaultFileName)
	//if !filehelper.FileExists(path) { //文件不存在会先创建
	//	filehelper.CreateFile(path)
	//}
	//return filepath.Join(baseDir, ".vitepress", defaultFileName), nil
}