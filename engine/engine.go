package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/tianye3017/gin-admin-backend/middleware"
)

var Router *gin.Engine

func init() {
	Router = gin.New()
	Router.Use(middleware.Cors())
}
