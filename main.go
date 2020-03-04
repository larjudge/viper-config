package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io/ioutil"
)

var (
	configFile = flag.StringP("config", "c", "", "specific config file to use")

	boolConfig   = flag.BoolP("boolConfig", "b", false, "some bool")
	intConfig    = flag.IntP("intConfig", "i", 0, "some int")
	stringConfig = flag.StringP("stringConfig", "s", "DEFAULT", "some string")

	flagConfig = &Config{
		AppSpec: AppSpec{
			BoolConfig:   *boolConfig,
			IntConfig:    *intConfig,
			StringConfig: *stringConfig,
		},
	}
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
	fmt.Printf("\n\nDefault CONFIG IS %+v\n\n", defaultConfig)
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




	fmt.Printf("\nkeys : %+v", viper.AllKeys())

	return &config, nil
}

func (c Config) GetMap(config interface{}) map[string]interface{} {

	var inInterface map[string]interface{}
	mapstructure.Decode(config, &inInterface)

	fmt.Printf("inInterface is %+v", inInterface)

	return inInterface
}

func main() {
	flag.Parse()

	x, _ := ioutil.ReadFile(*configFile)
	fmt.Println(string(x))


	fmt.Printf("\n\nFlagConfig IS %+v\n\n", flagConfig)

	config, _ := loadConfig(*configFile, *flag.CommandLine)
	fmt.Printf("\n\nFLAGS are \n\t%+v\n\t%+v\n\t%+v\n", *stringConfig, *intConfig, *boolConfig)
	fmt.Printf("\n\nCONFIG IS %+v\n\n", config)
}
