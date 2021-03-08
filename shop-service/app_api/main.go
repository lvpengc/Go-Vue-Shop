package main

import (
	"net/http"

	"github.com/micro/go-micro/util/log"

	"app_api/handler"

	"github.com/micro/go-micro/web"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("easymall.web.app_api"),
		web.Version("latest"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/app_api/call", handler.App_apiCall)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
