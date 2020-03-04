package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	configFile = flag.String("config", "", "specific config file to use")
)

type Config struct {
	AppSpec AppSpec `json:"spec"           yaml:"spec"          mapstructure:"spec"`
}

type AppSpec struct {
	BoolConfig   bool   `json:"boolConfig" yaml:"boolConfig" mapstructure:"boolConfig"`
	IntConfig    int    `json:"intConfig" yaml:"intConfig" mapstructure:"intConfig"`
	StringConfig string `json:"stringConfig" yaml:"stringConfig" mapstructure:"stringConfig"`
}

var (
	defaultConfig = &Config{
		AppSpec: AppSpec{
			BoolConfig:   false,
			IntConfig:    0,
			StringConfig: "",
		},
	}
)

func loadConfig(srcFile string) (*Config, error) {
	if srcFile != "" {
		viper.SetConfigFile(srcFile)
	} else {
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
		viper.AddConfigPath("/config")
		viper.AddConfigPath("/app/config")
	}

	var config Config

	// convert from struct to generic map (required for viper to merge correctly) and set the defaults (will be used if not explicitly set via environment or config file)
	viper.SetDefault("spec", config.GetMap(defaultConfig.AppSpec))

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println(viper.ConfigFileUsed())
			if err := viper.Unmarshal(&config); err != nil {
				fmt.Println(fmt.Errorf("Fatal unmarshal config file: %s \n", err))
				return nil, err
			}
			return &config, nil
		} else {
			// if no config file found use the default configuration
			fmt.Println("error:")
			fmt.Println(err)
			return defaultConfig, err
		}
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (c Config) GetMap(config interface{}) map[string]interface{} {
	var inInterface map[string]interface{}
	mapstructure.Decode(config, &inInterface)
	return inInterface
}

func main() {
	flag.Parse()

	config, _ := loadConfig(*configFile)

	fmt.Printf("\n\nCONFIG IS %+v\n\n", config)
}
