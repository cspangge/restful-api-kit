package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _SampleMgr struct {
	*_BaseMgr
}

// SampleMgr open func
func SampleMgr(db *gorm.DB) *_SampleMgr {
	if db == nil {
		panic(fmt.Errorf("SampleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SampleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sample"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SampleMgr) GetTableName() string {
	return "sample"
}

// Get 获取
func (obj *_SampleMgr) Get() (result Sample, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SampleMgr) Gets() (results []*Sample, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SampleMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_SampleMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_SampleMgr) GetByOption(opts ...Option) (result Sample, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SampleMgr) GetByOptions(opts ...Option) (results []*Sample, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_SampleMgr) GetFromID(id int) (result Sample, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_SampleMgr) GetBatchFromID(ids []int) (results []*Sample, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_SampleMgr) GetFromName(name string) (results []*Sample, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_SampleMgr) GetBatchFromName(names []string) (results []*Sample, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SampleMgr) FetchByPrimaryKey(id int) (result Sample, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}
