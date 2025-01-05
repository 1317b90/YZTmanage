package Data

import (
	"YZT/Config"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var Rdb = redis.NewClient(&redis.Options{
	Addr:     Config.Redis_Addr, // Redis 地址
	Password: "",                // 密码
	DB:       0,                 // 默认数据库
})

var CTX = context.Background()

func Redis_test() {
	_, err := Rdb.Ping(CTX).Result()
	if err != nil {
		fmt.Println("Redis连接失败:", err)
	} else {
		fmt.Println("Redis连接成功")
	}
}
