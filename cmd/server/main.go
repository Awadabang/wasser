package main

import (
	"flag"
	"fmt"
	"log"

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
	log.Println(viper.Get("name"))
}
