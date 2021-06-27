package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _TblRoleMgr struct {
	*_BaseMgr
}

// TblRoleMgr open func
func TblRoleMgr(db *gorm.DB) *_TblRoleMgr {
	if db == nil {
		panic(fmt.Errorf("TblRoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TblRoleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tbl_role"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TblRoleMgr) GetTableName() string {
	return "tbl_role"
}

// Get 获取
func (obj *_TblRoleMgr) Get() (result TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TblRoleMgr) Gets() (results []*TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_TblRoleMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithRoleID role_id获取
func (obj *_TblRoleMgr) WithRoleID(roleID uint) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithName name获取
func (obj *_TblRoleMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithActive active获取 1-active; 2-inactive
func (obj *_TblRoleMgr) WithActive(active int) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithURLs urls获取
func (obj *_TblRoleMgr) WithURLs(urls string) Option {
	return optionFunc(func(o *options) { o.query["urls"] = urls })
}

// WithCreatedAt created_at获取
func (obj *_TblRoleMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreatedBy created_by获取
func (obj *_TblRoleMgr) WithCreatedBy(createdBy int) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithUpdatedAt updated_at获取
func (obj *_TblRoleMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithUpdatedBy updated_by获取
func (obj *_TblRoleMgr) WithUpdatedBy(updatedBy int) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// GetByOption 功能选项模式获取
func (obj *_TblRoleMgr) GetByOption(opts ...Option) (result TblRole, err error) {
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
func (obj *_TblRoleMgr) GetByOptions(opts ...Option) (results []*TblRole, err error) {
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
func (obj *_TblRoleMgr) GetFromID(id uint) (result TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_TblRoleMgr) GetBatchFromID(ids []uint) (results []*TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromRoleID 通过role_id获取内容
func (obj *_TblRoleMgr) GetFromRoleID(roleID uint) (results []*TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`role_id` = ?", roleID).Find(&results).Error

	return
}

// GetBatchFromRoleID 批量查找
func (obj *_TblRoleMgr) GetBatchFromRoleID(roleIDs []uint) (results []*TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`role_id` IN (?)", roleIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_TblRoleMgr) GetFromName(name string) (results []*TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_TblRoleMgr) GetBatchFromName(names []string) (results []*TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容 1-active; 2-inactive
func (obj *_TblRoleMgr) GetFromActive(active int) (results []*TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`active` = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量查找 1-active; 2-inactive
func (obj *_TblRoleMgr) GetBatchFromActive(actives []int) (results []*TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`active` IN (?)", actives).Find(&results).Error

	return
}

// GetFromURLs 通过urls获取内容
func (obj *_TblRoleMgr) GetFromURLs(urls string) (results []*TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`urls` = ?", urls).Find(&results).Error

	return
}

// GetBatchFromURLs 批量查找
func (obj *_TblRoleMgr) GetBatchFromURLs(urlss []string) (results []*TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`urls` IN (?)", urlss).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_TblRoleMgr) GetFromCreatedAt(createdAt time.Time) (results []*TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_TblRoleMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容
func (obj *_TblRoleMgr) GetFromCreatedBy(createdBy int) (results []*TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量查找
func (obj *_TblRoleMgr) GetBatchFromCreatedBy(createdBys []int) (results []*TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_TblRoleMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_TblRoleMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容
func (obj *_TblRoleMgr) GetFromUpdatedBy(updatedBy int) (results []*TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量查找
func (obj *_TblRoleMgr) GetBatchFromUpdatedBy(updatedBys []int) (results []*TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` IN (?)", updatedBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TblRoleMgr) FetchByPrimaryKey(id uint) (result TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchIndexByNameIDx  获取多个内容
func (obj *_TblRoleMgr) FetchIndexByNameIDx(name string) (results []*TblRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error

	return
}
