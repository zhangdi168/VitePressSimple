package myhttp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string, data interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}

	return nil
}

func HttpGetWithHeader(url string, headers map[string]string) (*http.Response, error) {
	// 构造请求参数
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 请求头
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// HttpPostWhitHeader 发送http post请求
func HttpPostWhitHeader(url string, data interface{}, headers map[string]string) (*http.Response, error) {
	//组装请求数据
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	//构造请求参数
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	//请求头
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	//发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
