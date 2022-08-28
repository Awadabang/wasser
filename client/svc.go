package main

import (
	"log"
	"time"

	"github.com/Awadabang/wasser/types/worker"
)

type Client struct {
	LifeCircle worker.Worker_LifeCircleClient
	StatusSync worker.Worker_StatusSyncClient
}

func (c *Client) Run() {
	ticker := time.NewTicker(3 * time.Second)
	for ; true; <-ticker.C {
		if err := c.Register(); err != nil {
			log.Println(err)
			continue
		}
		log.Println("Register OK")
		break
	}

	c.AliveCheck()
}
