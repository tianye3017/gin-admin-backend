package api

import (
	"strconv"

	"github.com/tianye3017/gin-admin-backend/db"
	"github.com/tianye3017/gin-admin-backend/model/sysmodel"
	"github.com/tianye3017/gin-admin-backend/service"
	"github.com/gin-gonic/gin"
)

type menuEditStruct struct {
	Id uint `json:"id" required`
	menuCreateStruct
}

type menuCreateStruct struct {
	ParentId      uint   `json:"parent_id" required`
	Name          string `json:"name" binding:"required"`
	Uri           string `json:"uri" binding:"required"`
	Code          string `json:"code" binding:"required"`
	Memo          string `json:"memo"`
	Sequence      int    `json:"sequence" binding:"required,numeric"`
	MenuType      int    `json:"menu_type" binding:"required,numeric"`
	RequestMethod string `json:"request_method" binding:"required"`
	Status        int    `json:"status" binding:"required,numeric"`
	Icon          string `json:"icon"`
}

func MenuData(c *gin.Context) {
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
		menuName := c.Query("key")
		whereSql := "WHERE id !=1"
		if menuName != "" {
			whereSql = whereSql + " AND name LIKE '%" + menuName + "%' "
		}
		parentId := c.Query("parent_id")
		if parentId != "0" && parentId != "" {
			whereSql = whereSql + " AND parent_id = " + parentId
		}
		menuType := c.Query("type")
		if menuType != "" {
			whereSql = whereSql + " AND menu_type = " + menuType
		}
		var total int64
		total, err = db.DB.Sql("SELECT COUNT(*) FROM sys_menu " + whereSql).Count()
		if err != nil {
			service.ResFail(c, err.Error())
			return
		}
		menus := make([]sysmodel.SysMenu, 0)
		sql := "SELECT * FROM sys_menu " + whereSql + " ORDER BY " + sort + " LIMIT " + strconv.Itoa(limit) + " OFFSET " + strconv.Itoa((page-1)*limit) + ";"
		err = db.DB.Sql(sql).Find(&menus)
		// err = db.DB.Asc("sequence").Asc("id").Limit(limit, (page-1)*limit).Find(&roles)
		if err != nil {
			service.ResFail(c, err.Error())
			return
		}
		service.ResSuccess(c, "ok", gin.H{"total": total, "items": menus})
		return
	} else {
		menu := new(sysmodel.SysMenu)
		has, err := db.DB.Where("id = ?", c.Query("id")).Get(menu)
		if err != nil {
			service.ResFail(c, err.Error())
			return
		}
		if !has {
			service.ResFail(c, "数据获取失败")
			return
		}
		service.ResSuccess(c, "ok", &menu)
		return
	}
}

func MenuAll(c *gin.Context) {
	menuModel := new(sysmodel.SysMenu)
	allMenus, _ := menuModel.GetAllMenu()
	service.ResSuccess(c, "ok", allMenus)
}

func MenuEdit(c *gin.Context) {
	var mes menuEditStruct
	err := c.BindJSON(&mes)
	if err != nil {
		service.ResFail(c, "传值错误")
		return
	}
	menuModel := new(sysmodel.SysMenu)
	menuModel.ParentId = mes.ParentId
	menuModel.Name = mes.Name
	menuModel.Uri = mes.Uri
	menuModel.Code = mes.Code
	menuModel.Memo = mes.Memo
	menuModel.Sequence = mes.Sequence
	menuModel.MenuType = mes.MenuType
	menuModel.RequestMethod = mes.RequestMethod
	menuModel.Status = mes.Status
	menuModel.Icon = mes.Icon
	affected, err := db.DB.Id(mes.Id).Update(menuModel)
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

func MenuDelete(c *gin.Context) {
	var ids []uint
	err := c.Bind(&ids)
	if err != nil {
		service.ResFail(c, err.Error())
		return
	}
	menuModel := new(sysmodel.SysMenu)
	for _, id := range ids {
		err = menuModel.Delete(id)
		if err != nil {
			service.ResFail(c, err.Error())
			return
		}
	}
	service.ResSuccess(c, "菜单删除成功", "")
	return
}

func MenuCreate(c *gin.Context) {
	var mcs menuCreateStruct
	err := c.BindJSON(&mcs)
	if err != nil {
		service.ResFail(c, "传值错误")
		return
	}
	menuModel := new(sysmodel.SysMenu)
	menuModel.ParentId = mcs.ParentId
	menuModel.Name = mcs.Name
	menuModel.Uri = mcs.Uri
	menuModel.Code = mcs.Code
	menuModel.Memo = mcs.Memo
	menuModel.Sequence = mcs.Sequence
	menuModel.MenuType = mcs.MenuType
	menuModel.RequestMethod = mcs.RequestMethod
	menuModel.Status = mcs.Status
	menuModel.Icon = mcs.Icon
	affected, err := db.DB.Insert(menuModel)
	if err != nil {
		service.ResFail(c, err.Error())
		return
	}
	if affected == 0 {
		service.ResFail(c, "新增失败")
		return
	}
	service.ResSuccess(c, "新增成功", gin.H{"id": menuModel.Id})
	return
}
