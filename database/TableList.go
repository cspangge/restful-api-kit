package database

const (
	ROLE_ID = 100
	URL_ID  = 101
	USER_ID = 102
)

var TABLE_CATALOG = map[int]string{
	ROLE_ID: "tbl_role",
	URL_ID:  "tbl_url",
	USER_ID: "tbl_user",
}
