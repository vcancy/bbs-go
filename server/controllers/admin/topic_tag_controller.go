package admin

import (
	"github.com/kataras/iris"
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/bbs-go/services2"
	"github.com/mlogclub/simple"
	"strconv"
)

type TopicTagController struct {
	Ctx iris.Context
}

func (this *TopicTagController) GetBy(id int64) *simple.JsonResult {
	t := services2.TopicTagService.Get(id)
	if t == nil {
		return simple.JsonErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return simple.JsonData(t)
}

func (this *TopicTagController) AnyList() *simple.JsonResult {
	list, paging := services2.TopicTagService.Query(simple.NewParamQueries(this.Ctx).PageAuto().Desc("id"))
	return simple.JsonData(&simple.PageResult{Results: list, Page: paging})
}

func (this *TopicTagController) PostCreate() *simple.JsonResult {
	t := &model.TopicTag{}
	this.Ctx.ReadForm(t)

	err := services2.TopicTagService.Create(t)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(t)
}

func (this *TopicTagController) PostUpdate() *simple.JsonResult {
	id, err := simple.FormValueInt64(this.Ctx, "id")
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	t := services2.TopicTagService.Get(id)
	if t == nil {
		return simple.JsonErrorMsg("entity not found")
	}

	this.Ctx.ReadForm(t)

	err = services2.TopicTagService.Update(t)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(t)
}
