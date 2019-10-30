
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var ArticleTagService = &articleTagService {}

type articleTagService struct {
}

func (this *articleTagService) Get(id int64) *model.ArticleTag {
	ret := &model.ArticleTag{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *articleTagService) Take(where ...interface{}) *model.ArticleTag {
	ret := &model.ArticleTag{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *articleTagService) QueryCnd(cnd *simple.SqlCnd) (list []model.ArticleTag, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *articleTagService) Query(params *simple.QueryParams) (list []model.ArticleTag, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.ArticleTag{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *articleTagService) Create(t *model.ArticleTag) (*model.ArticleTag, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *articleTagService) Update(t *model.ArticleTag) error {
	return simple.DB().Save(t).Error
}

func (this *articleTagService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.ArticleTag{}).Where("id = ?", id).Updates(columns).Error
}

func (this *articleTagService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.ArticleTag{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *articleTagService) Delete(id int64) error {
	return simple.DB().Delete(&model.ArticleTag{}, "id = ?", id).Error
}

