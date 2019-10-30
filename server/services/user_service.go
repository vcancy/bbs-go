
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var UserService = &userService {}

type userService struct {
}

func (this *userService) Get(id int64) *model.User {
	ret := &model.User{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *userService) Take(where ...interface{}) *model.User {
	ret := &model.User{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *userService) QueryCnd(cnd *simple.SqlCnd) (list []model.User, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *userService) Query(params *simple.QueryParams) (list []model.User, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.User{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *userService) Create(t *model.User) (*model.User, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *userService) Update(t *model.User) error {
	return simple.DB().Save(t).Error
}

func (this *userService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.User{}).Where("id = ?", id).Updates(columns).Error
}

func (this *userService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.User{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *userService) Delete(id int64) error {
	return simple.DB().Delete(&model.User{}, "id = ?", id).Error
}

