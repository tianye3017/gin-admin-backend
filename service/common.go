package service

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS_CODE = 20000 //成功的状态码
	FAIL_CODE    = 60000 //失败的状态码
)

// ResSuccess 返回成功
func ResSuccess(c *gin.Context, m string, d interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": SUCCESS_CODE, "message": m, "data": d})
	c.Abort()
}

// ResFail 返回失败
func ResFail(c *gin.Context, m string) {
	c.JSON(http.StatusOK, gin.H{"code": FAIL_CODE, "message": m})
	c.Abort()
}

// GetStrBefore 获取字符串某一字符前的字符串
func GetStrBefore(str string, position string) (c string) {
	if strings.Contains(str, position) {
		return str[:strings.Index(str, position)]
	}
	return str
}

// GetStrAfter 获取字符串某一字符后的字符串
func GetStrAfter(str string, position string) (c string) {
	if strings.Contains(str, position) {
		return str[strings.Index(str, position)+1:]
	}
	return str
}

// ResolveSorKey 解析排序关键字
func ResolveSorKey(str string) (e error, s string) {
	if !strings.HasPrefix(str, "+") && !strings.HasPrefix(str, "-") {
		return errors.New("排序关键字错误"), s
	}
	if strings.HasPrefix(str, "+") {
		s = GetStrAfter(str, "+") + " ASC"
	} else {
		s = GetStrAfter(str, "-") + " DESC"
	}
	return
}
