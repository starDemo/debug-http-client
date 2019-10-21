package router

import (
    "debug-http-client/app/api/alertmanager"
    "debug-http-client/app/api/hello"
    "github.com/gogf/gf/frame/g"
)

// 统一路由注册.
func init() {
    g.Server().BindHandler("POST:/debug", hello.Handler)
    g.Server().BindHandler("POST:/alert", alertmanager.Alert)
    //TODO SKYEALKING报警接入
}
