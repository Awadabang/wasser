package main

import (
	"context"
	"flag"
	"fmt"

	cache "github.com/Awadabang/wasser/pkg/redis"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	Name     string
	Versiont string
	flagConf string
)

func init() {
	flag.StringVar(&flagConf, "conf", "configs/config.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	// config
	flag.Parse()
	viper.SetConfigFile(flagConf)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redisAddr"),
		Password: viper.GetString("redisPassword"),
		DB:       0,
	})

	pong, err := client.Ping(client.Context()).Result()
	fmt.Println(pong)

	var cacheMgr cache.Manager
	cacheMgr = cache.NewStreamMQ()

	cacheMgr.SendMsg(context.Background(), &cache.Msg{
		Topic: "Alert",
		Body:  []byte("test"),
	})
}
