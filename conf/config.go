package conf

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//Config represents server configuration
type Config struct {
	Path string
}

//Init server config
func Init(cfg string) error {
	c := Config{Path: cfg}
	if err := c.initConfig(); err != nil {
		return err
	}
	c.watch()
	return nil
}

func (c *Config) initConfig() error {
	if c.Path != "" {
		viper.SetConfigFile(c.Path)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func (c Config) watch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Config file changed: %s", in.Name)
	})
}
