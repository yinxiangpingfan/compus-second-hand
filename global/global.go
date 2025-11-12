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
	Server ServerConfig `yaml:"server"`
	MySQL  MySQLConfig  `yaml:"mysql"`
	JWT    JWTConfig    `yaml:"jwt"`
}

type ServerConfig struct {
	APIHost  string `yaml:"apiHost"`
	APIPort  int    `yaml:"apiPort"`
	UserHost string `yaml:"userHost"`
	UserPort int    `yaml:"userPort"`
}

type MySQLConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
}

type JWTConfig struct {
	SecretKey  string `yaml:"secretKey"`
	ExpireTime int64  `yaml:"expireTime"`
}
