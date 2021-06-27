package models

import (
	"time"
)

// Sample [...]
type Sample struct {
	ID   int    `gorm:"primaryKey;column:id;type:int;not null"`
	Name string `gorm:"column:name;type:varchar(45)"`
}

// SampleColumns get sql column name.获取数据库列名
var SampleColumns = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}

// TblRole role info
type TblRole struct {
	ID        uint      `gorm:"primaryKey;column:id;type:int unsigned;not null"`
	RoleID    uint      `gorm:"column:role_id;type:int unsigned;not null"`
	Name      string    `gorm:"index:name_idx;column:name;type:varchar(50);not null;default:''"`
	Active    int       `gorm:"column:active;type:int;not null;default:2"` // 1-active; 2-inactive
	URLs      string    `gorm:"column:urls;type:varchar(255);not null;default:''"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	CreatedBy int       `gorm:"column:created_by;type:int;not null;default:0"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedBy int       `gorm:"column:updated_by;type:int;not null;default:0"`
}

// TblRoleColumns get sql column name.获取数据库列名
var TblRoleColumns = struct {
	ID        string
	RoleID    string
	Name      string
	Active    string
	URLs      string
	CreatedAt string
	CreatedBy string
	UpdatedAt string
	UpdatedBy string
}{
	ID:        "id",
	RoleID:    "role_id",
	Name:      "name",
	Active:    "active",
	URLs:      "urls",
	CreatedAt: "created_at",
	CreatedBy: "created_by",
	UpdatedAt: "updated_at",
	UpdatedBy: "updated_by",
}

// TblURL url info
type TblURL struct {
	ID        uint      `gorm:"primaryKey;column:id;type:int unsigned;not null"`
	URLID     uint      `gorm:"column:url_id;type:int unsigned;not null"`
	URL       string    `gorm:"index:url_idx;column:url;type:varchar(255);not null;default:''"`
	Active    int       `gorm:"column:active;type:int;not null;default:2"` // 1-active; 2-inactive
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	CreatedBy int       `gorm:"column:created_by;type:int;not null;default:0"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedBy int       `gorm:"column:updated_by;type:int;not null;default:0"`
}

// TblURLColumns get sql column name.获取数据库列名
var TblURLColumns = struct {
	ID        string
	URLID     string
	URL       string
	Active    string
	CreatedAt string
	CreatedBy string
	UpdatedAt string
	UpdatedBy string
}{
	ID:        "id",
	URLID:     "url_id",
	URL:       "url",
	Active:    "active",
	CreatedAt: "created_at",
	CreatedBy: "created_by",
	UpdatedAt: "updated_at",
	UpdatedBy: "updated_by",
}

// TblUser user info
type TblUser struct {
	ID         uint      `gorm:"primaryKey;column:id;type:int unsigned;not null"`
	UserID     uint      `gorm:"column:user_id;type:int unsigned;not null"`
	Name       string    `gorm:"column:name;type:varchar(50);not null;default:''"`
	Email      string    `gorm:"unique;index:tbl_user_email_IDX;column:email;type:varchar(50);not null;default:''"`
	Pwd        string    `gorm:"index:tbl_user_email_IDX;column:pwd;type:varchar(50);not null;default:''"`
	Role       int       `gorm:"column:role;type:int;not null;default:0"`
	Active     int       `gorm:"column:active;type:int;not null;default:2"` // 1-active; 2-inactive
	DisableURL string    `gorm:"column:disable_url;type:varchar(255);not null;default:''"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	CreatedBy  int       `gorm:"column:created_by;type:int;not null;default:0"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedBy  int       `gorm:"column:updated_by;type:int;not null;default:0"`
}

// TblUserColumns get sql column name.获取数据库列名
var TblUserColumns = struct {
	ID         string
	UserID     string
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
	UserID:     "user_id",
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
