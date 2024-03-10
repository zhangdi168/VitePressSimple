package repository

import (
	"gorm.io/gorm"
	"wailstemplate/application/pkg/where"
)

func formatDbWhere(tx *gorm.DB, wheres []*where.Field) (*gorm.DB, error) {
	if wheres == nil || len(wheres) == 0 {
		return tx, nil
	}
	//where 构造
	whereStr, whereValues, err := where.ConvertToGormWhere(wheres)
	if err != nil {
		return tx, err
	}
	tx = tx.Where(whereStr, whereValues...)
	return tx, err
}

func formatDbExtra(tx *gorm.DB, extra *where.Extra) (*gorm.DB, error) {
	if extra == nil { //传入空值nil则使用零值
		extra = &where.Extra{}
	}
	//排序方式
	if extra.OrderByColumn != "" {
		if extra.OrderByDesc {
			tx = tx.Order(extra.OrderByColumn + " desc")
		} else {
			tx = tx.Order(extra.OrderByColumn + " desc")
		}
	}
	//limit
	if extra.PageSize > 0 {
		tx = tx.Limit(extra.PageSize)
	}

	//offset条件
	if extra.PageNum > 0 {
		tx = tx.Offset(calculateOffset(extra.PageNum, extra.PageSize))
	} else {
		//取消offset条件
		tx = tx.Offset(-1)
	}

	return tx, nil
}
