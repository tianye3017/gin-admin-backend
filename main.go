package main

import (
	_ "github.com/tianye3017/gin-admin-backend/config"
	_ "github.com/tianye3017/gin-admin-backend/db"
	"github.com/tianye3017/gin-admin-backend/engine"

	// _ "github.com/tianye3017/gin-admin-backend/model"
	_ "github.com/tianye3017/gin-admin-backend/router"
)

func main() {
	engine.Router.Run()
}
