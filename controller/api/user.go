package api

import (
	"time"

	"github.com/ahmetb/go-linq"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tianye3017/gin-admin-backend/config"
	"github.com/tianye3017/gin-admin-backend/db"
	"github.com/tianye3017/gin-admin-backend/middleware"
	"github.com/tianye3017/gin-admin-backend/model/sysmodel"
	"github.com/tianye3017/gin-admin-backend/service"
	"github.com/tianye3017/gin-admin-backend/tools"
)

type registAndLoginStuct struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type editUserInfoStruct struct {
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Memo     string `json:"memo"`
}

type editPasswordStruct struct {
	OldPassword      string `json:"old_password" binding:"required"`
	NewPassword      string `json:"new_password" binding:"required,min=6,max=10"`                           //长度大于等于6,小于等于10,与旧密码不同
	NewPasswordAgain string `json:"new_password_again" binding:"required,min=6,max=10,eqfield=NewPassword"` //长度大于等于6,小于等于10,与新密码相同
}

type menuMeta struct {
	Title   string `json:"title"`   // 标题
	Icon    string `json:"icon"`    // 图标
	NoCache bool   `json:"noCache"` // 是不是缓存
}

type menuDataStruct struct {
	Path      string           `json:"path"`      // 路由
	Component string           `json:"component"` // 对应vue中的map name
	Name      string           `json:"name"`      // 菜单名称
	Hidden    bool             `json:"hidden"`    // 是否隐藏
	Meta      menuMeta         `json:"meta"`      // 菜单信息
	Children  []menuDataStruct `json:"children"`  // 子级菜单
}

type userData struct {
	Menus        []menuDataStruct `json:"menus"`        // 菜单
	Introduction string           `json:"introduction"` // 介绍
	Avatar       string           `json:"avatar"`       // 图标
	Name         string           `json:"name"`         // 姓名
	Email        string           `json:"email"`        // 邮箱
	Phone        string           `json:"phone"`        // 手机
	Memo         string           `json:"memo"`         // 备注
}

// @Tags User
// @Summary 用户登录
// @Produce  application/json
// @Param data body api.registAndLoginStuct true "用户登录接口"
// @Success 200 {string} string "{"code":20000,"message":"ok","data":{}}"
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	var ral registAndLoginStuct
	err := c.BindJSON(&ral)
	if err != nil {
		service.ResFail(c, err.Error())
		return
	}
	var hasUser bool
	user := new(sysmodel.SysUser)
	enPass := tools.MD5V([]byte(ral.Password))
	hasUser, err = db.DB.Where("username = ?", ral.Username).And("password = ?", enPass).Get(user)
	if err != nil {
		service.ResFail(c, err.Error())
		return
	}
	if !hasUser {
		service.ResFail(c, "用户名或密码错误")
		return
	}
	tokenNext(c, user)
}

