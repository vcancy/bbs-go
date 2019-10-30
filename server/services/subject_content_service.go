package services

import (
	"github.com/mlogclub/simple"
	"github.com/sirupsen/logrus"

	"github.com/mlogclub/bbs-go/common"
	"github.com/mlogclub/bbs-go/common/subject"
	"github.com/mlogclub/bbs-go/model"
)

var SubjectContentService = &subjectContentService{}

type subjectContentService struct {
}

func (this *subjectContentService) Get(id int64) *model.SubjectContent {
	ret := &model.SubjectContent{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *subjectContentService) Take(where ...interface{}) *model.SubjectContent {
	ret := &model.SubjectContent{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *subjectContentService) QueryCnd(cnd *simple.SqlCnd) (list []model.SubjectContent, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *subjectContentService) Query(params *simple.QueryParams) (list []model.SubjectContent, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.SubjectContent{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *subjectContentService) Create(t *model.SubjectContent) (*model.SubjectContent, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *subjectContentService) Update(t *model.SubjectContent) error {
	return simple.DB().Save(t).Error
}

func (this *subjectContentService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.SubjectContent{}).Where("id = ?", id).Updates(columns).Error
}

func (this *subjectContentService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.SubjectContent{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *subjectContentService) Delete(id int64) error {
	return simple.DB().Delete(&model.SubjectContent{}, "id = ?", id).Error
}

func (this *subjectContentService) GetByEntity(entityType string, entityId int64) *model.SubjectContent {
	return this.Take("entity_type = ? and entity_id = ?", entityType, entityId)
}

// 分析文章
func (this *subjectContentService) AnalyzeArticle(article *model.Article) {
	subjectIds := subject.AnalyzeSubjects(article.UserId, article.Title, article.Content)
	if len(subjectIds) > 0 {
		for _, subjectId := range subjectIds {
			summary := article.Summary
			if summary == "" {
				summary = common.GetSummary(article.ContentType, article.Content)
			}
			_, err := this.Publish(subjectId, model.EntityTypeArticle, article.Id,
				article.Title, summary)
			if err != nil {
				logrus.Error(err)
			}
		}
	}
}

// 发布
func (this *subjectContentService) Publish(subjectId int64, entityType string, entityId int64, title, summary string) (c *model.SubjectContent, err error) {
	c = this.GetByEntity(entityType, entityId)
	if c != nil {
		c.SubjectId = subjectId
		c.EntityType = entityType
		c.EntityId = entityId
		c.Title = title
		c.Summary = summary
		c.Deleted = false
		c.CreateTime = simple.NowTimestamp()
		err = this.Update(c)
	} else {
		c := &model.SubjectContent{
			SubjectId:  subjectId,
			EntityType: entityType,
			EntityId:   entityId,
			Title:      title,
			Summary:    summary,
			Deleted:    false,
			CreateTime: simple.NowTimestamp(),
		}
		c, err = this.Create(c)
	}
	return
}

func (this *subjectContentService) DeleteByEntity(entityType string, entityId int64) {
	t := this.GetByEntity(entityType, entityId)
	if t != nil {
		_ = this.Delete(t.Id)
	}
}
