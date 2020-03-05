package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io/ioutil"
)

const (
	defaultHealthCheckURL = "/healthz"
)
var (
	configFile = flag.StringP("config", "c", "", "specific config file to use")
	boolConfig   = flag.BoolP("spec.boolConfig", "b", false, "some bool")
	intConfig    = flag.IntP("spec.intConfig", "i", 0, "some int")
	stringConfig = flag.StringP("spec.stringConfig", "s", "DEFAULT", "some string")
	healthCheckURL = flag.StringP("opSpec.healthCheckURL", "e", defaultHealthCheckURL, "some string")

	defaultConfig = &Config{
		AppSpec: AppSpec{
			BoolConfig:   *boolConfig,
			IntConfig:    *intConfig,
			StringConfig: *stringConfig,
		},
		OpSpec: OpSpec{HealthCheckURL: *healthCheckURL},
	}
)

type OpSpec struct{
	HealthCheckURL string `json:"healthCheckURL" yaml:"healthCheckURL" mapstructure:"healthCheckURL"`
}

type Config struct {
	AppSpec AppSpec `json:"spec" yaml:"spec" mapstructure:"spec"`
	OpSpec  OpSpec  `json:"opSpech" yaml:"opSpec mapstructure:"opSpec"`
}

type AppSpec struct {
	BoolConfig   bool   `json:"boolConfig" yaml:"boolConfig" mapstructure:"boolConfig"`
	IntConfig    int    `json:"intConfig" yaml:"intConfig" mapstructure:"intConfig"`
	StringConfig string `json:"stringConfig" yaml:"stringConfig" mapstructure:"stringConfig"`
}

func loadConfig(srcFile string, fl flag.FlagSet) (*Config, error) {
	viper.BindPFlags(&fl)

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
	viper.SetDefault("spec",  config.GetMap(defaultConfig.AppSpec))
	viper.SetDefault("opSpec",  config.GetMap(defaultConfig.OpSpec))

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
	config, _ := loadConfig(*configFile, *flag.CommandLine)

	yaml, _ := ioutil.ReadFile(*configFile)
	fmt.Printf("YAML \n%+v\n\n", string(yaml))
	fmt.Printf("DEFAULT CONFIG IS \n\t%+v\n\n", defaultConfig)
	fmt.Printf("\n\nFLAGS are \n\t%+v\n\t%+v\n\t%+v\n\t%+v\n", *stringConfig, *intConfig, *boolConfig, *healthCheckURL)
	fmt.Printf("\n\nCONFIG IS \n\t%+v\n\n", config)
}