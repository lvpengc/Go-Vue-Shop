package etcd_test

import (
	"testing"
	"time"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/store"
	storeEtcd "github.com/micro/go-micro/store/etcd"

	"github.com/hellozouyou/easy-mall/pkg/etcd"
)

var (
	conf config.Config
	data = &TestData{}
	s    = storeEtcd.NewStore(store.Nodes("etcd:2379"))
)

// TestData 测试数据格式
type TestData struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// TestEtcd 测试etcd连接并读取数据
func TestEtcd(t *testing.T) {
	var err error
	// 初始化etcd
	etcd.Init()
	t.Log("etcd 连接成功")
	// 获取测试数据
	err = etcd.Scan(data, "test")
	go func() {
		etcd.Watch(data, "test")
	}()
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second)
	t.Logf("测试数据初始值：%s", *data)
	// 主动写入测试数据
	writeData(t, &store.Record{
		Key:   "/test",
		Value: []byte(`{"name": "a", "value": "1"}`),
	})
	time.Sleep(time.Second)
	t.Logf("watch 测试数据：%s", *data)
	// 更改测试数据
	writeData(t, &store.Record{
		Key:   "/test",
		Value: []byte(`{"name": "b", "value": "3"}`),
	})
	time.Sleep(time.Second)
	t.Logf("watch 测试数据：%s", *data)
	time.Sleep(time.Second)
}

// writeData 写入测试数据到etcd
func writeData(t *testing.T, data *store.Record) {
	var err error
	t.Logf("写入测试数据：%s", *data)
	err = s.Write(data)
	if err != nil {
		t.Errorf("测试数据写入失败：%s", err)
	}
}
