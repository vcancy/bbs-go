
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var CommentService = &commentService {}

type commentService struct {
}

func (this *commentService) Get(id int64) *model.Comment {
	ret := &model.Comment{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *commentService) Take(where ...interface{}) *model.Comment {
	ret := &model.Comment{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *commentService) QueryCnd(cnd *simple.SqlCnd) (list []model.Comment, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *commentService) Query(params *simple.QueryParams) (list []model.Comment, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.Comment{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *commentService) Create(t *model.Comment) (*model.Comment, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *commentService) Update(t *model.Comment) error {
	return simple.DB().Save(t).Error
}

func (this *commentService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.Comment{}).Where("id = ?", id).Updates(columns).Error
}

func (this *commentService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.Comment{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *commentService) Delete(id int64) error {
	return simple.DB().Delete(&model.Comment{}, "id = ?", id).Error
}

