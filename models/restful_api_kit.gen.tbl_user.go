package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

const (
	ACTIVE   = 1
	INACTIVE = 2
)

type _TblUserMgr struct {
	*_BaseMgr
}

// TblUserMgr open func
func TblUserMgr(db *gorm.DB) *_TblUserMgr {
	if db == nil {
		panic(fmt.Errorf("TblUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TblUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tbl_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TblUserMgr) GetTableName() string {
	return "tbl_user"
}

// Get 获取
func (obj *_TblUserMgr) Get() (result TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TblUserMgr) Gets() (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_TblUserMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取
func (obj *_TblUserMgr) WithUserID(userID uint) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithName name获取
func (obj *_TblUserMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithEmail email获取
func (obj *_TblUserMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithPwd pwd获取
func (obj *_TblUserMgr) WithPwd(pwd string) Option {
	return optionFunc(func(o *options) { o.query["pwd"] = pwd })
}

// WithRole role获取
func (obj *_TblUserMgr) WithRole(role int) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// WithActive active获取 1-active; 2-inactive
func (obj *_TblUserMgr) WithActive(active int) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithDisableURL disable_url获取
func (obj *_TblUserMgr) WithDisableURL(disableURL string) Option {
	return optionFunc(func(o *options) { o.query["disable_url"] = disableURL })
}

// WithCreatedAt created_at获取
func (obj *_TblUserMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreatedBy created_by获取
func (obj *_TblUserMgr) WithCreatedBy(createdBy int) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithUpdatedAt updated_at获取
func (obj *_TblUserMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithUpdatedBy updated_by获取
func (obj *_TblUserMgr) WithUpdatedBy(updatedBy int) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// GetByOption 功能选项模式获取
func (obj *_TblUserMgr) GetByOption(opts ...Option) (result TblUser, err error) {
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
func (obj *_TblUserMgr) GetByOptions(opts ...Option) (results []*TblUser, err error) {
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
func (obj *_TblUserMgr) GetFromID(id uint) (result TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_TblUserMgr) GetBatchFromID(ids []uint) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_TblUserMgr) GetFromUserID(userID uint) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_TblUserMgr) GetBatchFromUserID(userIDs []uint) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_TblUserMgr) GetFromName(name string) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_TblUserMgr) GetBatchFromName(names []string) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容
func (obj *_TblUserMgr) GetFromEmail(email string) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`email` = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量查找
func (obj *_TblUserMgr) GetBatchFromEmail(emails []string) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`email` IN (?)", emails).Find(&results).Error

	return
}

// GetFromPwd 通过pwd获取内容
func (obj *_TblUserMgr) GetFromPwd(pwd string) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`pwd` = ?", pwd).Find(&results).Error

	return
}

// GetBatchFromPwd 批量查找
func (obj *_TblUserMgr) GetBatchFromPwd(pwds []string) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`pwd` IN (?)", pwds).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容
func (obj *_TblUserMgr) GetFromRole(role int) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找
func (obj *_TblUserMgr) GetBatchFromRole(roles []int) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容 1-active; 2-inactive
func (obj *_TblUserMgr) GetFromActive(active int) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`active` = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量查找 1-active; 2-inactive
func (obj *_TblUserMgr) GetBatchFromActive(actives []int) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`active` IN (?)", actives).Find(&results).Error

	return
}

// GetFromDisableURL 通过disable_url获取内容
func (obj *_TblUserMgr) GetFromDisableURL(disableURL string) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`disable_url` = ?", disableURL).Find(&results).Error

	return
}

// GetBatchFromDisableURL 批量查找
func (obj *_TblUserMgr) GetBatchFromDisableURL(disableURLs []string) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`disable_url` IN (?)", disableURLs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_TblUserMgr) GetFromCreatedAt(createdAt time.Time) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_TblUserMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容
func (obj *_TblUserMgr) GetFromCreatedBy(createdBy int) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量查找
func (obj *_TblUserMgr) GetBatchFromCreatedBy(createdBys []int) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_TblUserMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_TblUserMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容
func (obj *_TblUserMgr) GetFromUpdatedBy(updatedBy int) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量查找
func (obj *_TblUserMgr) GetBatchFromUpdatedBy(updatedBys []int) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`updated_by` IN (?)", updatedBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TblUserMgr) FetchByPrimaryKey(id uint) (result TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByTblUserUn primay or index 获取唯一内容
func (obj *_TblUserMgr) FetchUniqueByTblUserUn(email string) (result TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`email` = ?", email).Find(&result).Error

	return
}

// FetchIndexByTblUserEmailIDX  获取多个内容
func (obj *_TblUserMgr) FetchIndexByTblUserEmailIDX(email string, pwd string) (results []*TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`email` = ? AND `pwd` = ?", email, pwd).Find(&results).Error

	return
}

func (obj *_TblUserMgr) FetchByTblUserEmail(email string, pwd string) (results *TblUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`email` = ? AND `pwd` = ?", email, pwd).Find(&results).Error

	return
}
