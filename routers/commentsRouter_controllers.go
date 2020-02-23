package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/oam-apiserver/controllers:APPController"] = append(beego.GlobalControllerRouter["github.com/oam-apiserver/controllers:APPController"],
        beego.ControllerComments{
            Method: "CreateApp",
            Router: `/app`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/oam-apiserver/controllers:APPController"] = append(beego.GlobalControllerRouter["github.com/oam-apiserver/controllers:APPController"],
        beego.ControllerComments{
            Method: "GetApps",
            Router: `/app`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/oam-apiserver/controllers:APPController"] = append(beego.GlobalControllerRouter["github.com/oam-apiserver/controllers:APPController"],
        beego.ControllerComments{
            Method: "CreateChangeSheet",
            Router: `/app/changeSheet`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/oam-apiserver/controllers:APPController"] = append(beego.GlobalControllerRouter["github.com/oam-apiserver/controllers:APPController"],
        beego.ControllerComments{
            Method: "GetChangeSheet",
            Router: `/app/changeSheet`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
