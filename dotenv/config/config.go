package config

import "time"

// 定义配置常量，避免去拼写，防止错误
const (
	DEV  = "development"
	PROD = "production"
)

// Appf 提供一个全局环境变量，前提需要设置一下
var Appf Config

// Config 环境变量结构体
type Config struct {
	// http server config
	Port         int
	Mode         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	// database config
	DbType      string `env:"DBType"`
	DbHost      string `env:"DBHost"`
	DbPort      int    `env:"DBPort"`
	DbUser      string `env:"DBUser"`
	DbPswd      string `env:"DBPswd"`
	MaxIdleConn int
	MaxOpenConn int

	// token config
	JwtSecretKey        string
	AccessTokenExpires  time.Duration
	RefreshTokenExpires time.Duration

	// image upload config
	UploadImageAllowExts []string

	Nested struct {
		Foo string
		Bar time.Duration
		Qux bool
	}
}
