package orm

import (
	"cilicili/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() {
	db, err := gorm.Open(mysql.Open(config.Conf.MysqlConfig.DNS), &gorm.Config{})
	DB = db
	err = DB.AutoMigrate(&User{}, &Video{}, &Comment{}, &Good{}, &Collect{}, &BulletChat{})
	if err != nil {
		panic("db err")
	}
}
