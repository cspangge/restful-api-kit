package models

import (
	"time"
)

// TblUser user info
type TblUser struct {
	ID         uint      `gorm:"primaryKey;column:id;type:int unsigned;not null"`
	Name       string    `gorm:"column:name;type:varchar(50);not null;default:''"`
	Email      string    `gorm:"unique;index:tbl_user_email_IDX;column:email;type:varchar(50);not null;default:''"`
	Pwd        string    `gorm:"index:tbl_user_email_IDX;column:pwd;type:varchar(50);not null;default:''"`
	Role       int       `gorm:"column:role;type:int;not null;default:0"`
	Active     int       `gorm:"column:active;type:int;not null;default:0"`
	DisableURL string    `gorm:"column:disable_url;type:varchar(255);not null;default:''"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	CreatedBy  int       `gorm:"column:created_by;type:int;not null;default:0"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedBy  int       `gorm:"column:updated_by;type:int;not null;default:0"`
}

// TblUserColumns get sql column name.获取数据库列名
var TblUserColumns = struct {
	ID         string
	Name       string
	Email      string
	Pwd        string
	Role       string
	Active     string
	DisableURL string
	CreatedAt  string
	CreatedBy  string
	UpdatedAt  string
	UpdatedBy  string
}{
	ID:         "id",
	Name:       "name",
	Email:      "email",
	Pwd:        "pwd",
	Role:       "role",
	Active:     "active",
	DisableURL: "disable_url",
	CreatedAt:  "created_at",
	CreatedBy:  "created_by",
	UpdatedAt:  "updated_at",
	UpdatedBy:  "updated_by",
}
