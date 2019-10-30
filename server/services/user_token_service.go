
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var UserTokenService = &userTokenService {}

type userTokenService struct {
}

func (this *userTokenService) Get(id int64) *model.UserToken {
	ret := &model.UserToken{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *userTokenService) Take(where ...interface{}) *model.UserToken {
	ret := &model.UserToken{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *userTokenService) QueryCnd(cnd *simple.SqlCnd) (list []model.UserToken, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *userTokenService) Query(params *simple.QueryParams) (list []model.UserToken, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.UserToken{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *userTokenService) Create(t *model.UserToken) (*model.UserToken, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *userTokenService) Update(t *model.UserToken) error {
	return simple.DB().Save(t).Error
}

func (this *userTokenService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.UserToken{}).Where("id = ?", id).Updates(columns).Error
}

func (this *userTokenService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.UserToken{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *userTokenService) Delete(id int64) error {
	return simple.DB().Delete(&model.UserToken{}, "id = ?", id).Error
}

