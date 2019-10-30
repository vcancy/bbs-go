
package admin

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/bbs-go/services2"
	"github.com/mlogclub/simple"
	"github.com/kataras/iris"
	"strconv"
)

type ThirdAccountController struct {
	Ctx             iris.Context
}

func (this *ThirdAccountController) GetBy(id int64) *simple.JsonResult {
	t := services2.ThirdAccountService.Get(id)
	if t == nil {
		return simple.JsonErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return simple.JsonData(t)
}

func (this *ThirdAccountController) AnyList() *simple.JsonResult {
	list, paging := services2.ThirdAccountService.Query(simple.NewParamQueries(this.Ctx).PageAuto().Desc("id"))
	return simple.JsonData(&simple.PageResult{Results: list, Page: paging})
}

func (this *ThirdAccountController) PostCreate() *simple.JsonResult {
	t := &model.ThirdAccount{}
	err := this.Ctx.ReadForm(t)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	err = services2.ThirdAccountService.Create(t)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(t)
}

func (this *ThirdAccountController) PostUpdate() *simple.JsonResult {
	id, err := simple.FormValueInt64(this.Ctx, "id")
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	t := services2.ThirdAccountService.Get(id)
	if t == nil {
		return simple.JsonErrorMsg("entity not found")
	}

	err = this.Ctx.ReadForm(t)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	err = services2.ThirdAccountService.Update(t)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(t)
}

