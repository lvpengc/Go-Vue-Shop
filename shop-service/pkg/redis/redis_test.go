package redis_test

import (
	"testing"

	"github.com/hellozouyou/easy-mall/pkg/redis"
	jsoniter "github.com/json-iterator/go"
)

func TestRedisPool(t *testing.T) {
	m := struct {
		host         string
		password     string
		database     int
		maxOpenConns int
		maxIdleConns int
	}{
		host:         "127.0.0.1:6379",
		password:     "",
		database:     0,
		maxOpenConns: 10,
		maxIdleConns: 2,
	}

	r := redis.InitRedisPool(m.host, m.password, m.database, m.maxOpenConns, m.maxIdleConns)

	// 测试结构体
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	//json序列化
	type Movie struct {
		Title  string
		Year   int  `json:"released"`
		Color  bool `json:"color,omitempty"`
		Actors []string
	}

	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
		// ...
	}
	data, err := json.Marshal(movies)
	if err != nil {
		t.Fatalf("JSON marshaling failed: %s", err)
	}
	t.Logf("JSON: %s\n", data)

	// redis设置
	// v, err := r.Do("SET", "struct", data)
	v, err := r.SetString("struct", data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("redis set struct: %s\n", v)

	// redis获取值
	v2, _ := r.GetBytes("struct")
	t.Logf("redis get struct GetBytes: %s\n", v2)

	//json反序列化
	var data2 []Movie
	json.Unmarshal(v2, &data2)
	t.Logf("redis Unmarshal: %v\n", data2)

	var titles []struct{ Title string }
	if err := json.Unmarshal(v2, &titles); err != nil {
		t.Fatalf("JSON unmarshaling failed: %s", err)
	}
	t.Logf("JSON titles: %s\n", titles)

	// 测试整数
	v, err = r.Do("SET", "test", 123)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("redis set int: %v\n", v)

	// 自增
	r.Do("INCR", "test")

	// 获取值
	v, err = r.GetString("test")
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("redis get string: %s\n", v)

	v, err = r.GetInt("test")
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("redis get int: %d\n", v)

	// 测试删除
	v, err = r.DelKey("test")
	if err != nil {
		t.Fatalf(err.Error())
	}
	v, err = r.GetString("test")
	t.Logf("redis get deleted key: %v\n", v)
}
