package main

import (
	"flag"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"io/ioutil"
)

var (
	boolConfig   bool
	intConfig    int
	stringConfig string

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

	// makes config discoverable at /config/config.yaml

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

//
//func loadConfig(srcFile string) (*Config, error) {
//	if srcFile != "" {
//		fmt.Printf("\n Source file is %s\n", srcFile)
//		viper.SetConfigFile(srcFile)
//	}
//
//	// makes config discoverable at /config/config.yaml
//	viper.SetConfigType("yaml")
//	viper.SetConfigName("local")
//	viper.AddConfigPath("/config")
//	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
//	viper.AutomaticEnv()
//
//	var config Config
//
//	// convert from struct to generic map (required for viper to merge correctly) and set the defaults (will be used if not explicitly set via environment or config file)
//	viper.SetDefault("spec", config.GetMap(defaultConfig))
//
//	err := viper.ReadInConfig()
//	if err != nil {
//		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
//			fmt.Println("LAR LAR LAR")
//			fmt.Printf("\nviper.ConfigFileUsed(): \n %+v \n", viper.ConfigFileUsed())
//			if err := viper.Unmarshal(&config); err != nil {
//				fmt.Println(fmt.Errorf("Fatal unmarshal config file: %s \n", err))
//				return nil, err
//			}
//		}
//	}
//	if err := viper.Unmarshal(&config); err != nil {
//		return nil, err
//	}
//
//	//flag.IntVarP(&intConfig, "intConfig", "i", defaultConfig.IntConfig, "some int")
//	//flag.StringVarP(&stringConfig, "stringConfig", "s", defaultConfig.StringConfig, "some string")
//	//flag.BoolVarP(&boolConfig, "boolConfig", "b", defaultConfig.BoolConfig, "some bool")
//	//
//	////flag.Parse()
//	//
//	//err = viper.BindPFlags(flag.CommandLine)
//	//if err != nil {
//	//	log.Fatal("cannot bind command line flags")
//	//}
//	//
//	//config = Config{
//	//	BoolConfig:   boolConfig,
//	//	IntConfig:    intConfig,
//	//	StringConfig: stringConfig,
//	//}
//
//	return &config, nil
//}

func (c Config) GetMap(config interface{}) map[string]interface{} {

	var inInterface map[string]interface{}
	mapstructure.Decode(config, &inInterface)

	fmt.Printf("inInterface is %+v", inInterface)

	return inInterface
}

func main() {

	//flag.StringVarP(&configFile, "config", "c", "", "specify config path")

	//configFile := flag.String("config", "", "specific config file to use")
	flag.Parse()

	x, err := ioutil.ReadFile(*configFile)
	fmt.Println(string(x))
	fmt.Println(err)

	config, err := loadConfig(*configFile)
	fmt.Printf("err is %+v", err)
	fmt.Printf("\n\nCONFIG IS %+v\n\n", config)
}
