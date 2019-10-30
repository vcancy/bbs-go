
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var ArticleService = &articleService {}

type articleService struct {
}

func (this *articleService) Get(id int64) *model.Article {
	ret := &model.Article{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *articleService) Take(where ...interface{}) *model.Article {
	ret := &model.Article{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *articleService) QueryCnd(cnd *simple.SqlCnd) (list []model.Article, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *articleService) Query(params *simple.QueryParams) (list []model.Article, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.Article{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *articleService) Create(t *model.Article) (*model.Article, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *articleService) Update(t *model.Article) error {
	return simple.DB().Save(t).Error
}

func (this *articleService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.Article{}).Where("id = ?", id).Updates(columns).Error
}

func (this *articleService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.Article{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *articleService) Delete(id int64) error {
	return simple.DB().Delete(&model.Article{}, "id = ?", id).Error
}

