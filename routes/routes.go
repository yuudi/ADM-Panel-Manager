package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuudi/ADM-Panel-manager/containers"
	"github.com/yuudi/ADM-Panel-manager/panel"
)

const cookieKey = "adm-token"

func requireLogin(c *gin.Context) (*panel.UserToken, bool) {
	cookie, err := c.Cookie(cookieKey)
	if err != nil {
		return nil, false
	}
	token := panel.UserToken{}
	err = token.ParseJwtString(cookie)
	if err != nil {
		return nil, false
	}
	return &token, true
}

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api.POST("/login/pre", func(c *gin.Context) {
		var loginPreRequest panel.LoginPreRequest
		c.BindJSON(&loginPreRequest)
		loginRequestToken := panel.NewLoginRequestToken(loginPreRequest.Username, c.ClientIP())
		loginRequestTokenString, err := loginRequestToken.ToJwtString()
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
		c.JSON(200, gin.H{
			"login_request_token": loginRequestTokenString,
		})
	})
	api.POST("/login", func(c *gin.Context) {
		var loginRequest panel.LoginRequest
		c.BindJSON(&loginRequest)
		var loginRequestToken panel.LoginRequestToken
		err := loginRequestToken.ParseJwtString(loginRequest.LoginRequestTokenString)
		if err != nil {
			c.JSON(401, gin.H{
				"code":    401001,
				"message": "Unauthorized",
			})
			return
		}
		if loginRequestToken.RemoteAddress != c.ClientIP() {
			c.JSON(401, gin.H{
				"code":    401002,
				"message": "Unauthorized",
			})
			return
		}
		if loginRequestToken.Timestamp+60 < time.Now().Unix() {
			c.JSON(401, gin.H{
				"code":    401003,
				"message": "Unauthorized",
			})
			return
		}
		ok, err := loginRequest.VerifyPassword(panel.GetPanelInstance().Config.Auth.PasswordHmacHex) //TODO: password based on username
		if !ok {
			c.JSON(401, gin.H{
				"code":    401004,
				"message": err.Error(),
			})
			return
		}
		userToken := panel.UserToken{
			Username:   loginRequestToken.Username,
			ExpireTime: time.Now().Add(time.Hour * 24 * 7).Unix(),
		}
		userTokenString, err := userToken.ToJwtString()
		if err != nil {
			c.JSON(500, gin.H{
				"code":    500001,
				"message": "Internal Server Error",
			})
			return
		}
		c.SetCookie(cookieKey, userTokenString, 60*60*24*7, "/", "", false, true)
		c.JSON(200, gin.H{
			"message": "Authorized",
			"code":    0,
		})
	})
	api.GET("/authping", func(c *gin.Context) {
		token, ok := requireLogin(c)
		if !ok {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		c.JSON(200, gin.H{
			"message":   "Authorized",
			"user_name": token.Username,
		})
	})
	api.GET("/containers", func(c *gin.Context) {
		_, ok := requireLogin(c)
		if !ok {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		_, err := containers.ListContainers()
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
		c.JSON(200, gin.H{
			"containers": "", //TODO: Implement
		})
	})
}
