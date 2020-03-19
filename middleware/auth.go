package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tianye3017/gin-admin-backend/config"
	"github.com/tianye3017/gin-admin-backend/db"
	"github.com/tianye3017/gin-admin-backend/model/sysmodel"
	"github.com/tianye3017/gin-admin-backend/service"
)

// AuthVerify 鉴权
func AuthVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		claimsData := claims.(*CustomClaims)
		has := false
		var uri, method string
		if claimsData.Id == config.SysConfig.System.SuperId {
			has = true
		} else {
			// 获取请求的URI
			uri = service.GetStrBefore(c.Request.URL.RequestURI(), "?")
			// 获取请求方法
			method = c.Request.Method
			//获取用户的角色
			syr := new(sysmodel.SysUserRole)
			roleIds, err := syr.GetUserRoleId(claimsData.Id)
			if err != nil {
				service.ResFail(c, err.Error())
				return
			}
			has, _ = db.DB.Cols("sys_menu.id").Table("sys_menu").Join("INNER", "sys_role_menu", "sys_menu.id=sys_role_menu.menu_id").Join("INNER", "sys_role", "sys_role.id=sys_role_menu.role_id").Where("sys_menu.uri = ?", uri).Where("sys_menu.request_method = ?", method).In("sys_role.id", roleIds).Get(new(sysmodel.SysMenu))
		}
		if has {
			c.Next()
		} else {
			menuModel := new(sysmodel.SysMenu)
			has, _ := db.DB.Cols("name").Where("uri = ?", uri).Where("request_method = ?", method).Get(menuModel)
			var resStr string
			if has {
				resStr = "缺少权限：" + menuModel.Name
			} else {
				resStr = "未定义权限：" + uri + "（" + method + "），请联系管理员添加后授权。"
			}
			service.ResFail(c, resStr)
			return
		}
	}
}
