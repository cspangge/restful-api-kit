package model

// TblUser user info
type TblUser struct {
	ID    uint   `gorm:"primaryKey;column:id;type:int unsigned;not null"`
	Name  string `gorm:"column:name;type:varchar(50);not null"`
	Email string `gorm:"unique;index:tbl_user_email_IDX;column:email;type:varchar(50);not null"`
	Pwd   string `gorm:"index:tbl_user_email_IDX;column:pwd;type:varchar(50);not null"`
}

// TblUserColumns get sql column name.获取数据库列名
var TblUserColumns = struct {
	ID    string
	Name  string
	Email string
	Pwd   string
}{
	ID:    "id",
	Name:  "name",
	Email: "email",
	Pwd:   "pwd",
}
