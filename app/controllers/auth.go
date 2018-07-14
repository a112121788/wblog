package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/a112121788/wblog/app/helpers"
	"net/http"
	"github.com/a112121788/wblog/config"
)

func AuthGet(c *gin.Context) {
	authType := c.Param("authType")

	session := sessions.Default(c)
	uuid := helpers.UUID()
	session.Delete(SESSION_GITHUB_STATE)
	session.Set(SESSION_GITHUB_STATE, uuid)
	session.Save()

	authurl := "/signin"
	switch authType {
	case "github":
		authurl = fmt.Sprintf(config.GetConfiguration().GithubAuthUrl, config.GetConfiguration().GithubClientId, uuid)
	case "weibo":
	case "qq":
	case "wechat":
	case "oschina":
	default:
	}
	c.Redirect(http.StatusFound, authurl)
}
