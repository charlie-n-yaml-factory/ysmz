package server

import (
	"time"

	"github.com/charlie-n-yaml-factory/ysmz/ysmz-backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))

	healthController := new(controllers.HealthController)
	oauthGoogleController := new(controllers.OAuthGoogleController)

	r.GET("/health", healthController.Status)

	v1 := r.Group("/oauth/google")
	{
		v1.GET("/state", oauthGoogleController.GetState)
		v1.GET("/callback", oauthGoogleController.GetTokens)
		v1.GET("/user-info", oauthGoogleController.GetUserInfo)
	}

	return r
}
