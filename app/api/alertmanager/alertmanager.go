package alertmanager

import (
	"debug-http-client/library/response"
	"debug-http-client/library/types"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)
var alertmanager types.Alertmanager
func Alert(r *ghttp.Request){
	if err:=gjson.DecodeTo(r.GetRaw(), &alertmanager); err != nil {
		glog.Error("返回JSON 反序列化失败 ",err)
	}
	glog.Info(alertmanager.Status,":",alertmanager.Alerts[0].Labels.AlertName)
	response.Json(r, 0, "ok")
}
