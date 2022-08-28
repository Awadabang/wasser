package main

import (
	"errors"

	"github.com/Awadabang/wasser/types/worker"
)

func (c *Client) Register() error {
	err := c.LifeCircle.Send(&worker.LifeCircleReq{
		Register:   true,
		Alivecheck: false,
	})
	if err != nil {
		return err
	}

	res, err := c.LifeCircle.Recv()
	if err != nil {
		return err
	}
	if res.Status == "OK" {
		return nil
	} else {
		return errors.New("Register Not OK")
	}
}
