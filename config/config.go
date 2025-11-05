package config

import "github.com/spf13/viper"

type App struct {
	AppPort string `json:"app_port"`
	AppEnv  string `json:"app_env"`

	JwtSecretKey string `json:"jwt_secret_key"`
	JwtIssuer    string `json:"jwt_issuer"`
}

type PsqlDB struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	DBName    string `json:"db_name"`
	DBMaxOpen int    `json:"db_max_open"`
	DBMaxIdle int    `json:"db_max_idle"`
}

type Config struct {
	App    App
	PsqlDB PsqlDB
}

func NewConfig() *Config {
	return &Config{
		App: App{
			AppPort:      viper.GetString("app.port"),
			AppEnv:       viper.GetString("app.env"),
			JwtSecretKey: viper.GetString("jwt.secret_key"),
			JwtIssuer:    viper.GetString("jwt.issuer"),
		},
		PsqlDB: PsqlDB{
			Host:      viper.GetString("database.host"),
			Port:      viper.GetInt("database.port"),
			User:      viper.GetString("database.user"),
			Password:  viper.GetString("database.password"),
			DBName:    viper.GetString("database.name"),
			DBMaxOpen: viper.GetInt("database.max_open"),
			DBMaxIdle: viper.GetInt("database.max_idle"),
		},
	}
}
