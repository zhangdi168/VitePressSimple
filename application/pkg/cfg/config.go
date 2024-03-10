package cfg

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	mu sync.Mutex
	v  *viper.Viper
}

// NewConfig 函数创建一个Config实例。
func NewConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")

	// 设置默认配置选项
	// ...

	return &Config{v: v}, nil
}

// LoadConfig 函数从YAML文件加载配置。
func (c *Config) LoadConfig(path string) error {
	// 检查文件
	err := c.checkFile(path)
	if err != nil {
		return err
	}

	// 读取YAML配置文件
	c.mu.Lock()
	defer c.mu.Unlock()

	c.v.SetConfigFile(path)
	if err := c.v.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %s", path)
	}

	// 验证配置
	if err := c.validateConfig(); err != nil {
		return fmt.Errorf("无效的配置文件: %s", err)
	}

	return nil
}

func (c *Config) checkFile(configPath string) error {
	//判断文件configPath_的目录不存在则先创建目录，再判断configPath_文件不存在则创建文件
	// 获取目录路径
	dir := filepath.Dir(configPath)
	// 判断目录是否存在，不存在则创建
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	// 判断文件是否存在，不存在则创建
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		_, err = os.Create(configPath)
		if err != nil {
			return err
		}
	}

	return nil
}

// validateConfig 函数可用于验证加载的配置文件。
func (c *Config) validateConfig() error {
	// 验证必需的配置选项
	// ...

	return nil
}

// Get 函数从配置中获取指定键的值。
// 如果键不存在，则返回defaultValue（如果提供了），否则返回错误。
func (c *Config) Get(key string, defaultValue ...interface{}) (interface{}, error) {
	if !c.v.IsSet(key) {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return nil, errors.New("未知的配置键")
	}

	return c.v.Get(key), nil
}

// GetString 函数从配置中获取指定键的值。
// 如果键不存在，则返回defaultValue（如果提供了），否则返回错误。
func (c *Config) GetString(key string, defaultValue ...string) (string, error) {
	if !c.v.IsSet(key) {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return "", errors.New("未知的配置键")
	}

	return c.v.GetString(key), nil
}

// GetInt 函数从配置中获取指定键的值。
// 如果键不存在，则返回defaultValue（如果提供了），否则返回错误。
func (c *Config) GetInt(key string, defaultValue ...int) (int, error) {
	if !c.v.IsSet(key) {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return 0, errors.New("未知的配置键")
	}

	return c.v.GetInt(key), nil
}

// GetBool 函数从配置中获取指定键的值。
// 如果键不存在，则返回defaultValue（如果提供了），否则返回错误。
func (c *Config) GetBool(key string, defaultValue ...bool) (bool, error) {
	if !c.v.IsSet(key) {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return false, errors.New("未知的配置键")
	}

	return c.v.GetBool(key), nil
}

// Set 函数在配置中设置指定键的值。
// 它还将更改写入配置文件。
func (c *Config) Set(key string, value interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.v.Set(key, value)

	// 将配置写入文件
	if err := c.v.WriteConfig(); err != nil {
		return errors.New("写入配置文件失败:" + err.Error())

	}
	return nil
}
