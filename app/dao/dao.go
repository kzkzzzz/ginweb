package dao

import (
	"ginweb/app/conf"
	"ginweb/common/mysql"
	"gorm.io/gorm"
)

type Dao struct {
	DB *gorm.DB
}

func New(conf *conf.Data) (d *Dao) {
	d = &Dao{
		DB: mysql.NewDB(conf.Mysql),
	}
	return
}
