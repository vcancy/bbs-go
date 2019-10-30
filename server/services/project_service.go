
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var ProjectService = &projectService {}

type projectService struct {
}

func (this *projectService) Get(id int64) *model.Project {
	ret := &model.Project{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *projectService) Take(where ...interface{}) *model.Project {
	ret := &model.Project{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *projectService) QueryCnd(cnd *simple.SqlCnd) (list []model.Project, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *projectService) Query(params *simple.QueryParams) (list []model.Project, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.Project{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *projectService) Create(t *model.Project) (*model.Project, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *projectService) Update(t *model.Project) error {
	return simple.DB().Save(t).Error
}

func (this *projectService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.Project{}).Where("id = ?", id).Updates(columns).Error
}

func (this *projectService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.Project{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *projectService) Delete(id int64) error {
	return simple.DB().Delete(&model.Project{}, "id = ?", id).Error
}

