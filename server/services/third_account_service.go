
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var ThirdAccountService = &thirdAccountService {}

type thirdAccountService struct {
}

func (this *thirdAccountService) Get(id int64) *model.ThirdAccount {
	ret := &model.ThirdAccount{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *thirdAccountService) Take(where ...interface{}) *model.ThirdAccount {
	ret := &model.ThirdAccount{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *thirdAccountService) QueryCnd(cnd *simple.SqlCnd) (list []model.ThirdAccount, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *thirdAccountService) Query(params *simple.QueryParams) (list []model.ThirdAccount, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.ThirdAccount{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *thirdAccountService) Create(t *model.ThirdAccount) (*model.ThirdAccount, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *thirdAccountService) Update(t *model.ThirdAccount) error {
	return simple.DB().Save(t).Error
}

func (this *thirdAccountService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.ThirdAccount{}).Where("id = ?", id).Updates(columns).Error
}

func (this *thirdAccountService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.ThirdAccount{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *thirdAccountService) Delete(id int64) error {
	return simple.DB().Delete(&model.ThirdAccount{}, "id = ?", id).Error
}

