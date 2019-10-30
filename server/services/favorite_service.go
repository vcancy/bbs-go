
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var FavoriteService = &favoriteService {}

type favoriteService struct {
}

func (this *favoriteService) Get(id int64) *model.Favorite {
	ret := &model.Favorite{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *favoriteService) Take(where ...interface{}) *model.Favorite {
	ret := &model.Favorite{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *favoriteService) QueryCnd(cnd *simple.SqlCnd) (list []model.Favorite, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *favoriteService) Query(params *simple.QueryParams) (list []model.Favorite, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.Favorite{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *favoriteService) Create(t *model.Favorite) (*model.Favorite, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *favoriteService) Update(t *model.Favorite) error {
	return simple.DB().Save(t).Error
}

func (this *favoriteService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.Favorite{}).Where("id = ?", id).Updates(columns).Error
}

func (this *favoriteService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.Favorite{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *favoriteService) Delete(id int64) error {
	return simple.DB().Delete(&model.Favorite{}, "id = ?", id).Error
}

