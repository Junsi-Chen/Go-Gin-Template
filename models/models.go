package models

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"sync"
	"template/tool/log"
	"template/tool/mysql"
	"time"
)

var (
	db   *gorm.DB
	once sync.Once
)

// InitDb 初始化数据库
func InitDb(dbEngine *gorm.DB) {
	once.Do(func() {
		db = dbEngine
	})
}

// GetDb 获取数据库实例 -> 为了单例
func GetDb() *gorm.DB {
	return db
}

// Model 基类
type Model struct {
	Id         int64 `gorm:"primaryKey"`
	CreateTime time.Time
	UpdateTime time.Time
	IsDelete   soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
}

// AutoMigrate 表同步 - 每次启动项目都会同步数据库的表，
func AutoMigrate() {
	tables := []interface{}{
		//new(user.User),
	}
	err := mysql.AutoMigrate(db, tables...)
	if err != nil {
		log.Logger.Error("mysql AutoMigrate err:" + err.Error())
		panic(err)
	}
}
