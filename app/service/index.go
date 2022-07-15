package service

import (
	"ginweb/app/model"
	"ginweb/common/errm"
)

func (s *Service) Index(arg *model.UserSearch) ([]*model.User, error) {
	user, err := s.dao.ListUser(arg)
	if err != nil {
		return nil, errm.DbError(err.Error())
	}
	return user, nil
}
