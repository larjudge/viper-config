package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var (
	boolConfig   bool
	intConfig    int
	stringConfig string
	configFile = flag.String("config", "", "specific config file to use")
)

type Config struct {
	BoolConfig   bool   `yaml:"boolConfig"`
	IntConfig    int    `yaml:"intConfig"`
	StringConfig string `yaml:"stringConfig"`
}

var (
	defaultConfig = &Config{
		BoolConfig:   false,
		IntConfig:    0,
		StringConfig: "",
	}
)

func loadConfig(srcFile string) (*Config, error) {
	if srcFile != "" {
		viper.SetConfigFile(srcFile)
	}

	// makes config discoverable at /config/config.yaml
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("/config")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var config Config

	// convert from struct to generic map (required for viper to merge correctly) and set the defaults (will be used if not explicitly set via environment or config file)
	viper.SetDefault("spec", config.GetMap(defaultConfig))

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("LAR LAR LAR")
			fmt.Println(viper.ConfigFileUsed())
			if err := viper.Unmarshal(&config); err != nil {
				fmt.Println(fmt.Errorf("Fatal unmarshal config file: %s \n", err))
				return nil, err
			}
		}
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	flag.IntVarP(&intConfig, "intConfig", "i", defaultConfig.IntConfig, "some int")
	flag.StringVarP(&stringConfig, "stringConfig", "s", defaultConfig.StringConfig, "some string")
	flag.BoolVarP(&boolConfig, "boolConfig", "b", defaultConfig.BoolConfig, "some bool")

	flag.Parse()

	err = viper.BindPFlags(flag.CommandLine)
	if err != nil {
		log.Fatal("cannot bind command line flags")
	}

	config = Config{
		BoolConfig:   boolConfig,
		IntConfig:    intConfig,
		StringConfig: stringConfig,
	}

	return &config, nil
}

func (c Config) GetMap(config interface{}) map[string]interface{} {

	var inInterface map[string]interface{}
	mapstructure.Decode(config, &inInterface)
	return inInterface
}

func main() {
	config, _ := loadConfig("local.yaml")
	fmt.Printf("\n\nCONFIG IS %+v\n\n", config)
}
