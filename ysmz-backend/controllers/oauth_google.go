package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/charlie-n-yaml-factory/ysmz/ysmz-backend/config"
	"github.com/charlie-n-yaml-factory/ysmz/ysmz-backend/databases"
	"github.com/charlie-n-yaml-factory/ysmz/ysmz-backend/services"
)

type OAuthGoogleController struct{}

func (o *OAuthGoogleController) GetState(c *gin.Context) {
	redisDB := databases.RedisDB()

	state, err := services.SaveOAuthGoogleState(redisDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not retrieve google oauth state",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"state": state,
	})
}

func (o *OAuthGoogleController) GetTokens(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "code query parameter is required",
		})
		return
	}

	state := c.Query("state")
	if state == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "state query parameter is required",
		})
		return
	}

	conf := config.Config()
	redisDB := databases.RedisDB()

	oauthGoogleTokens, err := services.GetOAuthGoogleTokens(code, state, conf, redisDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.SetCookie("ysmz_google_access_token", oauthGoogleTokens.AccessToken, 300, "/", "", true, true)

	clientOrigin := conf.ClientOrigin
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not parse client origin",
		})
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, clientOrigin)
}

func (o *OAuthGoogleController) GetUserInfo(c *gin.Context) {
	// get tokens from cookies
	googleAccessToken, err := c.Cookie("ysmz_google_access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "could not retrieve google access token",
		})
		return
	}

	// get user info
	userInfo, err := services.GetOAuthGoogleUserInfo(googleAccessToken)
	c.JSON(http.StatusOK, gin.H{
		"user_given_name": userInfo["given_name"],
	})
}
