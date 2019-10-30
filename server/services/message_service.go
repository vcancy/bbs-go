
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var MessageService = &messageService {}

type messageService struct {
}

func (this *messageService) Get(id int64) *model.Message {
	ret := &model.Message{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *messageService) Take(where ...interface{}) *model.Message {
	ret := &model.Message{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *messageService) QueryCnd(cnd *simple.SqlCnd) (list []model.Message, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *messageService) Query(params *simple.QueryParams) (list []model.Message, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.Message{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *messageService) Create(t *model.Message) (*model.Message, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *messageService) Update(t *model.Message) error {
	return simple.DB().Save(t).Error
}

func (this *messageService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.Message{}).Where("id = ?", id).Updates(columns).Error
}

func (this *messageService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.Message{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *messageService) Delete(id int64) error {
	return simple.DB().Delete(&model.Message{}, "id = ?", id).Error
}

