package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()

	logger := log.With(
		log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	errChan := make(chan error)
	go func() {
		errChan <- newService(context.Background())
	}()

	select {
	case err := <-errChan:
		if err != nil {
			panic(err)
		}
	case <-sigChan:
		logger.Log(log.LevelFatal, "Bye!", "Bro!")
		return
	}
}

func newService(ctx context.Context) error {
	time.Sleep(10 * time.Second)
	return nil
}
