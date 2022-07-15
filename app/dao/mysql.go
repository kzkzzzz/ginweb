package dao

import (
	"ginweb/app/model"
)

func (d *Dao) ListUser(arg *model.UserSearch) (res []*model.User, err error) {
	err = d.DB.Table("user").
		Where("id > ?", arg.Id).
		Limit(15).Find(&res).Error
	return
}
