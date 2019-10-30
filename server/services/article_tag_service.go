package services

import (
	"github.com/jinzhu/gorm"
	"github.com/mlogclub/simple"

	"github.com/mlogclub/bbs-go/model"
)

var ArticleTagService = &articleTagService{}

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

func (this *articleTagService) DeleteByArticleId(topicId int64) {
	simple.DB().Model(model.ArticleTag{}).Where("topic_id = ?", topicId).UpdateColumn("status", model.ArticleTagStatusDeleted)
}

func (this *articleTagService) CreateArticleTags(db *gorm.DB, articleId int64, tagIds []int64) {
	if articleId <= 0 || len(tagIds) == 0 {
		return
	}

	for _, tagId := range tagIds {
		db.Create(&model.ArticleTag{
			ArticleId:  articleId,
			TagId:      tagId,
			CreateTime: simple.NowTimestamp(),
		})
	}
}

func (this *articleTagService) DeleteArticleTags(db *gorm.DB, articleId int64) {
	if articleId <= 0 {
		return
	}

	db.Where("article_id = ?", articleId).Delete(model.ArticleTag{})
}

func (this *articleTagService) GetUnique(db *gorm.DB, articleId, tagId int64) *model.ArticleTag {
	ret := &model.ArticleTag{}
	if err := db.First(ret, "article_id = ? and tag_id = ?", articleId, tagId).Error; err != nil {
		return nil
	}
	return ret
}

func (this *articleTagService) GetByArticleId(db *gorm.DB, articleId int64) ([]model.ArticleTag, error) {
	return this.QueryCnd(simple.NewSqlCnd("article_id = ?", articleId))
}
