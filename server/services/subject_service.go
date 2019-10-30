
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var SubjectService = &subjectService {}

type subjectService struct {
}

func (this *subjectService) Get(id int64) *model.Subject {
	ret := &model.Subject{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *subjectService) Take(where ...interface{}) *model.Subject {
	ret := &model.Subject{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *subjectService) QueryCnd(cnd *simple.SqlCnd) (list []model.Subject, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *subjectService) Query(params *simple.QueryParams) (list []model.Subject, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.Subject{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *subjectService) Create(t *model.Subject) (*model.Subject, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *subjectService) Update(t *model.Subject) error {
	return simple.DB().Save(t).Error
}

func (this *subjectService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.Subject{}).Where("id = ?", id).Updates(columns).Error
}

func (this *subjectService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.Subject{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *subjectService) Delete(id int64) error {
	return simple.DB().Delete(&model.Subject{}, "id = ?", id).Error
}

