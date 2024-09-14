package configs

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Client struct {
	Address      string        `yaml:"address"`
	Timeout      time.Duration `yaml:"timeout"`
	RetriesCount int           `yaml:"retries_count"`
}

type ClientConfig struct {
	Auth      Client `yaml:"auth"`
	AppSecret string `yaml:"app_secret" env-required:"true" env:"APP_SECRET"`
}

func MustLoad() *ClientConfig {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config file is empty")
	}

	return MustLoadByPath(configPath)
}

func MustLoadByPath(configPath string) *ClientConfig {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config path does not exist:" + configPath)
	}

	var config ClientConfig

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		panic("config file is empty:" + configPath)
	}

	return &config
}
func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "config.yaml", "config file path")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
