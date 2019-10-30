package services

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/mlogclub/simple"

	"github.com/mlogclub/bbs-go/model"
)

var TagService = &tagService{}

type tagService struct {
}

func (this *tagService) Get(id int64) *model.Tag {
	ret := &model.Tag{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *tagService) Take(where ...interface{}) *model.Tag {
	ret := &model.Tag{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *tagService) QueryCnd(cnd *simple.SqlCnd) (list []model.Tag, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *tagService) Query(params *simple.QueryParams) (list []model.Tag, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.Tag{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *tagService) Create(t *model.Tag) (*model.Tag, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *tagService) Update(t *model.Tag) error {
	return simple.DB().Save(t).Error
}

func (this *tagService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.Tag{}).Where("id = ?", id).Updates(columns).Error
}

func (this *tagService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.Tag{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *tagService) Delete(id int64) error {
	return simple.DB().Delete(&model.Tag{}, "id = ?", id).Error
}

func (this *tagService) GetByName(name string) *model.Tag {
	if len(name) == 0 {
		return nil
	}
	return this.Take("name = ?", name)
}

// 自动完成
func (this *tagService) Autocomplete(input string) []model.Tag {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return nil
	}
	list, _ := this.QueryCnd(simple.NewSqlCnd("status = ? and name like ?", model.TagStatusOk, "%"+input+"%").Size(6))
	return list
}

func (this *tagService) GetOrCreates(db *gorm.DB, tags []string) (tagIds []int64) {
	for _, tagName := range tags {
		tagName = strings.TrimSpace(tagName)
		tag, err := TagService.GetOrCreate(db, tagName)
		if err == nil {
			tagIds = append(tagIds, tag.Id)
		}
	}
	return
}

func (this *tagService) GetOrCreate(db *gorm.DB, name string) (*model.Tag, error) {
	if len(name) == 0 {
		return nil, errors.New("标签为空")
	}
	tag := this.GetByName(name)
	if tag != nil {
		return tag, nil
	} else {
		tag = &model.Tag{
			Name:       name,
			Status:     model.TagStatusOk,
			CreateTime: simple.NowTimestamp(),
			UpdateTime: simple.NowTimestamp(),
		}
		if err := db.Create(tag).Error; err != nil {
			return nil, err
		}
		return tag, nil
	}
}

func (this *tagService) ListAll(categoryId int64) ([]model.Tag, error) {
	return this.QueryCnd(simple.NewSqlCnd("category_id = ? and status = ?", categoryId, model.TagStatusOk))
}

func (this *tagService) GetTags() []model.TagResponse {
	list, err := this.QueryCnd(simple.NewSqlCnd("status = ?", model.TagStatusOk))
	if err != nil {
		return nil
	}

	var tags []model.TagResponse
	for _, tag := range list {
		tags = append(tags, model.TagResponse{TagId: tag.Id, TagName: tag.Name})
	}
	return tags
}

func (this *tagService) GetTagInIds(tagIds []int64) []model.Tag {
	if len(tagIds) == 0 {
		return nil
	}
	var tags []model.Tag
	simple.DB().Where("id in (?)", tagIds).Find(&tags)
	return tags
}
