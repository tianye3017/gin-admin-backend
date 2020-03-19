package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tianye3017/gin-admin-backend/db"
	"github.com/tianye3017/gin-admin-backend/model/sysmodel"
)

func Test(c *gin.Context) {
	sql_2_1 := "select * from sys_user where id=9"
	var aaa []sysmodel.SysUser
	db.DB.Sql(sql_2_1).Find(&aaa)
	fmt.Println(aaa)
}
