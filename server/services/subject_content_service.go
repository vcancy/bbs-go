
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var SubjectContentService = &subjectContentService {}

type subjectContentService struct {
}

func (this *subjectContentService) Get(id int64) *model.SubjectContent {
	ret := &model.SubjectContent{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *subjectContentService) Take(where ...interface{}) *model.SubjectContent {
	ret := &model.SubjectContent{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *subjectContentService) QueryCnd(cnd *simple.SqlCnd) (list []model.SubjectContent, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *subjectContentService) Query(params *simple.QueryParams) (list []model.SubjectContent, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.SubjectContent{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *subjectContentService) Create(t *model.SubjectContent) (*model.SubjectContent, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *subjectContentService) Update(t *model.SubjectContent) error {
	return simple.DB().Save(t).Error
}

func (this *subjectContentService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.SubjectContent{}).Where("id = ?", id).Updates(columns).Error
}

func (this *subjectContentService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.SubjectContent{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *subjectContentService) Delete(id int64) error {
	return simple.DB().Delete(&model.SubjectContent{}, "id = ?", id).Error
}

