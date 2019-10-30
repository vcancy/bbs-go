
package admin

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/bbs-go/services2"
	"github.com/mlogclub/simple"
	"github.com/kataras/iris"
	"strconv"
)

type TopicLikeController struct {
	Ctx             iris.Context
}

func (this *TopicLikeController) GetBy(id int64) *simple.JsonResult {
	t := services2.TopicLikeService.Get(id)
	if t == nil {
		return simple.JsonErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return simple.JsonData(t)
}

func (this *TopicLikeController) AnyList() *simple.JsonResult {
	list, paging := services2.TopicLikeService.Query(simple.NewParamQueries(this.Ctx).PageAuto().Desc("id"))
	return simple.JsonData(&simple.PageResult{Results: list, Page: paging})
}

func (this *TopicLikeController) PostCreate() *simple.JsonResult {
	t := &model.TopicLike{}
	err := this.Ctx.ReadForm(t)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	err = services2.TopicLikeService.Create(t)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(t)
}

func (this *TopicLikeController) PostUpdate() *simple.JsonResult {
	id, err := simple.FormValueInt64(this.Ctx, "id")
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	t := services2.TopicLikeService.Get(id)
	if t == nil {
		return simple.JsonErrorMsg("entity not found")
	}

	err = this.Ctx.ReadForm(t)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	err = services2.TopicLikeService.Update(t)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(t)
}

