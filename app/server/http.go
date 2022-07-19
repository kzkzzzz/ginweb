package server

import (
	"ginweb/app/conf"
	"ginweb/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	svc    *service.Service
	engine *gin.Engine
)

func NewHttp(c *conf.Config, s *service.Service) (h *http.Server) {
	svc = s
	h = &http.Server{
		Addr: c.Server.Addr,
	}
	engine = gin.New()
	engine.Use(gin.Logger(), Recovery())
	router()
	h.Handler = engine
	return
}

func router() {
	engine.GET("/", ApiWrap(index))
}
