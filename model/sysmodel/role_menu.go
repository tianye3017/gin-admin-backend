package sysmodel

import (
	"log"
	"time"

	"github.com/tianye3017/gin-admin-backend/db"
)

// SysRoleMenu 角色菜单关联表
type SysRoleMenu struct {
	Id        uint      `json:"id" xorm:"pk autoincr"`
	RoleId    uint      `json:"role_id" xorm:"notnull"`
	MenuId    uint      `json:"menu_id" xorm:"notnull"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
}

func init() {
	RoleMenuModel := new(SysRoleMenu)
	if err := db.DB.Sync(RoleMenuModel); err != nil {
		log.Fatal("数据表同步失败:", err)
	}
}
