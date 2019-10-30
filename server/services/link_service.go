
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var LinkService = &linkService {}

type linkService struct {
}

func (this *linkService) Get(id int64) *model.Link {
	ret := &model.Link{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *linkService) Take(where ...interface{}) *model.Link {
	ret := &model.Link{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *linkService) QueryCnd(cnd *simple.SqlCnd) (list []model.Link, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *linkService) Query(params *simple.QueryParams) (list []model.Link, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.Link{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *linkService) Create(t *model.Link) (*model.Link, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *linkService) Update(t *model.Link) error {
	return simple.DB().Save(t).Error
}

func (this *linkService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.Link{}).Where("id = ?", id).Updates(columns).Error
}

func (this *linkService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.Link{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *linkService) Delete(id int64) error {
	return simple.DB().Delete(&model.Link{}, "id = ?", id).Error
}

