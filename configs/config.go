package configs

import (
	"flag"
	"github.com/spf13/viper"
)

type config struct {
	paths *struct {
		templatePath string
		outPath      string
	}
	server *struct {
		httpPort string
	}
}

func init() {
	var configPath, configFile string

	flag.StringVar(&configPath, "path", "configs", "Path to config file")
	flag.StringVar(&configFile, "config", "config", "Name of config file")
	flag.StringVar(&configPath, "p", "configs", "Path to config file")
	flag.StringVar(&configFile, "c", "config", "Name of config file")
	flag.Parse()

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configFile)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func NewConfig() *config {
	return &config{
		paths: &struct {
			templatePath string
			outPath      string
		}{
			templatePath: viper.GetString("paths.templatePath"),
			outPath:      viper.GetString("paths.outPath"),
		},
		server: &struct {
			httpPort string
		}{
			httpPort: viper.GetString("server.httpPort"),
		},
	}
}
