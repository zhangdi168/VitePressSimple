// Package myexcel
// @Author: zhangdi
// @File: write
// @Version: 1.0.0
// @Date: 2023/9/21 15:01
package myexcel

import (
	"fmt"
)

var colNames = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

type ExcelHelper struct {
	f            *excelize.File
	currRowIndex int      //当前写入的函数
	colHeaders   []string //表头
}

func NewExcelHelper() *ExcelHelper {
	file := excelize.NewFile()
	return &ExcelHelper{
		f:            file,
		currRowIndex: 1,
	}
}

// SetHeaders 设置表头和默认列宽20
func (h *ExcelHelper) SetHeaders(colHeaders []string, cloWidth float64) string {

	h.colHeaders = colHeaders
	for i, headerName := range colHeaders {
		colName := colNames[i]
		//设置列宽
		err := h.f.SetColWidth("Sheet1", colName, colName, cloWidth)
		if err != nil {
			return err.Error()
		}
		//设置表头
		err = h.f.SetCellValue("Sheet1", colName+"1", headerName)
		if err != nil {
			return err.Error()
		}
	}
	h.currRowIndex++
	return ""
}

// AppendRows 向最后一行后面插入多行数据(批量时推荐该方法)
// rows := [][]any{
// {"8881", "sasada", "dadada", "dadadad", "", ""},
// {"8882", "sasada", "dadada", "dadadad", "", ""},
// {"8883", "sasada", "dadada", "dadadad", "", ""},
// }
func (h *ExcelHelper) AppendRows(rows [][]any) string {
	if h.colHeaders == nil {
		return "未设置表头"
	}
	// 批量插入数据到工作表中
	for rowIndex, rowData := range rows {
		err := h.f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", rowIndex+h.currRowIndex), &rowData)
		if err != nil {
			return "插入数据失败:" + err.Error()
		}
	}

	h.currRowIndex += len(rows)
	return ""
}

// AppendRow 向最后一行插入一行数据
func (h *ExcelHelper) AppendRow(row ...string) string {
	if h.colHeaders == nil {
		return "未设置表头"
	}
	if len(row) != len(h.colHeaders) {
		return fmt.Sprintf("插入项数[%d]与表头列数[%d]不一致", len(row), len(h.colHeaders))
	}
	rowIndexStr := fmt.Sprintf("%v", h.currRowIndex)
	for i, item := range row {
		colName := colNames[i] //列名 如B
		err := h.f.SetCellValue("Sheet1", colName+rowIndexStr, item)
		if err != nil {
			return err.Error()
		}
	}
	h.currRowIndex++
	return ""
}

func (h *ExcelHelper) SaveToFile(filePath string) error {
	return h.f.SaveAs(filePath)
}

// Close 关闭excel
func (h *ExcelHelper) Close() error {
	return h.f.Close()
}
