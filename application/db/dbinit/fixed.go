// Package dbinit
// @Author: zhangdi
// @File: fixed
// @Version: 1.0.0
// @Date: 2023/11/22 15:15
package dbinit

import (
	"gorm.io/gorm"
	"log"
	"path/filepath"
)

var fixedDb *gorm.DB

// GetFixedDb 获取固定数据的示例
func GetFixedDb() *gorm.DB {
	db := *fixedDb
	return &db
}

// 固定的数据 不跟随年份的改变而改变，也就是不管哪一年都在同一库同一个文件中保存
func initFixedDb(baseDir string, migrateEntityList []any) {
	//	固定数据，不管哪年都能查到
	dbFile := filepath.Join(baseDir, "fixed.db")
	//获取一个实例化对象
	var err error
	fixedDb, err = NewDbInstance(dbFile)
	if err != nil {
		log.Println(err.Error())
	}

	//迁移固定表：
	err = fixedDb.AutoMigrate(migrateEntityList...)
	if err != nil {
		log.Println(err.Error())
	}
}
