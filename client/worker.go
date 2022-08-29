package main

import (
	"context"
	"log"
	"time"

	"github.com/Awadabang/wasser/workerclient"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/zrpc"
)

func main() {
	cli := zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{"127.0.0.1:8080"},
	})
	defer cli.Conn().Close()

	workerClient := workerclient.NewWorker(cli)

	workerId := uuid.NewString()

	ticker := time.NewTicker(5 * time.Second)
	for ; true; <-ticker.C {
		lifeCircle, err := workerClient.LifeCircle(context.Background())
		if err != nil {
			log.Println(err)
			continue
		}
		statusSync, err := workerClient.StatusSync(context.Background())
		if err != nil {
			log.Println(err)
			continue
		}

		client := &Client{
			LifeCircle: lifeCircle,
			StatusSync: statusSync,

			WorkerId: workerId,
		}

		// 注册、健康检查
		client.Run()

		// monitor掉线
		log.Println("Ymonitor 掉线")
	}
}
