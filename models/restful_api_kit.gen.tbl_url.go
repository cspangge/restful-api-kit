package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _TblURLMgr struct {
	*_BaseMgr
}

// TblURLMgr open func
func TblURLMgr(db *gorm.DB) *_TblURLMgr {
	if db == nil {
		panic(fmt.Errorf("TblURLMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TblURLMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tbl_url"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TblURLMgr) GetTableName() string {
	return "tbl_url"
}

// Get 获取
func (obj *_TblURLMgr) Get() (result TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TblURLMgr) Gets() (results []*TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_TblURLMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithURLID url_id获取
func (obj *_TblURLMgr) WithURLID(urlID uint) Option {
	return optionFunc(func(o *options) { o.query["url_id"] = urlID })
}

// WithURL url获取
func (obj *_TblURLMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithActive active获取 1-active; 2-inactive
func (obj *_TblURLMgr) WithActive(active int) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithCreatedAt created_at获取
func (obj *_TblURLMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreatedBy created_by获取
func (obj *_TblURLMgr) WithCreatedBy(createdBy int) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithUpdatedAt updated_at获取
func (obj *_TblURLMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithUpdatedBy updated_by获取
func (obj *_TblURLMgr) WithUpdatedBy(updatedBy int) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// GetByOption 功能选项模式获取
func (obj *_TblURLMgr) GetByOption(opts ...Option) (result TblURL, err error) {
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
func (obj *_TblURLMgr) GetByOptions(opts ...Option) (results []*TblURL, err error) {
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
func (obj *_TblURLMgr) GetFromID(id uint) (result TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_TblURLMgr) GetBatchFromID(ids []uint) (results []*TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromURLID 通过url_id获取内容
func (obj *_TblURLMgr) GetFromURLID(urlID uint) (results []*TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`url_id` = ?", urlID).Find(&results).Error

	return
}

// GetBatchFromURLID 批量查找
func (obj *_TblURLMgr) GetBatchFromURLID(urlIDs []uint) (results []*TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`url_id` IN (?)", urlIDs).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容
func (obj *_TblURLMgr) GetFromURL(url string) (results []*TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`url` = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量查找
func (obj *_TblURLMgr) GetBatchFromURL(urls []string) (results []*TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`url` IN (?)", urls).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容 1-active; 2-inactive
func (obj *_TblURLMgr) GetFromActive(active int) (results []*TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`active` = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量查找 1-active; 2-inactive
func (obj *_TblURLMgr) GetBatchFromActive(actives []int) (results []*TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`active` IN (?)", actives).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_TblURLMgr) GetFromCreatedAt(createdAt time.Time) (results []*TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_TblURLMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容
func (obj *_TblURLMgr) GetFromCreatedBy(createdBy int) (results []*TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量查找
func (obj *_TblURLMgr) GetBatchFromCreatedBy(createdBys []int) (results []*TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_TblURLMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_TblURLMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容
func (obj *_TblURLMgr) GetFromUpdatedBy(updatedBy int) (results []*TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量查找
func (obj *_TblURLMgr) GetBatchFromUpdatedBy(updatedBys []int) (results []*TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` IN (?)", updatedBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TblURLMgr) FetchByPrimaryKey(id uint) (result TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchIndexByURLIDx  获取多个内容
func (obj *_TblURLMgr) FetchIndexByURLIDx(url string) (results []*TblURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`url` = ?", url).Find(&results).Error

	return
}
