package service

import (
	"ginweb/app/model"
)

func (s *Service) Index(arg *model.UserSearch) ([]*model.User, error) {
	return s.dao.ListUser(arg)
}
