package sysmodel

import (
	"errors"
	"log"
	"time"

	"github.com/tianye3017/gin-admin-backend/db"
)

// SysMenu 菜单表
type SysMenu struct {
	Id            uint      `json:"id" xorm:"pk autoincr"`
	ParentId      uint      `json:"parent_id" xorm:"notnull"`
	Name          string    `json:"name" xorm:"varchar(32) notnull"`
	Uri           string    `json:"uri" xorm:"varchar(64)"`
	Memo          string    `json:"memo" xorm:"varchar(64)"`
	Sequence      int       `json:"sequence" xorm:"notnull"`
	MenuType      int       `json:"menu_type" xorm:"tinyint(1) notnull"`
	RequestMethod string    `json:"request_method" xorm:"varchar(10)"`
	Status        int       `json:"status" xorm:"tinyint(1) notnull"`
	Code          string    `json:"code" xorm:"varchar(32) notnull"`
	Icon          string    `json:"icon" xorm:"varchar(32)"`
	CreatedAt     time.Time `json:"created_at" xorm:"created"`
	UpdatedAt     time.Time `json:"updated_at" xorm:"updated"`
}

func init() {
	MenuModel := new(SysMenu)
	if err := db.DB.Sync(MenuModel); err != nil {
		log.Fatal("数据表同步失败:", err)
	}
}

// GetAllMenu 获取全部菜单
func (m *SysMenu) GetAllMenu() (a []SysMenu, e error) {
	allMenus := make([]SysMenu, 0)
	err := db.DB.Asc("parent_id").Asc("sequence").Find(&allMenus)
	return allMenus, err
}

// Delete 根据传入id删除数据以及关联表数据
func (m *SysMenu) Delete(id uint) (err error) {
	// 首先删除关联表的相关数据
	toDelRoleMenu := SysRoleMenu{MenuId: id}
	session := db.DB.NewSession()
	defer session.Close()
	err = session.Begin()
	_, err = db.DB.Delete(toDelRoleMenu)
	if err != nil {
		session.Rollback()
		return err
	}
	// 删除菜单
	affected, err := db.DB.Id(id).Delete(m)
	if err == nil && affected == 0 {
		session.Rollback()
		err = errors.New("角色删除失败")
		return err
	}
	err = session.Commit()
	if err != nil {
		session.Rollback()
		err = errors.New("操作失败")
		return err
	}
	return err
}
