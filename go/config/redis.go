package config

import (
	"github.com/go-redis/redis"
	"tiktok/go/util"
)

// 声明一个全局的redisDb变量
//var redisDb *redis.Client

// InitRedisClient 根据redis配置初始化一个客户端
func InitRedisClient() (redisDb *redis.Client, err error) {
	// 替换你的账号密码
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err = redisDb.Ping().Result()
	if err != nil {
		util.LogFatal(err.Error())
		panic(err)
		return nil, err
	}
	util.Log("redis 初始化成功")
	return redisDb, nil
}

// RedisClient 选择redis的数据库 -> i
func RedisClient(i int) (*redis.Client, error) {
	redisDb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       i,                // 默认数据库，默认是0
	})
	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err := redisDb.Ping().Result()
	if err != nil {
		util.LogFatal(err.Error())
		panic(err)
		return nil, err
	}
	return redisDb, nil
}
