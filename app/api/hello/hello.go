package hello

import (
    "debug-http-client/library/response"
    "github.com/gogf/gf/net/ghttp"
    "github.com/gogf/gf/os/glog"
)

// Hello World
func Handler(r *ghttp.Request) {
    data := r.GetPostMapStrStr()
    glog.Info(data)
    response.Json(r, 0, "ok")
}
