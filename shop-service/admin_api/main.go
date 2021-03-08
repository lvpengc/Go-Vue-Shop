// admin_api 管理后台后端API
package main

import (
	"net/http"
	"time"

	"github.com/hellozouyou/easy-mall/admin_api/middleware"
	"github.com/hellozouyou/easy-mall/admin_api/router"
	"github.com/hellozouyou/easy-mall/pkg/casbin"
	"github.com/hellozouyou/easy-mall/pkg/constant"
	"github.com/hellozouyou/easy-mall/pkg/etcd"
	"github.com/hellozouyou/easy-mall/pkg/mysql"
	"github.com/hellozouyou/easy-mall/pkg/redis"
	etcdStruct "github.com/hellozouyou/easy-mall/pkg/structure/etcd"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
)

var (
	redisCfg = &etcdStruct.Redis{}
	mysqlCfg = &etcdStruct.Mysql{}
	webCfg   = &etcdStruct.Web{}
	gormCfg  = &etcdStruct.Gorm{}
)

func main() {
	// 初始化etcd
	etcd.Init()
	// 从etcd获取配置信息
	etcd.Scan(mysqlCfg, "easymall", "config", "mysql")
	etcd.Scan(webCfg, "easymall", "admin_api", "web")
	etcd.Scan(gormCfg, "easymall", "admin_api", "gorm")
	etcd.Scan(redisCfg, "easymall", "config", "redis")
	// 将mysql信息写入gorm的连接
	gormCfg.DSN = mysqlCfg.DSN()
	// 初始化gorm
	mysql.Init(gormCfg)
	// 初始化redis
	redis.Init(redisCfg)
	// 初始化casbin
	casbin.InitEnforcer()
	// 启动web服务
	initWeb()
}

// initWeb 启动 gin web 服务
func initWeb() {
	// 确定gin是否debug
	gin.SetMode(webCfg.SetMode)
	// 初始化gin
	app := gin.Default()
	// 404 路由找不到
	app.NoRoute(middleware.NoRouteHandler())
	// 405 方法不允许
	app.NoMethod(middleware.NoMethodHandler())
	// 崩溃恢复
	app.Use(middleware.RecoveryMiddleware())
	// 注册路由
	router.Init(app)

	// 定义http.Server参数
	server := &http.Server{
		Addr:         webCfg.Addr,
		ReadTimeout:  time.Duration(webCfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(webCfg.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(webCfg.IdleTimeout) * time.Second,
		Handler:      app,
	}
	// 初始化micro的web服务
	service := web.NewService(
		web.Address(webCfg.Addr),
		web.Name(constant.WEBADMINAPI),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*10),
		web.Server(server),
	)
	// 将web服务的所有路由指向到gin
	service.Handle("/", app)
	log.Infof("service running on %v", webCfg.Addr)
	// 启动运行服务
	if err := service.Run(); err != nil {
		panic(err)
	}
}
