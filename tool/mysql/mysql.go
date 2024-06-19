package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Instance struct {
	User         string `toml:"user"`
	Password     string `toml:"password"`
	Host         string `toml:"host"`
	DatabaseName string `toml:"databaseName"`
	Charset      string `toml:"charset"`
	LogShow      bool   `toml:"logShow"`
}

// InitEngine 初始化数据库实例
func InitEngine(ins *Instance) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		ins.User, ins.Password, ins.Host, ins.DatabaseName, ins.Charset)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DefaultStringSize:        255, // string 类型字段的默认长度
		DisableDatetimePrecision: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项 不会在尾部加"s"
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// SetMaxIdleConns设置空闲状态下的最大连接数
	sqlDB.SetMaxIdleConns(5)
	// SetMaxOpenConns设置数据库打开的最大连接数
	sqlDB.SetMaxOpenConns(10)
	// ping
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}
	return db, err
}

// AutoMigrate 表同步 - 每次启动项目都会同步数据库的表，
func AutoMigrate(db *gorm.DB, tables ...interface{}) (err error) {
	// 表同步
	err = db.Debug().AutoMigrate(tables...)
	if err != nil {
		return err
	}
	return
}
