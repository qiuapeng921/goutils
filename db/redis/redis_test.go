package redis

import (
	"fmt"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	var config Config
	config = Config{
		Address:     "127.0.0.1:6379",
		Password:    "",
		DB:          5,
		DialTimeout: time.Duration(10),
		CmdTimeout:  time.Duration(10),
		PoolSize:    10,
		MinIdle:     10,
	}

	if err := Init(config); err != nil {
		fmt.Println(err.Error())
	}
}

func TestClient(t *testing.T) {
	Client().HSet("test", "test","test")
}