package global

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var Logger *slog.Logger

var DB *sqlx.DB

var Engine *gin.Engine

var Configs Config

type Config struct {
	Server ServerConfig `mapstructure:"server" yaml:"server"`
	MySQL  MySQLConfig  `mapstructure:"mysql"  yaml:"mysql"`
}

type ServerConfig struct {
	APIHost  string `mapstructure:"apiHost" yaml:"apiHost"`
	APIPort  int    `mapstructure:"apiPort" yaml:"apiPort"`
	UserHost string `mapstructure:"userHost" yaml:"userHost"`
	UserPort int    `mapstructure:"userPort" yaml:"userPort"`
}

type MySQLConfig struct {
	User     string `mapstructure:"user"     yaml:"user"`
	Password string `mapstructure:"password" yaml:"password"`
	Host     string `mapstructure:"host"     yaml:"host"`
	Port     int    `mapstructure:"port"     yaml:"port"`
	Name     string `mapstructure:"name"     yaml:"name"`
}
