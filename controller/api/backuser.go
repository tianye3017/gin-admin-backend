package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tianye3017/gin-admin-backend/db"
	"github.com/tianye3017/gin-admin-backend/model/sysmodel"
	"github.com/tianye3017/gin-admin-backend/service"
)

type backUserEditStruct struct {
	Id uint `json:"id" required`
	backUserOtherStruct
}

type backUserCreateStruct struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	backUserOtherStruct
}

type backUserOtherStruct struct {
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Status   int    `json:"status" binding:"required,numeric"`
	Memo     string `json:"memo"`
}

func BackUserData(c *gin.Context) {
	if c.Query("id") == "" {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
		var err error
		var sort string
		err, sort = service.ResolveSorKey(c.DefaultQuery("sort", "+id"))
		if err != nil {
			service.ResFail(c, err.Error())
			return
		}
		whereSql := " WHERE 1=1"
		backUserName := c.Query("key")
		if backUserName != "" {
			whereSql = whereSql + " AND username LIKE '%" + backUserName + "%' "
		}
		status := c.Query("status")
		if status != "" {
			whereSql = whereSql + " AND status = " + status
		}
		var total int64
		total, err = db.DB.Sql("SELECT COUNT(*) FROM sys_user " + whereSql).Count()
		if err != nil {
			service.ResFail(c, err.Error())
			return
		}
		backUsers := make([]sysmodel.SysUser, 0)
		sql := "SELECT * FROM sys_user " + whereSql + " ORDER BY " + sort + " LIMIT " + strconv.Itoa(limit) + " OFFSET " + strconv.Itoa((page-1)*limit) + ";"
		err = db.DB.Sql(sql).Find(&backUsers)
		if err != nil {
			service.ResFail(c, err.Error())
			return
		}
		service.ResSuccess(c, "ok", gin.H{"total": total, "items": backUsers})
		return
	} else {
		backUser := new(sysmodel.SysUser)
		has, err := db.DB.Where("id = ?", c.Query("id")).Get(backUser)
		if err != nil {
			service.ResFail(c, err.Error())
			return
		}
		if !has {
			service.ResFail(c, "数据获取失败")
			return
		}
		service.ResSuccess(c, "ok", &backUser)
		return
	}
}

func BackUserEdit(c *gin.Context) {
	var bes backUserEditStruct
	err := c.BindJSON(&bes)
	if err != nil {
		service.ResFail(c, "传值错误")
		return
	}
	backUserModel := new(sysmodel.SysUser)
	backUserModel.NickName = bes.NickName
	backUserModel.Email = bes.Email
	backUserModel.Phone = bes.Phone
	backUserModel.Status = bes.Status
	backUserModel.Memo = bes.Memo
	affected, err := db.DB.Id(bes.Id).Update(backUserModel)
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

func BackUserDelete(c *gin.Context) {
	var ids []uint
	err := c.Bind(&ids)
	if err != nil {
		service.ResFail(c, err.Error())
		return
	}
	backUserModel := new(sysmodel.SysUser)
	for _, id := range ids {
		err = backUserModel.Delete(id)
		if err != nil {
			service.ResFail(c, err.Error())
			return
		}
	}
	service.ResSuccess(c, "菜单删除成功", "")
	return
}

func BackUserCreate(c *gin.Context) {
	var bcs backUserCreateStruct
	err := c.BindJSON(&bcs)
	if err != nil {
		service.ResFail(c, "传值错误")
		return
	}
	backUserModel := new(sysmodel.SysUser)
	backUserModel.Username = bcs.Username
	backUserModel.NickName = bcs.NickName
	backUserModel.Email = bcs.Email
	backUserModel.Phone = bcs.Phone
	backUserModel.Status = bcs.Status
	backUserModel.Memo = bcs.Memo
	affected, err := db.DB.Insert(backUserModel)
	if err != nil {
		service.ResFail(c, err.Error())
		return
	}
	if affected == 0 {
		service.ResFail(c, "新增失败")
		return
	}
	service.ResSuccess(c, "新增成功", gin.H{"id": backUserModel.Id})
	return
}

func BackUserRoleList(c *gin.Context) {
	userId := c.Query("userid")
	if userId == "" {
		service.ResFail(c, "传值错误")
		return
	}
	syr := new(sysmodel.SysUserRole)
	userIdInt, _ := strconv.Atoi(userId)
	roleIds, err := syr.GetUserRoleId(uint(userIdInt))
	if err != nil {
		service.ResFail(c, err.Error())
		return
	}
	service.ResSuccess(c, "ok", roleIds)
}

func BackUserSetRole(c *gin.Context) {
	getUserId := c.Query("userid")
	if getUserId == "" {
		service.ResFail(c, "传值错误")
		return
	}
	userIdInt, _ := strconv.Atoi(getUserId)
	var roleIds []uint
	err := c.BindJSON(&roleIds)
	if err != nil {
		service.ResFail(c, "参数错误")
		return
	}
	userId := uint(userIdInt)
	toDelUserRole := sysmodel.SysUserRole{UserId: userId}
	session := db.DB.NewSession()
	defer session.Close()
	err = session.Begin()
	_, err = db.DB.Delete(toDelUserRole)
	if err != nil {
		session.Rollback()
		service.ResFail(c, err.Error())
		return
	}
	toInsertUserRole := make([]*sysmodel.SysUserRole, 0)
	for i, v := range roleIds {
		toInsertUserRole = append(toInsertUserRole, new(sysmodel.SysUserRole))
		toInsertUserRole[i].RoleId = v
		toInsertUserRole[i].UserId = userId
	}
	_, err = db.DB.Insert(&toInsertUserRole)
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