func EditUserInfo(c *gin.Context) {
	var infoData editUserInfoStruct
	c.BindJSON(&infoData)
	user := new(sysmodel.SysUser)
	claims, _ := c.Get("claims")
	claimsData := claims.(*middleware.CustomClaims)
	user.NickName = infoData.NickName
	user.Email = infoData.Email
	user.Phone = infoData.Phone
	user.Memo = infoData.Memo
	affected, err := db.DB.Id(claimsData.Id).Update(user)
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

// @Tags User
// @Summary 修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body api.editPasswordStruct true "修改密码接口"
// @Success 200 {string} string "{"code":20000,"message":"修改成功","data":""}"
// @Router /user/editpwd [post]
func EditPassword(c *gin.Context) {
	var pwdData editPasswordStruct
	err := c.BindJSON(&pwdData)
	if err != nil {
		service.ResFail(c, "传值错误")
		return
	}
	var hasUser bool
	user := new(sysmodel.SysUser)
	oldPass := tools.MD5V([]byte(pwdData.OldPassword))
	newPass := tools.MD5V([]byte(pwdData.NewPassword))
	claims, _ := c.Get("claims")
	claimsData := claims.(*middleware.CustomClaims)
	hasUser, err = db.DB.Cols("password").Where("id = ?", claimsData.Id).And("password = ?", oldPass).Get(user)
	if !hasUser {
		service.ResFail(c, "原密码错误")
		return
	}
	user.Password = newPass
	affected, err := db.DB.Id(claimsData.Id).Update(user)
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

// @Tags User
// @Summary 获取用户菜单列表以及昵称
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"code":20000,"message":"ok","data":""}"
// @Router /user/info [get]
func UserInfo(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsData := claims.(*middleware.CustomClaims)
	menuModel := new(sysmodel.SysMenu)
	var err error
	var userMenus []sysmodel.SysMenu
	if claimsData.Id == config.SysConfig.System.SuperId {
		userMenus, err = menuModel.GetAllMenu()
	} else {
		//获取用户的角色
		syr := new(sysmodel.SysUserRole)
		var roleIds []uint
		roleIds, err = syr.GetUserRoleId(claimsData.Id)
		if err != nil {
			service.ResFail(c, err.Error())
			return
		}
		// 去除重复menu
		err = db.DB.Distinct("sys_menu.*").Table("sys_menu").Join("INNER", "sys_role_menu", "sys_menu.id=sys_role_menu.menu_id").Join("INNER", "sys_role", "sys_role.id=sys_role_menu.role_id").In("sys_role.id", roleIds).Asc("sys_menu.parent_id").Asc("sys_menu.sequence").Find(&userMenus)
	}
	if err != nil {
		service.ResFail(c, err.Error())
		return
	}
	var menus []menuDataStruct
	if len(userMenus) > 0 {
		var topmenuid uint = userMenus[0].ParentId
		if topmenuid == 0 {
			topmenuid = userMenus[0].Id
		}
		menus = setMenu(userMenus, topmenuid)
	}
	resData := userData{Menus: menus}
	user := new(sysmodel.SysUser)
	db.DB.Cols("nick_name, avatar, email, phone, memo").Where("id = ?", claimsData.Id).Get(user)
	resData.Name = user.NickName
	resData.Avatar = user.Avatar
	resData.Email = user.Email
	resData.Phone = user.Phone
	resData.Memo = user.Memo
	service.ResSuccess(c, "ok", resData)
	return
}

func tokenNext(c *gin.Context, user *sysmodel.SysUser) {
	j := &middleware.JWT{
		[]byte(config.SysConfig.JWT.SigningKey), // 唯一签名
	}
	clams := middleware.CustomClaims{
		Id:       user.Id,
		NickName: user.NickName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),       // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60*24*7), // 过期时间 一周
			Issuer:    config.SysConfig.JWT.SigningKey,       //签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		service.ResFail(c, "获取token失败")
		return
	}
	service.ResSuccess(c, "ok", gin.H{"token": token})
	return
	// if config.SysConfig.System.UseMultipoint {
	// 	var loginJwt model.JwtBlacklist
	// 	loginJwt.Jwt = token
	// 	err, jwtStr := loginJwt.GetRedisJWT(user.Username)
	// 	if err == redis.Nil {
	// 		err2 := loginJwt.SetRedisJWT(user.Username)
	// 		if err2 != nil {
	// 			c.JSON(http.StatusOK, gin.H{"code": 50008, "message": "设置登录状态失败"})
	// 			c.Abort()
	// 			return
	// 		} else {
	// 			c.JSON(http.StatusOK, gin.H{"code": 20000, "data": "登录成功"})
	// 			return
	// 		}
	// 	} else if err != nil {
	// 		c.JSON(http.StatusOK, gin.H{"code": 50008, "message": fmt.Sprintf("%v", err)})
	// 		c.Abort()
	// 		return
	// 	} else {
	// 		var blackjWT model.JwtBlacklist
	// 		blackjWT.Jwt = jwtStr
	// 		// err3 := blackjWT.JsonInBlacklist()
	// 		var err3 interface{}
	// 		if err3 != nil {
	// 			c.JSON(http.StatusOK, gin.H{"code": 50008, "message": "jwt作废失败"})
	// 			c.Abort()
	// 			return
	// 		} else {
	// 			err2 := loginJwt.SetRedisJWT(user.Username)
	// 			if err2 != nil {
	// 				c.JSON(http.StatusOK, gin.H{"code": 50008, "message": "设置登录状态失败"})
	// 				c.Abort()
	// 				return
	// 			} else {
	// 				c.JSON(http.StatusOK, gin.H{"code": 20000, "message": "登录成功"})
	// 				c.Abort()
	// 				return
	// 			}
	// 		}
	// 	}
	// } else {
	// c.JSON(http.StatusOK, gin.H{"code": 20000, "message": "ok", "data": gin.H{"token": token}})
	// return
	// }
}

// 递归菜单
func setMenu(menus []sysmodel.SysMenu, parentId uint) (out []menuDataStruct) {
	var menuArr []sysmodel.SysMenu
	linq.From(menus).Where(func(c interface{}) bool {
		return c.(sysmodel.SysMenu).ParentId == parentId
	}).OrderBy(func(c interface{}) interface{} {
		return c.(sysmodel.SysMenu).Sequence
	}).ToSlice(&menuArr)
	if len(menuArr) == 0 {
		return
	}
	noCache := false
	for _, item := range menuArr {
		menu := menuDataStruct{
			Path:      item.Uri,
			Component: item.Code,
			Name:      item.Name,
			Meta:      menuMeta{Title: item.Name, Icon: item.Icon, NoCache: noCache},
			Children:  []menuDataStruct{}}
		if item.MenuType == 3 {
			menu.Hidden = true
		}
		//查询是否有子级
		menuChildren := setMenu(menus, item.Id)
		if len(menuChildren) > 0 {
			menu.Children = menuChildren
		}
		if item.MenuType == 2 {
			// 添加子级首页，有这一级NoCache才有效
			menuIndex := menuDataStruct{
				Path:      "index",
				Component: item.Code,
				Name:      item.Name,
				Meta:      menuMeta{Title: item.Name, Icon: item.Icon, NoCache: noCache},
				Children:  []menuDataStruct{}}
			menu.Children = append(menu.Children, menuIndex)
			menu.Name = menu.Name + "index"
			menu.Meta = menuMeta{}
		}
		out = append(out, menu)
	}
	return
}

func Logout(c *gin.Context) {
	service.ResSuccess(c, "ok", "")
	return
}
