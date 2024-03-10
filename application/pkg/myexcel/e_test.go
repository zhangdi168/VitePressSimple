// Package myexcel
// @Author: zhangdi
// @File: e_test
// @Version: 1.0.0
// @Date: 2023/11/24 10:45
package myexcel

import (
	"fmt"
	"testing"
)

func TestMapToexcel(t *testing.T) {
	data := map[string]any{
		"Name":  "John22",
		"Age":   1,
		"Email": 88,
	}

	err := WriteMapToExcel(data, "output.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
