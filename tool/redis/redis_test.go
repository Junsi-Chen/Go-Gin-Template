package redis

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	ins := &Instance{
		DBNumber:  0,
		Addresses: []string{"127.0.0.1:6379"},
		Password:  "",
		MaxIdle:   30,
		MaxActive: 30,
	}
	err := InitClient(ins)
	if err != nil {
		t.Fatal(err)
	}

	err = GetRdb().Set(context.Background(), "test1", "hello world", time.Minute).Err()
	if err != nil {
		t.Fatal(err)
	}
	value, err := GetRdb().Get(context.Background(), "test1").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(value)
}
