
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var CollectArticleService = &collectArticleService {}

type collectArticleService struct {
}

func (this *collectArticleService) Get(id int64) *model.CollectArticle {
	ret := &model.CollectArticle{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *collectArticleService) Take(where ...interface{}) *model.CollectArticle {
	ret := &model.CollectArticle{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *collectArticleService) QueryCnd(cnd *simple.SqlCnd) (list []model.CollectArticle, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *collectArticleService) Query(params *simple.QueryParams) (list []model.CollectArticle, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.CollectArticle{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *collectArticleService) Create(t *model.CollectArticle) (*model.CollectArticle, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *collectArticleService) Update(t *model.CollectArticle) error {
	return simple.DB().Save(t).Error
}

func (this *collectArticleService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.CollectArticle{}).Where("id = ?", id).Updates(columns).Error
}

func (this *collectArticleService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.CollectArticle{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *collectArticleService) Delete(id int64) error {
	return simple.DB().Delete(&model.CollectArticle{}, "id = ?", id).Error
}

