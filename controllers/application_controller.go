package controllers

import (
	"encoding/json"
	"github.com/oam-backend/modules"
	"time"

	"github.com/astaxie/beego"
	"github.com/oam-backend/common"
	"github.com/oam-backend/models"
	"github.com/oam-backend/modules/cache"
)

type APPController struct {
	beego.Controller
}

// @Title CreateApp
// @Description  创建 APP
// @Success 200  {object} common.Response  "统一返回结构"
// @router /app [post]
func (c *APPController) CreateApp() {
	input := models.ApplicationMetaInfo{}
	resp := common.Response{}
	defer func() {
		c.Data["json"] = resp
		c.ServeJSON()
	}()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &input); err != nil {
		resp = resp.WarnDump(common.PARAM_ERROR, err.Error())
		return
	}
	input.CreatedTime = time.Now()
	input.UpdatedTime = input.CreatedTime
	resp = cache.StoreApp(input)
	return
}

// @Title GetApps
// @Description  获取APP 信息
// @Success 200  {object} common.Response  "统一返回结构"
// @router /app [get]
func (c *APPController) GetApps() {
	resp := common.Response{}
	defer func() {
		c.Data["json"] = resp
		c.ServeJSON()
	}()
	resp = cache.GetApps()
	return
}


// @Title CreateApp
// @Description  创建 APP
// @Success 200  {object} common.Response  "统一返回结构"
// @router /app/changeSheet [post]
func (c *APPController) CreateChangeSheet() {
	input := models.ChangeSheet{}
	resp := common.Response{}
	defer func() {
		c.Data["json"] = resp
		c.ServeJSON()
	}()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &input); err != nil {
		resp = resp.WarnDump(common.PARAM_ERROR, err.Error())
		return
	}
	resp = modules.HandleTask(input)
	return
}

// @Title GetApps
// @Description  获取APP 信息
// @Success 200  {object} common.Response  "统一返回结构"
// @router /app/changeSheet [get]
func (c *APPController) GetChangeSheet() {
	resp := common.Response{}
	defer func() {
		c.Data["json"] = resp
		c.ServeJSON()
	}()
	resp = cache.GetChangeSheet()
	return
}