package router

import (
	"auth0-sample/platform/authenticator"
	"auth0-sample/platform/middleware"
	"auth0-sample/web/app/callback"
	"auth0-sample/web/app/login"
	"auth0-sample/web/app/logout"
	"auth0-sample/web/app/user"
	"encoding/gob"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	gob.Register(map[string]interface{}{})
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", middleware.IsAuthenticated, user.Handler)
	router.GET("/logout", logout.Handler)

	return router

}
