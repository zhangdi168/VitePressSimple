// Package myexcel
// @Author: zhangdi
// @File: map_excel
// @Version: 1.0.0
// @Date: 2023/11/24 10:44
package myexcel

import (
	"fmt"
)

func WriteMapToExcel(data map[string]any, savePath string) error {
	// 创建一个新的Excel文件
	f := excelize.NewFile()

	// 设置工作表名称为"Sheet1"
	index, _ := f.NewSheet("Sheet1")

	row := 1
	// 将map的键和值写入Excel的第一列和第二列
	for key, value := range data {
		// 写入第一列（A列）
		cell := fmt.Sprintf("A%d", row)
		f.SetCellValue("Sheet1", cell, key)

		// 写入第二列（B列）
		cell = fmt.Sprintf("B%d", row)
		f.SetCellValue("Sheet1", cell, value)

		row++
	}

	// 设置活动工作表为"Sheet1"
	f.SetActiveSheet(index)

	// 保存Excel文件
	err := f.SaveAs(savePath)
	if err != nil {
		return err
	}

	return nil
}
