package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofrs/uuid"
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

	//邮件主题为"Hello"
	subject := "Hello"
	// 邮件正文
	body := "bad"
	err := helpers.SendEMail(subject, body, to)
	if err != nil {
		fmt.Println(err)
	}
}

func TestUuid(t *testing.T) {
	// Creating UUID Version 4
	// panic on error
	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)

	// or error handling
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("UUIDv4: %s\n", u2)

	// Parsing UUID from string input
	u2, err = uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)
}
