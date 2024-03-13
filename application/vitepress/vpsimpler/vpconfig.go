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

var ConfigNames = []string{"config.mts", "config.ts", "config.js"}
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
	template := "{\n  title: \"{title}\",\n  description: \"{description}\",\n  srcDir: \"./docs\",\n  themeConfig: {\n    //导航\n    nav: [\n      {\n        text: \"Home\",\n        link: \"/\"\n      },\n      {\n        text: \"About\",\n        link: \"/about\"\n      },\n      {\n        text: \"Contact\",\n        link: \"/contact\"\n      }\n    ],\n    //侧栏\n    sidebar: [\n      {\n        text: \"Examples\",\n        items: [\n          { text: \"Markdown Examples\", link: \"/markdown-examples\" },\n          { text: \"Runtime API Examples\", link: \"/api-examples\" }\n        ]\n      }\n    ],\n    //社交账号\n    socialLinks: [\n      { icon: \"github\", link: \"https://github.com/vuejs/vitepress\" }\n    ]\n  }\n}"
	newContent := strings.ReplaceAll(template, "{title}", title)
	newContent = strings.ReplaceAll(newContent, "{description}", description)
	return newContent
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
}
