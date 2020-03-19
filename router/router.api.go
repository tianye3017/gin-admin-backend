package router

import (
	"github.com/tianye3017/gin-admin-backend/controller/api"
	"github.com/tianye3017/gin-admin-backend/engine"
	"github.com/tianye3017/gin-admin-backend/middleware"
)

func init() {
	// 登录用户
	userGroup := engine.Router.Group("user").Use(middleware.JWTAuth())
	{
		userGroup.GET("info", api.UserInfo)
		userGroup.POST("editpwd", api.EditPassword)
		userGroup.POST("logout", api.Logout)
	}

	// 角色管理
	roleGroup := engine.Router.Group("role").Use(middleware.JWTAuth()).Use(middleware.AuthVerify())
	{
		roleGroup.GET("", api.RoleData)
		roleGroup.PUT("", api.RoleEdit)
		roleGroup.POST("", api.RoleCreate)
		roleGroup.DELETE("", api.RoleDelete)
		roleGroup.GET("rolemenulist", api.RoleMenuList) // 获取全部权限节点
		roleGroup.POST("setrole", api.SetRole)          // 设置角色
	}

	// 菜单管理
	menuGroup := engine.Router.Group("menu").Use(middleware.JWTAuth())
	{
		menuGroup.GET("", api.MenuData)
		menuGroup.GET("all", api.MenuAll)
		menuGroup.PUT("", api.MenuEdit)
		menuGroup.DELETE("", api.MenuDelete)
		menuGroup.POST("", api.MenuCreate)
	}

	// 后台用户管理
	backUserGroup := engine.Router.Group("backuser").Use(middleware.JWTAuth())
	{
		backUserGroup.GET("", api.BackUserData)
		backUserGroup.PUT("", api.BackUserEdit)
		backUserGroup.DELETE("", api.BackUserDelete)
		backUserGroup.POST("", api.BackUserCreate)
		backUserGroup.GET("rolelist", api.BackUserRoleList) // 获取后台用户角色列表
		backUserGroup.PUT("setrole", api.BackUserSetRole)   // 设置后台用户权限
	}
}
