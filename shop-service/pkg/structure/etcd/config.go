// Package etcd config.go 定义etcd中配置数据的格式
package etcd

import "fmt"

// Redis REDIS连接配置
// key: /yun/system/config/redis
type Redis struct {
	Host         string `json:"host"`
	Password     string `json:"password"`
	Database     int    `json:"database"`
	MaxOpenConns int    `json:"maxOpenConns"`
	MaxIdleConns int    `json:"maxIdleConns"`
}

// Mysql 数据库连接配置
// key: /yun/system/config/mysql
type Mysql struct {
	Host     string `json:"host"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
	Charset  string `json:"charset"`
}

// Nsq NSQ的连接配置
// key: /yun/system/config/nsq
type Nsq struct {
	Host string `json:"host"`
}

// Web 站点配置参数
type Web struct {
	Addr         string `json:"addr"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
	IdleTimeout  int    `json:"idle_timeout"`
	SetMode      string `json:"set_mode"`
}

// Gorm gorm配置参数
type Gorm struct {
	Debug        bool   `json:"debug"`
	DBType       string `json:"db_type"`
	MaxLifetime  int    `json:"max_lifetime"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
	DSN          string `json:"dsn"`
}

// DSN MySQL 数据库连接串
func (m *Mysql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", m.User, m.Password, m.Host, m.Database, m.Charset) + "&parseTime=True&loc=Asia%2FShanghai"
}
