package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
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
