package dbinit

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"path/filepath"
	"sync"
	"time"
)

var db *gorm.DB
var CurrYear int
var mu = sync.Mutex{}
var StartYear = 2023 //起始年份

func init() {
	CurrYear = time.Now().Year()
}

// GetDbInstance 获取今年的db示例
func GetDbInstance() *gorm.DB {
	mu.Lock()
	defer mu.Unlock()
	db := *dbs[CurrYear]
	return &db
}

func GetDbInstanceYear(year int) (*gorm.DB, error) {
	db, ok := dbs[year]
	if ok {
		return db, nil
	}
	return nil, fmt.Errorf("所选年份无数据")
}

func InitDb(baseDir string, migrateEntityList []any) {

	//initYearDb(baseDir)
	//初始化固定表
	initFixedDb(baseDir, migrateEntityList)

}
func needMigrate(migrateEntityList any) bool {
	if migrateEntityList != nil {
		// 表示没有传入任何参数，等同于认为是空列表
		fmt.Println("No entities were provided for migration.")
		return true
	} else {
		// 迭代检查每个元素是否为 nil（假设它们是指针或其他可以为 nil 的类型）
		//for _, entity := range migrateEntityList {
		//	if entity != nil {
		//		arr, ok := entity.([]any)
		//		if ok {
		//			// 如果是切片，那么检查是否为空
		//			if len(arr) > 0 && arr[0] != nil {
		//				return true
		//			}
		//		}
		//	}
		//}

		return true
	}
}

func initYearDb(baseDir string, entityList ...any) {
	var dbFilePathList []string
	//依次构造db年份
	for i := StartYear; i <= time.Now().Year(); i++ {
		currYear := fmt.Sprintf("%v", i)
		dbFilePathList = append(dbFilePathList, filepath.Join(baseDir, "data", "dbs", currYear+".db"))
	}

	//开始实例化数据
	dbs = make(map[int]*gorm.DB, time.Now().Year()-StartYear+1)
	for i, path := range dbFilePathList {
		//获取一个实例化对象
		db, err := NewDbInstance(path)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		//将db实例存入map
		if db != nil {
			dbs[i+StartYear] = db
		}

		//迁移表(只迁移当前年的)
		if i == (len(dbFilePathList) - 1) {
			//自动迁移表结构
			err := db.AutoMigrate(entityList)
			if err != nil {
				fmt.Println("迁移表结构失败" + err.Error())
				panic(err)
			}
		}

	}

}
