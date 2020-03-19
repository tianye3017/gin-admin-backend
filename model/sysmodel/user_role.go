package sysmodel

import (
	"log"
	"time"

	"github.com/tianye3017/gin-admin-backend/db"
)

// SysUserRole 用户角色关联表
type SysUserRole struct {
	Id        uint      `json:"id" xorm:"pk autoincr"`
	UserId    uint      `json:"user_id" xorm:"notnull"`
	RoleId    uint      `json:"role_id" xorm:"notnull"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
}

func init() {
	UserRoleModel := new(SysUserRole)
	if err := db.DB.Sync(UserRoleModel); err != nil {
		log.Fatal("数据表同步失败:", err)
	}
}

// GetUserRoleId 根据用户id获取用户的角色
func (m *SysUserRole) GetUserRoleId(userId uint) (r []uint, e error) {
	userRole := new([]SysUserRole)
	e = db.DB.Cols("role_id").Where("user_id = ?", userId).Find(userRole)
	if e != nil {
		return r, e
	}
	roleIds := []uint{}
	for _, v := range *userRole {
		roleIds = append(roleIds, v.RoleId)
	}
	return roleIds, e
}
