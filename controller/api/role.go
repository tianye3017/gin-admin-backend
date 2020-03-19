package api

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tianye3017/gin-admin-backend/db"
	"github.com/tianye3017/gin-admin-backend/model/sysmodel"
	"github.com/tianye3017/gin-admin-backend/service"
)

type roleCreateStruct struct {
	Name     string `json:"name" binding:"required"`
	Sequence int    `json:"sequence" binding:"required,numeric"`
	Memo     string
}

type roleEditStruct struct {
	Id uint `json:"id" required`
	roleCreateStruct
}

func RoleData(c *gin.Context) {
	if c.Query("id") == "" {
		if c.Query("page") == "" {
			roles, err := new(sysmodel.SysRole).GetAllRole()
			if err != nil {
				service.ResFail(c, err.Error())
				return
			}
			service.ResSuccess(c, "ok", roles)
			return
		} else {
			page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
			limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
			// total, err := db.DB.Count(new(sysmodel.SysRole))
			// if err != nil {
			// 	service.ResFail(c, err.Error())
			// 	return
			// }
			var err error
			var sort string
			err, sort = service.ResolveSorKey(c.DefaultQuery("sort", "+id"))
			if err != nil {
				service.ResFail(c, err.Error())
				return
			}
			roleName := c.Query("key")
			whereSql := " WHERE 1=1"
			if roleName != "" {
				whereSql = whereSql + " AND name LIKE '%" + roleName + "%'"
			}
			var total int64
			total, err = db.DB.Sql("SELECT COUNT(*) FROM sys_role " + whereSql).Count()
			if err != nil {
				service.ResFail(c, err.Error())
				return
			}
			roles := make([]sysmodel.SysRole, 0)
			sql := "SELECT * FROM sys_role " + whereSql + " ORDER BY " + sort + " LIMIT " + strconv.Itoa(limit) + " OFFSET " + strconv.Itoa((page-1)*limit) + ";"
			err = db.DB.Sql(sql).Find(&roles)
			if err != nil {
				service.ResFail(c, err.Error())
				return
			}
			service.ResSuccess(c, "ok", gin.H{"total": total, "items": roles})
			return
		}
	} else {
		role := new(sysmodel.SysRole)
		has, err := db.DB.Where("id = ?", c.Query("id")).Get(role)
		if err != nil {
			service.ResFail(c, err.Error())
			return
		}
		if !has {
			service.ResFail(c, "数据获取失败")
			return
		}
		service.ResSuccess(c, "ok", &role)
		return
	}
}

func RoleCreate(c *gin.Context) {
	var cae roleCreateStruct
	err := c.BindJSON(&cae)
	if err != nil {
		service.ResFail(c, "传值错误")
		return
	}
	roleModel := new(sysmodel.SysRole)
	roleModel.Name = cae.Name
	roleModel.Sequence = cae.Sequence
	roleModel.Memo = cae.Memo
	roleModel.Status = 1
	affected, err := db.DB.Insert(roleModel)
	if err != nil {
		service.ResFail(c, err.Error())
		return
	}
	if affected == 0 {
		service.ResFail(c, "新增失败")
		return
	}
	service.ResSuccess(c, "新增成功", gin.H{"id": roleModel.Id})
	return
}

func RoleDelete(c *gin.Context) {
	// roleModel := new(sysmodel.SysRole)
	// affected, err := db.DB.Id(c.Param("id")).Delete(roleModel)
	// if err != nil {
	// 	service.ResFail(c, err.Error())
	// 	return
	// }
	// if affected == 0 {
	// 	service.ResFail(c, "删除失败")
	// 	return
	// }
	// service.ResSuccess(c, "删除成功", "")
	// return
	var ids []uint
	err := c.Bind(&ids)
	if err != nil {
		service.ResFail(c, err.Error())
		return
	}
	roleModel := new(sysmodel.SysRole)
	for _, id := range ids {
		err = roleModel.Delete(id)
		if err != nil {
			service.ResFail(c, err.Error())
			return
		}
	}
	service.ResSuccess(c, "角色删除成功", "")
	return
}

func RoleEdit(c *gin.Context) {
	var cae roleEditStruct
	err := c.BindJSON(&cae)
	if err != nil {
		service.ResFail(c, "传值错误")
		return
	}
	roleModel := new(sysmodel.SysRole)
	roleModel.Name = cae.Name
	roleModel.Sequence = cae.Sequence
	roleModel.Memo = cae.Memo
	affected, err := db.DB.Id(cae.Id).Update(roleModel)
	if err != nil {
		service.ResFail(c, err.Error())
		return
	}
	if affected == 0 {
		service.ResFail(c, "修改失败")
		return
	}
	service.ResSuccess(c, "修改成功", "")
	return
}

func RoleMenuList(c *gin.Context) {
	var getRoleId string
	getRoleId = c.Query("roleid")
	if getRoleId == "" {
		service.ResFail(c, "参数错误")
		return
	}
	roleId, err := strconv.Atoi(getRoleId)
	if err != nil {
		log.Fatal("类型转换错误:", err)
	}
	roleList := make([]sysmodel.SysMenu, 0)
	db.DB.Table("sys_menu").Join("INNER", "sys_role_menu", "sys_menu.id=sys_role_menu.menu_id").Join("INNER", "sys_role", "sys_role.id=sys_role_menu.role_id").Where("sys_role.id = ?", roleId).Find(&roleList)
	service.ResSuccess(c, "ok", roleList)

}

func SetRole(c *gin.Context) {
	var getRoleId string
	getRoleId = c.Query("roleid")
	if getRoleId == "" {
		service.ResFail(c, "参数错误")
		return
	}
	roleIdInt, err := strconv.Atoi(getRoleId)
	if err != nil {
		log.Fatal("类型转换错误:", err)
	}
	roleId := uint(roleIdInt)
	var menuIds []uint
	err = c.BindJSON(&menuIds)
	if err != nil {
		service.ResFail(c, "参数错误")
		return
	}
	toDelRoleMenu := sysmodel.SysRoleMenu{RoleId: roleId}
	session := db.DB.NewSession()
	defer session.Close()
	err = session.Begin()
	_, err = db.DB.Delete(toDelRoleMenu)
	if err != nil {
		session.Rollback()
		service.ResFail(c, err.Error())
		return
	}
	toInsertRoleMenu := make([]*sysmodel.SysRoleMenu, 0)
	for i, v := range menuIds {
		toInsertRoleMenu = append(toInsertRoleMenu, new(sysmodel.SysRoleMenu))
		toInsertRoleMenu[i].RoleId = roleId
		toInsertRoleMenu[i].MenuId = v
	}
	_, err = db.DB.Insert(&toInsertRoleMenu)
	if err != nil {
		session.Rollback()
		service.ResFail(c, err.Error())
		return
	}
	err = session.Commit()
	if err != nil {
		session.Rollback()
		service.ResFail(c, "操作失败")
		return
	}
	service.ResSuccess(c, "ok", "操作成功")
	return
}
