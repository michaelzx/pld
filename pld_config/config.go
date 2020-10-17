package pld_config

import (
	"github.com/BurntSushi/toml"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"time"
)

type MixinConfig interface {
	// 加载配置，用什么实现可自行定义。下面提供了加载toml的加载方式
	Load(configFilePath string)
}

func LoadToml(path string, cfg MixinConfig) {
	_, err := toml.DecodeFile(path, cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
}
func LoadYaml(path string, cfg MixinConfig) error {
	// 读取yaml文件
	v := viper.New()
	// 设置读取的配置文件名
	v.SetConfigName(filepath.Base(path))
	// windows环境下为%GOPATH，linux环境下为$GOPATH
	v.AddConfigPath(filepath.Dir(path))
	// 设置配置文件类型
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	// 也可以直接反序列化为Struct
	if err := v.Unmarshal(cfg); err != nil {
		return err
	}
	return nil
}

type EmailConfig struct {
	From  string
	Title string
	Host  string
	Port  int
	Usr   string
	Psw   string
}
type AppConfig struct {
	Name               string // 程序中写死
	Version            string // 程序中写死
	CdnUrl             string
	BaseUrl            string
	Theme              string
	AdminSalt          string
	AdminJwtSignKey    string
	AdminJwtExpireHour int64
	AdminJwtAesKey     string
}
type ServerConfig struct {
	Port            int
	BodyLimit       int // 单位：M
	ShutDownWaitSec int
	SnowflakeNode   int64
}
type LoggerConfig struct {
	Dev     bool
	FileLog string // 只有dev=false的时候，才有用
	Router  bool
}
type DbConfig struct {
	Host          string
	Port          int
	DbName        string
	Usr           string
	Psw           string
	MaxConnection int
	MaxIdleConns  int
	MaxLifetime   time.Duration
	MaxOpenConns  int
	Debug         bool
}
