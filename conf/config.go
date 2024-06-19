package conf

import (
	"github.com/spf13/viper"
	"template/tool/http"
	"template/tool/mysql"
)

var Conf *Config

type App struct {
	ServerName string `toml:"serverName"`
	Debug      bool   `toml:"debug"`
	JwtSecret  string `toml:"jwtSecret"`
}

type SignGrpc struct {
	Host string `toml:"host"`
}

type Config struct {
	App      *App
	Mysql    *mysql.Instance
	Http     *http.Instance
	SignGrpc *SignGrpc
}

func Init() error {
	v := viper.New()
	v.AddConfigPath("conf")
	v.SetConfigType("toml")
	err := v.ReadInConfig()
	if err != nil {
		return err
	}
	err = v.Unmarshal(&Conf)
	return err
}
