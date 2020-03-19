package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/tianye3017/gin-admin-backend/config"
)

var DB *xorm.Engine

func init() {
	DB, _ = xorm.NewEngine("mysql", config.SysConfig.Database.Username+":"+config.SysConfig.Database.Password+"@tcp("+config.SysConfig.Database.Path+")/"+config.SysConfig.Database.Dbname+"?"+config.SysConfig.Database.Config)
	// DB.ShowSQL(true)
}
