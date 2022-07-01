package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/dev-heeyoung/bookstore/pkg/config"
	"github.com/dev-heeyoung/bookstore/pkg/models"
	"github.com/go-redis/redis/v9"
)

var (
	RedisClient = config.ConnectRedis()
	Ctx         = context.Background()
)

func GetBook(id string) *models.Book {
	pong, err := RedisClient.Ping(Ctx).Result()
	fmt.Println(pong, err)

	var book models.Book
	value, err := RedisClient.Get(Ctx, id).Result()

	if err == redis.Nil {
		return &book
	} else if err != nil {
		panic(err)
	} else if value != "" {
		err = json.Unmarshal([]byte(value), &book)
		if err != nil {
			panic(err)
		}
	}

	return &book
}

func SaveBook(book *models.Book) {
	bookBytes, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}

	err = RedisClient.Set(Ctx, strconv.FormatUint(uint64(book.ID), 10), bookBytes, 0).Err()

	if err != nil {
		panic(err)
	}
}

func DeleteBook(id string) {
	_, err := RedisClient.Get(Ctx, id).Result()

	if err == redis.Nil {
		return
	} else if err != nil {
		panic(err)
	}
	RedisClient.Del(Ctx, id)
}
