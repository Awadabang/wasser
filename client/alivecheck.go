package main

import (
	"log"
	"time"

	"github.com/Awadabang/wasser/types/worker"
)

func (c *Client) AliveCheck() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for ; true; <-ticker.C {
		err := c.LifeCircle.Send(&worker.LifeCircleReq{
			Alivecheck: true,
		})
		if err != nil {
			return
		}

		res, err := c.LifeCircle.Recv()
		if err != nil {
			return
		}

		log.Println("Worker Alive: ", res.Alive)

		if !res.Alive {
			for ; true; <-ticker.C {
				if err := c.Register(); err != nil {
					log.Println("重连失败", err)
				} else {
					log.Println("重连成功")
					break
				}
			}
		}
	}
}
