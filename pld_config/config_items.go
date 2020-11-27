package pld_config

import "time"

// Web服务器配置
type WebServerConfig struct {
	Port            int
	BodyLimit       int // 单位：M
	ShutDownWaitSec int
}

// 应用全局配置
type AppConfig struct {
	Name    string // 程序中写死
	Version string // 程序中写死
	CdnUrl  string
	BaseUrl string
	Theme   string
}

// 日志配置
type LoggerConfig struct {
	Mode string // prod or dev
	// FileLog string // 只有dev=false的时候，才有用
	// Router  bool
}

// Snowflake配置
type SnowflakeConfig struct {
	Node int64
}

// MySQL配置
type MysqlConfig struct {
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

// Email配置
type EmailConfig struct {
	From  string
	Title string
	Host  string
	Port  int
	Usr   string
	Psw   string
}
