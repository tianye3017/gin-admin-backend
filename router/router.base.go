package router

import (
	"github.com/tianye3017/gin-admin-backend/controller/api"
	"github.com/tianye3017/gin-admin-backend/engine"
)

func init() {
	engine.Router.POST("/user/login", api.UserLogin)
	engine.Router.GET("/test", api.Test)
}
