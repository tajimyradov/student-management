package config

import (
	"github.com/spf13/viper"
	"time"
)

type AppConfig struct {
	HTTP      HTTP     `mapstructure:"http"`
	StudentDB Postgres `mapstructure:"student_postgres"`
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  bool   `mapstructure:"sslmode"`
}

type HTTP struct {
	Host               string        `mapstructure:"host"`
	Port               string        `mapstructure:"port"`
	ReadTimeout        time.Duration `mapstructure:"readTimeout"`
	WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
	MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
}

func NewAppConfig(configFile string) (*AppConfig, error) {
	config, err := loadConfig(configFile)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func loadConfig(configFile string) (*AppConfig, error) {
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var appConfig AppConfig
	err = viper.Unmarshal(&appConfig)
	if err != nil {
		return nil, err
	}

	return &appConfig, err
}
