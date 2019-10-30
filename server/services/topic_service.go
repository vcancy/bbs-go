
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var TopicService = &topicService {}

type topicService struct {
}

func (this *topicService) Get(id int64) *model.Topic {
	ret := &model.Topic{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *topicService) Take(where ...interface{}) *model.Topic {
	ret := &model.Topic{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *topicService) QueryCnd(cnd *simple.SqlCnd) (list []model.Topic, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *topicService) Query(params *simple.QueryParams) (list []model.Topic, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.Topic{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *topicService) Create(t *model.Topic) (*model.Topic, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *topicService) Update(t *model.Topic) error {
	return simple.DB().Save(t).Error
}

func (this *topicService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.Topic{}).Where("id = ?", id).Updates(columns).Error
}

func (this *topicService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.Topic{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *topicService) Delete(id int64) error {
	return simple.DB().Delete(&model.Topic{}, "id = ?", id).Error
}

