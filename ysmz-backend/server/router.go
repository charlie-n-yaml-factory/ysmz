package server

import (
	"github.com/charlie-n-yaml-factory/ysmz/ysmz-backend/controllers"
	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	r := gin.Default()

	health := new(controllers.HealthController)

	r.GET("/health", health.Status)

	return r
}
