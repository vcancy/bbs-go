
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var CategoryService = &categoryService {}

type categoryService struct {
}

func (this *categoryService) Get(id int64) *model.Category {
	ret := &model.Category{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *categoryService) Take(where ...interface{}) *model.Category {
	ret := &model.Category{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *categoryService) QueryCnd(cnd *simple.SqlCnd) (list []model.Category, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *categoryService) Query(params *simple.QueryParams) (list []model.Category, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.Category{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *categoryService) Create(t *model.Category) (*model.Category, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *categoryService) Update(t *model.Category) error {
	return simple.DB().Save(t).Error
}

func (this *categoryService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.Category{}).Where("id = ?", id).Updates(columns).Error
}

func (this *categoryService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.Category{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *categoryService) Delete(id int64) error {
	return simple.DB().Delete(&model.Category{}, "id = ?", id).Error
}

