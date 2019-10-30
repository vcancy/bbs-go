package admin

import (
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/mlogclub/bbs-go/services2"
	"github.com/mlogclub/simple"
	"strconv"
)

type SysConfigController struct {
	Ctx iris.Context
}

func (this *SysConfigController) GetBy(id int64) *simple.JsonResult {
	t := services2.SysConfigService.Get(id)
	if t == nil {
		return simple.JsonErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return simple.JsonData(t)
}

func (this *SysConfigController) AnyList() *simple.JsonResult {
	list, paging := services2.SysConfigService.Query(simple.NewParamQueries(this.Ctx).PageAuto().Desc("id"))
	return simple.JsonData(&simple.PageResult{Results: list, Page: paging})
}

func (this *SysConfigController) GetAll() *simple.JsonResult {
	list := services2.SysConfigService.GetAll()
	return simple.JsonData(list)
}

func (this *SysConfigController) PostSave() *simple.JsonResult {
	config := this.Ctx.FormValue("config")
	data := make(map[string]string)
	err := json.Unmarshal([]byte(config), &data)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	err = services2.SysConfigService.SetAll(data)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonSuccess()
}
