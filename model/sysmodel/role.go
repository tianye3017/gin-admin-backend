package sysmodel

import (
	"errors"
	"log"
	"time"

	"github.com/tianye3017/gin-admin-backend/db"
)

// SysRole 角色表
type SysRole struct {
	Id        uint      `json:"id" xorm:"pk autoincr"`
	Name      string    `json:"name" xorm:"varchar(32) notnull"`
	Memo      string    `json:"memo" xorm:"varchar(64)"`
	Sequence  int       `json:"sequence" xorm:"notnull"`
	Status    int       `json:"status" xorm:"tinyint(1) notnull"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
}

func init() {
	RoleModel := new(SysRole)
	if err := db.DB.Sync(RoleModel); err != nil {
		log.Fatal("数据表同步失败:", err)
	}
}

// Delete 根据传入id删除数据，删除前判断是否有用户使用。
func (m *SysRole) Delete(id uint) (err error) {
	hasUserUse, err := db.DB.Where("role_id = ?", id).Get(new(SysUserRole))
	if err == nil {
		if hasUserUse == true {
			err = errors.New("有用户正在使用角色禁止删除")
		} else {
			affected, err := db.DB.Id(id).Delete(m)
			if err == nil && affected == 0 {
				err = errors.New("角色删除失败")
			}
		}
	}
	return err
}

// GetAllRole 获取全部角色
func (m *SysRole) GetAllRole() (a []SysRole, e error) {
	allRoles := make([]SysRole, 0)
	err := db.DB.Asc("sequence").Find(&allRoles)
	return allRoles, err
}
