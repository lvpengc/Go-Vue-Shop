// Package mysql 初始化mysql连接
package mysql

import (
	"log"
	"os"
	"time"

	etcdStruct "github.com/hellozouyou/easy-mall/pkg/structure/etcd"

	"github.com/jinzhu/gorm"
)

// Client mysql 的本地连接
var Client *gorm.DB

// Init DB初始化
func Init(config *etcdStruct.Gorm) {
	var err error
	Client, err = gorm.Open(config.DBType, config.DSN)
	if err != nil {
		log.Fatalf("mysql connec fail: %v", err)
	}
	Client.SingularTable(true)
	if config.Debug {
		Client.LogMode(config.Debug)
		Client.SetLogger(log.New(os.Stdout, "\r\n", 0))
	}
	Client.DB().SetMaxIdleConns(config.MaxIdleConns)
	Client.DB().SetMaxOpenConns(config.MaxOpenConns)
	Client.DB().SetConnMaxLifetime(time.Duration(config.MaxLifetime) * time.Second)
}
