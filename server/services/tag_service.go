
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var TagService = &tagService {}

type tagService struct {
}

func (this *tagService) Get(id int64) *model.Tag {
	ret := &model.Tag{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *tagService) Take(where ...interface{}) *model.Tag {
	ret := &model.Tag{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *tagService) QueryCnd(cnd *simple.SqlCnd) (list []model.Tag, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *tagService) Query(params *simple.QueryParams) (list []model.Tag, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.Tag{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *tagService) Create(t *model.Tag) (*model.Tag, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *tagService) Update(t *model.Tag) error {
	return simple.DB().Save(t).Error
}

func (this *tagService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.Tag{}).Where("id = ?", id).Updates(columns).Error
}

func (this *tagService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.Tag{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *tagService) Delete(id int64) error {
	return simple.DB().Delete(&model.Tag{}, "id = ?", id).Error
}

