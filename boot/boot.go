package boot

import "github.com/gogf/gf/frame/g"

func init() {
    g.Server().SetPort(8199)
    g.Server().SetDumpRouteMap(false)
    g.Server().SetAccessLogEnabled(true)
}