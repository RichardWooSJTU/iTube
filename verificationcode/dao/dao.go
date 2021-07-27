package dao

import "github.com/go-redis/redis/v8"

var DefaultRedis *redis.Client

func Init() {
	DefaultRedis = redis.NewClient(&redis.Options{
		Addr:     "1.15.72.208:6379",
		Password: "liqqjw213u3o8joiuehi23e983h98HH8H*(8hoh8", // no password set
		DB:       0,  // use default DB
	})
}
