package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"restful-api-kit/config"
	"restful-api-kit/helpers"
	"testing"
)

var ctx = context.Background()

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123", // no password set
		DB:       0,     // use default DB
	})

	//err := rdb.Set(ctx, "key", "value", 0).Err()
	//if err != nil {
	//	panic(err)
	//}

	val, err := rdb.Get(ctx, "a").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func TestEmail(t *testing.T) {
	to := "mihanglaoban@outlook.com"
	config.Setup("../config/dbStaging.json")
	helpers.InitMail()
	//定义收件人
	mailTo := []string{
		to,
	}
	//邮件主题为"Hello"
	subject := "Hello"
	// 邮件正文
	body := "Good"
	err := helpers.SendMail(mailTo, subject, body)
	if err != nil {
		fmt.Println(err)
	}
}
