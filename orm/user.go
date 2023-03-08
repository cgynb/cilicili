package orm

import (
	"cilicili/config"
	"cilicili/utils"
	"gorm.io/gorm"
)

const (
	user  = 0
	admin = 1
)

func GetUser(by, info string) (user *User, ok bool) {
	// id name
	result := DB.Where(by+" = ?", info).First(&user)
	return user, result.Error == nil
}

func GetUsers(condition string, page int) (users []*User, ok bool) {
	result := DB.Where(condition).Limit(config.Conf.PageSize).Offset(config.Conf.PageSize * (page - 1)).Find(&users)
	return users, result.Error == nil
}

func CreateUser(name, email, pwd string) (u *User, ok bool, msg string) {
	hpwd, err := utils.GenHashPassword(pwd)
	if err != nil {
		return nil, false, "please register again"
	}
	u = &User{Name: name, Password: hpwd, Email: email, Role: user, Forbid: 0}
	result := DB.Create(u)
	if result.Error != nil {
		return nil, false, "username has been used"
	}
	return u, true, "ok"
}

func UpdateUser(userId uint, attr string, attrVal string) (ok bool) {
	result := DB.Model(&User{
		Model: gorm.Model{
			ID: userId,
		},
	}).Update(attr, attrVal)
	return result.Error == nil
}
func BanUser(userId uint) (ok bool) {
	result := DB.Model(&User{
		Model: gorm.Model{
			ID: userId,
		},
	}).Update("forbid", 1)
	return result.Error == nil
}
