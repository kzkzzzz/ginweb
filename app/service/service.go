package service

import (
	"ginweb/app/conf"
	"ginweb/app/dao"
)

type Service struct {
	dao *dao.Dao
}

func New(conf *conf.Config) (s *Service) {
	s = &Service{
		dao: dao.New(conf.Data),
	}
	return
}
