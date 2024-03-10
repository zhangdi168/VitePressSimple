// Package repository
// @Author: zhangdi
// @File: base_db
// @Version: 1.0.0
// @Date: 2023/11/22 15:29
package repository

import (
	"gorm.io/gorm"
	"wailstemplate/application/pkg/utils"
	"wailstemplate/application/pkg/where"
)

// CreateWithDb 创建
func CreateWithDb(newDtoCreate any, newEntity any, newDtoInfo any, db *gorm.DB) error {
	entityInfo := newEntity                     //为分配一个内存空间
	utils.DtoToEntity(newDtoCreate, entityInfo) //将dto转成entity
	result := db.Create(entityInfo)
	utils.DtoToEntity(entityInfo, newDtoInfo)
	return result.Error
}

// FindOneWithDb 查找一个
func FindOneWithDb(wheres []*where.Field, newEntity any, getSoftDelData bool, db *gorm.DB) error {
	whereStr, whereValues, err := where.ConvertToGormWhere(wheres)
	if err != nil {
		return err
	}
	tx := db.Where(whereStr, whereValues...)
	if getSoftDelData {
		//Unscoped()，如果关联模型被软删除了，它们将被包括在查询结果中
		tx = tx.Unscoped()
	}
	if err = tx.Find(newEntity).Error; err != nil {
		return err
	}

	return err
}

// FindListWithDb 查找列表  --getSoftDelData是否获取软删除数据
func FindListWithDb(
	wheres []*where.Field, extra *where.Extra, newEntityList any, getSoftDelData bool, db *gorm.DB) error {

	db, _ = formatDbExtra(db, extra)     //组装附加条件
	db, err := formatDbWhere(db, wheres) //组装wheres
	if err != nil {
		return err
	}

	if getSoftDelData {
		//Unscoped()，如果关联模型被软删除了，它们将被包括在查询结果中
		db = db.Unscoped()
	}

	err = db.Find(newEntityList).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateWithDb  指定配置更新
// 要特别注意零值问题，updateColumnsCfg可以指定要更新的字段类型或者字段名，如果不传入则默认更新非零值字段。
// updateColumnsCfg只传入一个值时会先跟update_config里面的几个常量进行匹配,
// 匹配到了则按照里面的规则进行更新，如UpdateColumnsNotZeroAndNumber表示更新所有非0值和整型的字段
// 否则会按照其是一个指定字段来更新，传入多个值是会认为是手动指定要更新字段。
func UpdateWithDb(
	wheres []*where.Field,
	newEntity any, //更新的entity
	dtoUpdate any, //更新的dto
	db *gorm.DB, //db实例
	updateColumnsCfg ...string, //指定更新的字段，不管是否为零值(默认零值字段是不更新的)
) error {
	updateEntity := newEntity //为分配一个内存空间

	utils.DtoToEntity(dtoUpdate, updateEntity) //将dto转成entity

	updateColumns := make([]string, 0) //要更新的字段
	if updateColumnsCfg != nil && len(updateColumnsCfg) > 0 {
		if len(updateColumnsCfg) == 1 && utils.ArrContains(UpdateTypList, updateColumnsCfg[0]) {
			//按照指定的更新规则来更新
			updateColumns = GetUpdateColumns(dtoUpdate, updateColumnsCfg[0])
		} else {
			//多个值 或1个值但是不属于类型指定 则组装更新的字段列表
			for _, column := range updateColumnsCfg {
				updateColumns = append(updateColumns, column)
			}
		}
	}

	whereStr, whereValues, err := where.ConvertToGormWhere(wheres)
	if err != nil {
		return err
	}

	dbx := db.Model(newEntity).Where(whereStr, whereValues...)
	if updateColumns == nil || len(updateColumns) == 0 {
		//不指定字段，gorm内部会默认更新所有非零值数据
		dbx.Updates(updateEntity)
	} else {
		//指定要更新的字段（解决零值问题）
		dbx.Select(updateColumns).Updates(updateEntity)
	}

	if dbx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return dbx.Error
}

// DeleteWithDb 删除,softDelete--是否软删除。deleted_at字段存在时默认会软删除
func DeleteWithDb(wheres []*where.Field, newEntity any, softDelete bool, db *gorm.DB) error {
	whereStr, whereValues, err := where.ConvertToGormWhere(wheres)
	if err != nil {
		return err
	}

	tx := db.
		Where(whereStr, whereValues...)
	if !softDelete {
		//Unscoped可以在执行时忽略软删除功能，实现真实删除
		tx = tx.Unscoped()
	}
	err = tx.Delete(newEntity).Error
	if err != nil {
		return err
	}
	return nil
}

// GetTotalWithDb 获取符合某个条件的数量
func GetTotalWithDb(wheres []*where.Field, newEntity any, db *gorm.DB) (int64, error) {
	var count int64
	whereStr, whereValues, err := where.ConvertToGormWhere(wheres)
	if err != nil {
		return -1, err
	}

	err = db.
		Model(newEntity).Where(whereStr, whereValues...).
		Count(&count).Error
	if err != nil {
		return -1, err
	}
	return count, nil
}
