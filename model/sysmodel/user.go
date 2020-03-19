package sysmodel

import (
	"errors"
	"log"
	"time"

	"github.com/tianye3017/gin-admin-backend/db"
)

// SysUser 用户表
type SysUser struct {
	Id        uint      `json:"id" xorm:"pk autoincr"`
	Username  string    `json:"username" xorm:"notnull unique"`
	Password  string    `json:"password" xorm:"notnull"`
	NickName  string    `json:"nick_name" xorm:"varchar(64)"`
	Avatar    string    `json:"avatar"`
	Email     string    `json:"email" xorm:"varchar(64)"`
	Phone     string    `json:"phone" xorm:"varchar(20)"`
	Memo      string    `json:"memo" xorm:"varchar(100)"`
	Status    int       `json:"status" xorm:"tinyint(1) notnull"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
}

func init() {
	UserModel := new(SysUser)
	if err := db.DB.Sync(UserModel); err != nil {
		log.Fatal("数据表同步失败:", err)
	}
}

func (m *SysUser) Delete(id uint) (err error) {
	// 首先删除关联表的相关数据
	toDelUserRole := SysUserRole{UserId: id}
	session := db.DB.NewSession()
	defer session.Close()
	err = session.Begin()
	_, err = db.DB.Delete(toDelUserRole)
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
