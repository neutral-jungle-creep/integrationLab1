package configs

import (
	"flag"
	"github.com/spf13/viper"
)

type config struct {
	Paths *struct {
		TemplatePath string
		OutPath      string
	}
	Server *struct {
		HttpPort string
	}
}

func init() {
	var configPath, configFile string

	flag.StringVar(&configPath, "path", "configs", "Path to config file")
	flag.StringVar(&configFile, "config", "configTemplate", "Name of config file")
	flag.StringVar(&configPath, "p", "configs", "Path to config file")
	flag.StringVar(&configFile, "c", "configTemplate", "Name of config file")
	flag.Parse()

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configFile)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func NewConfig() *config {
	return &config{
		Paths: &struct {
			TemplatePath string
			OutPath      string
		}{
			TemplatePath: viper.GetString("paths.templatePath"),
			OutPath:      viper.GetString("paths.outPath"),
		},
		Server: &struct {
			HttpPort string
		}{
			HttpPort: viper.GetString("server.httpPort"),
		},
	}
}
