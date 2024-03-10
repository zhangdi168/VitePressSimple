// Package script
// @Author: zhangdi
// @File: gen_column_const_test
// @Version: 1.0.0
// @Date: 2023/8/31 13:47
package scripts

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"wailstemplate/application/pkg/varname"
	setting "wailstemplate/settings"
)

func TestGenCol(t1 *testing.T) {

	for _, entity_ := range setting.EntityAutoMigrateList {
		t := reflect.TypeOf(entity_).Elem()
		fileName := strings.ToLower(t.Name()) + ".go"
		content := "package columns \n\n"
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			constName := t.Name() + field.Name
			fieldName := varname.NameToLine(field.Name)
			if fieldName == "model" { //gorm.model
				content += fmt.Sprintf("const %s = \"%s\"\n\n", t.Name()+"ID", "id")
				content += fmt.Sprintf("const %s = \"%s\"\n\n", t.Name()+"Created", "created_at")
				content += fmt.Sprintf("const %s = \"%s\"\n\n", t.Name()+"Updated", "updated_at")
			} else {
				content += fmt.Sprintf("const %s = \"%s\"\n\n", constName, fieldName)
			}

		}
		filePath := filepath.Join(GetColumnsSaveDir(), fileName)
		err := ioutil.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			continue
		}
		fmt.Println("Data written to file.")
	}

}

func GetColumnsSaveDir() string {
	return filepath.Join(ProjectDir(), "application", "constant", "columns")
}
