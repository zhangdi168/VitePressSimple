// Package cfg
// @Author: zhangdi
// @File: helper_test
// @Version: 1.0.0
// @Date: 2023/11/22 12:45
package cfg

import (
	"fmt"
	"sync"
	"testing"
)

// 测试多线程下批量写入,是否安全
func TestSaveBatch(t *testing.T) {
	err := InitCfg("cfg1.yaml")
	if err != nil {
		println(err.Error())
	}
	var wg sync.WaitGroup
	wg.Add(6)
	go writeKey("test1.key", &wg)
	go writeKey("test2.key", &wg)
	go writeKey("test3.key", &wg)
	go writeKey("test4.key", &wg)
	go writeKey("test1.key", &wg)
	go writeKey("test2.key", &wg)
	wg.Wait()
}

func TestReadBatch(t *testing.T) {
	err := InitCfg("cfg1.yaml")
	if err != nil {
		println(err.Error())
	}
	var wg sync.WaitGroup
	wg.Add(6)
	go Readkey("test1.key", &wg)
	go Readkey("test2.key", &wg)
	go Readkey("test3.key", &wg)
	go Readkey("test4.key", &wg)
	go Readkey("test1.key", &wg)
	go Readkey("test2.key", &wg)
	wg.Wait()
}

func Readkey(baseKey string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		k := fmt.Sprintf("%s%d", baseKey, i)
		println(Get(k))
	}
}
func writeKey(baseKey string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		k := fmt.Sprintf("%s%d", baseKey, i)
		err := Set(k, i)
		if err != nil {
			println(k + err.Error())
		}
	}
}
