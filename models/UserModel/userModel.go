package UserModel

import (
	"gorm.io/gorm"
	"restful-api-kit/models"
)

func GetUserByEmail(db *gorm.DB, email string) (models.TblUser, error) {
	userModel := models.TblUserMgr(db)
	return userModel.GetByOption(userModel.WithEmail(email))
}

func GetUserByEmailPwd(db *gorm.DB, email, pwd string) (*models.TblUser, error) {
	userModel := models.TblUserMgr(db)
	return userModel.FetchByTblUserEmail(email, pwd)
}
