package config

type Config struct {
	JWT      JWT      `json:"jwt"`
	Database Database `json:"database"`
	Redis    Redis    `json:"redis"`
	System   System   `json:"system"`
}

type System struct { // 系统配置
	UseMultipoint bool   `json:"useMultipoint"`
	Env           string `json:"env"`
	Addr          int    `json:"addr"`
	SuperId       uint   `json:"super_id"` // 超级管理员id
}

type JWT struct { // jwt签名
	SigningKey string `json:"signingKey"`
}

type Database struct { // 数据库配置
	Username string `json:"username"`
	Password string `json:"password"`
	Path     string `json:"path"`
	Dbname   string `json:"dbname"`
	Config   string `json:"config"`
}

type Redis struct { // Redis admin 数据库配置
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

var SysConfig = Config{
	JWT: JWT{
		SigningKey: "Lothar",
	},
	Database: Database{
		Username: "root",
		Password: "admin",
		Path:     "127.0.0.1:3306",
		Dbname:   "vue-ele-admin",
		Config:   "charset=utf8mb4",
	},
	Redis: Redis{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	},
	System: System{
		UseMultipoint: false,
		Env:           "develop",
		Addr:          8888,
		SuperId:       1,
	},
}
