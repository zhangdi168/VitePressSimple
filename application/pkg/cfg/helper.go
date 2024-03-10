// Package cfg
// @Author: zhangdi
// @File: helper
// @Version: 1.0.0
// @Date: 2023/11/22 12:21
package cfg

import "path/filepath"

var instance *Config

const DefaultYamlPath = "configs.yaml"

// InitCfg 初始化
func InitCfg(yamlPath string) error {
	var err error
	instance, err = NewConfig()
	if err != nil {
		return err
	}
	if yamlPath == "" {
		yamlPath = DefaultYamlPath
	}
	err = instance.LoadConfig(yamlPath)
	if err != nil {
		return err
	}
	return nil
}

func checkInstance() {
	if instance == nil {
		err := InitCfg(filepath.Join(DefaultYamlPath))
		if err != nil {
			println(err.Error())
			return
		}
	}
}

func Set(k string, v any) error {
	checkInstance()
	err := instance.Set(k, v)
	return err
}

// SetDefault 设置初始值
func SetDefault(k string, v any) error {
	checkInstance()
	val := Get(k)
	if val == nil || val == "" {
		err := instance.Set(k, v)
		return err
	}
	return nil
}

func Get(k string) any {
	checkInstance()
	val, err := instance.Get(k)
	if err != nil {
		return nil
	}
	return val
}

func GetString(k string) string {
	checkInstance()
	val, err := instance.GetString(k)
	if err != nil {
		return ""
	}
	return val
}
func GetStringDefault(k string, dv string) string {
	checkInstance()
	val, err := instance.GetString(k, dv)
	if err != nil {
		return ""
	}
	return val
}

func GetBool(k string) bool {
	checkInstance()
	val, _ := instance.GetString(k)
	if val == "yes" || val == "ok" || val == "1" {
		return true
	}
	if val == "no" || val == "ng" || val == "0" {
		return false
	}

	return false
}

func GetInt(k string) int {
	checkInstance()
	val, err := instance.GetInt(k)
	if err != nil {
		return 0
	}
	return val
}
