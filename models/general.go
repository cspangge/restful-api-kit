package models

import (
	"fmt"
	"gorm.io/gorm"
)

func Exist(tx *gorm.DB, tbl string, query interface{}, args ...interface{}) bool {
	var count int64
	tx.Table(tbl).Where(query, args...).Count(&count)
	if count == 0 {
		return false
	}
	return true
}

func Count(tx *gorm.DB, tbl string, query interface{}, args ...interface{}) int64 {
	var count int64
	tx.Table(tbl).Where(query, args...).Count(&count)
	return count
}

func GetMaxId(tx *gorm.DB, tbl, field string) uint {
	var id uint
	tx.Raw(fmt.Sprintf("SELECT MAX(%s) FROM %s", field, tbl)).Scan(&id)
	return id
}
