package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

const (
	boolValue   = false
	stringValue = "some string"
	intValue    = 1234
)

var (
	boolConfig   bool
	intConfig    int
	stringConfig string
)

type AppConfig struct {
	BoolConfig   bool   `yaml:"boolConfig"`
	IntConfig    int    `yaml:"intConfig"`
	StringConfig string `yaml:"stringConfig"`
}

func LoadConfig() (*AppConfig, error) {
	flag.IntVarP(&intConfig, "intConfig", "i", intValue, "some int")
	flag.StringVarP(&stringConfig, "stringConfig", "s", stringValue, "some string")
	flag.BoolVarP(&boolConfig, "boolConfig", "b", boolValue, "some bool")

	flag.Parse()

	err := viper.BindPFlags(flag.CommandLine)
	if err != nil {
		log.Fatal("cannot bind command line flags")
	}

	config := &AppConfig{
		BoolConfig:   boolConfig,
		IntConfig:    intConfig,
		StringConfig: stringConfig,
	}

	return config, nil
}

func main () {
	config, _ := LoadConfig()
	fmt.Printf("\n\nCONFIG IS %+v\n\n", config)
}