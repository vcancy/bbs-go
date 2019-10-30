
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var TopicTagService = &topicTagService {}

type topicTagService struct {
}

func (this *topicTagService) Get(id int64) *model.TopicTag {
	ret := &model.TopicTag{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *topicTagService) Take(where ...interface{}) *model.TopicTag {
	ret := &model.TopicTag{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *topicTagService) QueryCnd(cnd *simple.SqlCnd) (list []model.TopicTag, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *topicTagService) Query(params *simple.QueryParams) (list []model.TopicTag, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.TopicTag{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *topicTagService) Create(t *model.TopicTag) (*model.TopicTag, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *topicTagService) Update(t *model.TopicTag) error {
	return simple.DB().Save(t).Error
}

func (this *topicTagService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.TopicTag{}).Where("id = ?", id).Updates(columns).Error
}

func (this *topicTagService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.TopicTag{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *topicTagService) Delete(id int64) error {
	return simple.DB().Delete(&model.TopicTag{}, "id = ?", id).Error
}

