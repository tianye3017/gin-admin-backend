package model

import (
	"log"
	"time"

	"github.com/tianye3017/gin-admin-backend/db"
)

type JwtBlacklist struct {
	Id         int       `xorm:"pk autoincr"`
	Jwt        string    `xorm:"TEXT"`
	Updated_at time.Time `xorm:"updated"`
	Created_at time.Time `xorm:"created"`
	Deleted_at time.Time `xorm:"deleted"`
}

var JwtBlacklistModel *JwtBlacklist

func init() {
	JwtBlacklistModel = new(JwtBlacklist)
	if err := db.DB.Sync(JwtBlacklistModel); err != nil {
		log.Fatal("数据表同步失败:", err)
	}
}

// func (j *JwtBlacklist) JsonInBlacklist() (err error) {
// 	err = qmsql.DEFAULTDB.Create(j).Error
// 	return
// }

//判断JWT是否在黑名单内部
func (j *JwtBlacklist) IsBlacklist(Jwt string) bool {
	has, _ := db.DB.Where("jwt = ?", Jwt).Get(j)
	return has
}

//判断当前用户是否在线
func (j *JwtBlacklist) GetRedisJWT(userName string) (err error, RedisJWT string) {
	RedisJWT, err = db.Redis.Get(userName).Result()
	return err, RedisJWT
}

//设置当前用户在线
func (j *JwtBlacklist) SetRedisJWT(userName string) (err error) {
	err = db.Redis.Set(userName, j.Jwt, 1000*1000*1000*60*60*24*7).Err()
	return err
}
