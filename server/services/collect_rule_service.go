
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var CollectRuleService = &collectRuleService {}

type collectRuleService struct {
}

func (this *collectRuleService) Get(id int64) *model.CollectRule {
	ret := &model.CollectRule{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *collectRuleService) Take(where ...interface{}) *model.CollectRule {
	ret := &model.CollectRule{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *collectRuleService) QueryCnd(cnd *simple.SqlCnd) (list []model.CollectRule, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *collectRuleService) Query(params *simple.QueryParams) (list []model.CollectRule, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.CollectRule{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *collectRuleService) Create(t *model.CollectRule) (*model.CollectRule, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *collectRuleService) Update(t *model.CollectRule) error {
	return simple.DB().Save(t).Error
}

func (this *collectRuleService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.CollectRule{}).Where("id = ?", id).Updates(columns).Error
}

func (this *collectRuleService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.CollectRule{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *collectRuleService) Delete(id int64) error {
	return simple.DB().Delete(&model.CollectRule{}, "id = ?", id).Error
}

