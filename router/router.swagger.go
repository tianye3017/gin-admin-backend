package router

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/tianye3017/gin-admin-backend/docs"
	"github.com/tianye3017/gin-admin-backend/engine"
)

func init() {
	engine.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
