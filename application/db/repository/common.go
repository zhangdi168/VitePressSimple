package repository

import (
	"fmt"
	"gorm.io/gorm"
	"wailstemplate/application/pkg/where"
)

// FindDistinct 查找1个字段的所有重复项,minRepeat重复数
// entity 传入示例&user{}或new(user),sql demo如下:
// SELECT bottle_number as value, COUNT(*) as count
// FROM bottle_data
// GROUP BY bottle_number
// HAVING COUNT(bottle_number) > 1
func FindDistinct(entity any, columnName string,
	wheres []*where.Field, extra *where.Extra, db *gorm.DB) (
	[]*DistinctItem, error) {

	minRepeat := 1
	var err error
	var resultList []*DistinctItem

	db, _ = formatDbExtra(db.Model(entity), extra)

	db = db.
		Select(fmt.Sprintf(" %s as Value, COUNT(*) as Count", columnName)).
		Group(fmt.Sprintf("%s", columnName)).
		Having(fmt.Sprintf("COUNT(%s) > %d", columnName, minRepeat))
	db, err = formatDbWhere(db, wheres)
	db.Find(&resultList)

	if db.Error != nil {
		return nil, err
	}
	return resultList, err
}

// 根据页码和页面大小计算偏移量
func calculateOffset(page, pageSize int) int {
	offset := (page - 1) * pageSize
	return offset
}
