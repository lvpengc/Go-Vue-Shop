package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/micro/go-micro/client"
	app_api "path/to/service/proto/app_api"
)

func App_apiCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	app_apiClient := app_api.NewApp_apiService("easymall.srv.app_api", client.DefaultClient)
	rsp, err := app_apiClient.Call(context.TODO(), &app_api.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
