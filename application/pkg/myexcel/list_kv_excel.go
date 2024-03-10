// Package myexcel
// @Author: zhangdi
// @File: list_data_excel
// @Version: 1.0.0
// @Date: 2023/11/24 11:13
package myexcel

import (
	"fmt"
)

type KVItem struct {
	K string
	V any
}

// WriteKvToExcel 仿制Map 写入一个KV 目的是为了保证顺序可空 kWidth
func WriteKvToExcel(data []KVItem, savePath string, kWidth float64, vWidth float64) error {
	// 创建一个新的Excel文件
	f := excelize.NewFile()

	// 设置工作表名称为"Sheet1"
	index, _ := f.NewSheet("Sheet1")

	// 将map的键和值写入Excel的第一列和第二列
	for i, item := range data {
		row := i + 1
		// 写入第一列（A列）
		cell := fmt.Sprintf("A%d", row)
		f.SetCellValue("Sheet1", cell, item.K)

		// 写入第二列（B列）
		cell = fmt.Sprintf("B%d", row)
		f.SetCellValue("Sheet1", cell, item.V)
	}

	// 设置活动工作表为"Sheet1"
	f.SetActiveSheet(index)
	f.SetColWidth("Sheet1", "A", "A", kWidth)
	f.SetColWidth("Sheet1", "B", "B", vWidth)

	// 保存Excel文件
	err := f.SaveAs(savePath)
	if err != nil {
		return err
	}

	return nil
}
