package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

func init() {
	glog.SetPath("/tmp/debug-client")
	glog.SetStdoutPrint(false)
    g.Server().SetPort(8199)
    g.Server().SetDumpRouteMap(false)
    g.Server().SetAccessLogEnabled(true)
}