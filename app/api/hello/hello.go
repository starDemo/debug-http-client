package hello

import (
	"debug-http-client/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

// Hello World
func Handler(r *ghttp.Request) {
    data, err := r.GetJson()
    if err != nil{
    	glog.Error(err)
	}
    glog.Info(data)
    response.Json(r, 0, "ok")
}
