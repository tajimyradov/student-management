package config

import (
	"github.com/spf13/viper"
	"time"
)

type AppConfig struct {
	HTTP      HTTP     `mapstructure:"http"`
	StudentDB Postgres `mapstructure:"student_postgres"`
	Secrets   Secrets  `mapstructure:"secrets"`
	Domains   Domains  `mapstructure:"domains"`
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

type Secrets struct {
	AccessSecret string `mapstructure:"access_secret"`
	PasswordSalt string `mapstructure:"password_salt"`
}

type Domains struct {
	Image string `mapstructure:"image"`
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
