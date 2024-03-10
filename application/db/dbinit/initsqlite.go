// Package dbinit
// @Author: zhangdi
// @File: initsqlite
// @Version: 1.0.0
// @Date: 2023/9/14 18:13
package dbinit

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"sync"
	filehelper "wailstemplate/application/pkg/filehelper"
)

var (
	Mutex sync.Mutex       // 定义一个互斥锁对象
	dbs   map[int]*gorm.DB //数据库对象列表，按年分库，一年1个db文件,key为年份
)

//// GetDb 使用互斥锁保护 db 对象，避免多个 goroutine 同时访问引发竞态条件
//func GetDb(year int) *gorm.DB {
//	Mutex.Lock()
//	defer Mutex.Unlock()
//	return dbs[year]
//}

// NewDbInstance 初始化连接Sqlite
func NewDbInstance(dbPath string) (*gorm.DB, error) {
	//判断文件夹是否存在，不存在先创建
	baseDir := filepath.Dir(dbPath)
	if !filehelper.FileExists(baseDir) {
		err := os.MkdirAll(baseDir, 0755)
		if err != nil {
			return nil, err
		}
	}
	if !filehelper.FileExists(dbPath) {
		file, err := os.Create(dbPath)
		if err != nil {
			file.Close()
			return nil, err
		}
		file.Close()
	}
	//连接数据库

	db, err := gorm.Open(sqlite.Open(dbPath+"?_pragma=encoding=UTF-8"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 调整最大打开连接数以匹配您的池大小
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	//SetMaxIdleConns:用于设置连接池中空闲连接的最大数量。在使用GORM进行数据库操作时，每次GORM调用完毕后，
	//数据库连接将变为空闲状态，等待下一次使用。因此，当空闲连接达到最大值时，
	//多余的连接就会被关闭。这样可以避免浪费资源和过多的数据库连接
	sqlDB.SetMaxIdleConns(10)

	//SetMaxOpenConns:用于设置同时打开的最大连接数（包括空闲和正在使用的连接）。
	//如果当前连接数已满，则进入等待状态，并在空闲连接没有足够可用时创建新连接。
	//通过适当地调整这些参数，可以更好地平衡数据库连接的使用和性能。
	//但应该注意的是，在这里设置的最大连接数不应该超过数据库的实际最大连接数限制。
	sqlDB.SetMaxOpenConns(20)

	//迁移表结构

	return db, nil
}
