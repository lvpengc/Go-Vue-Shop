// Package etcd 初始化etcd连接并提供相关操作函数
package etcd

import (
	"flag"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/etcd"
	"github.com/micro/go-micro/util/log"
)

var (
	conf     config.Config
	etcdHost = flag.String("http", "etcd:2379", "etcd host for configure")
)

func init() {
	flag.Parse()
}

// Init 初始化etcd
func Init() {
	// var etcdHost string
	// etcdHost = os.Getenv("ETCD_HOST")

	// 初始化etcd接入
	etcdSource := etcd.NewSource(
		// optionally specify etcd address; default to localhost:8500
		etcd.WithAddress(*etcdHost),
		// optionally specify prefix; defaults to /micro/config
		etcd.WithPrefix("/"),
		// optionally strip the provided prefix from the keys, defaults to false
		etcd.StripPrefix(true),
	)
	conf = config.NewConfig()
	err := conf.Load(etcdSource)
	if err != nil {
		log.Fatalf("etcd config init fail: %s", err)
	}
	return
}

// Scan 获取指定path下的数据并watch它
func Scan(datav interface{}, path ...string) (err error) {
	v := conf.Get(path...)
	err = v.Scan(datav)
	return
}

// Watch watch在etcd数据中变化并更新到data变量
func Watch(data interface{}, path ...string) error {
	watcher, err := conf.Watch(path...)
	if err != nil {
		log.Errorf("etcd watch failed with path %s : %s", path, err)
		return err
	}
	for {
		rv, err := watcher.Next()
		if err != nil {
			log.Errorf("etcd watch failed with path %s : %s", path, err)
			return err
		}
		log.Errorf("config changed: %v", string(rv.Bytes()))
		rv.Scan(data)
	}
}
