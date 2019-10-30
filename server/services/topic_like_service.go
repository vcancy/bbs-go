
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var TopicLikeService = &topicLikeService {}

type topicLikeService struct {
}

func (this *topicLikeService) Get(id int64) *model.TopicLike {
	ret := &model.TopicLike{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *topicLikeService) Take(where ...interface{}) *model.TopicLike {
	ret := &model.TopicLike{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *topicLikeService) QueryCnd(cnd *simple.SqlCnd) (list []model.TopicLike, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *topicLikeService) Query(params *simple.QueryParams) (list []model.TopicLike, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.TopicLike{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *topicLikeService) Create(t *model.TopicLike) (*model.TopicLike, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *topicLikeService) Update(t *model.TopicLike) error {
	return simple.DB().Save(t).Error
}

func (this *topicLikeService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.TopicLike{}).Where("id = ?", id).Updates(columns).Error
}

func (this *topicLikeService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.TopicLike{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *topicLikeService) Delete(id int64) error {
	return simple.DB().Delete(&model.TopicLike{}, "id = ?", id).Error
}

