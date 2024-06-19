package mysql

import (
	"testing"
)

// TestMysql 测试数据库连接
func TestMysql(t *testing.T) {
	ins := &Instance{
		User:         "root",
		Password:     "root",
		Host:         "127.0.0.1:3306",
		DatabaseName: "",
		Charset:      "utf8mb4",
		LogShow:      false,
	}
	db, err := InitEngine(ins)
	if err != nil {
		t.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	err = sqlDB.Ping()
	if err != nil {
		t.Fatal(err)
	}
}
