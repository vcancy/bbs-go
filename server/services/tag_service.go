package services

import (
	"strings"

	"github.com/mlogclub/simple"

	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/bbs-go/repositories"
)

var TagService = newTagService()

func newTagService() *tagService {
	return &tagService{}
}

type tagService struct {
}

func (this *tagService) Get(id int64) *model.Tag {
	return repositories.TagRepository.Get(simple.DB(), id)
}

func (this *tagService) Take(where ...interface{}) *model.Tag {
	return repositories.TagRepository.Take(simple.DB(), where...)
}

func (this *tagService) Find(cnd *simple.SqlCnd) (list []model.Tag, err error) {
	return repositories.TagRepository.Find(simple.DB(), cnd)
}

func (this *tagService) FindPageByParams(params *simple.QueryParams) (list []model.Tag, paging *simple.Paging) {
	return repositories.TagRepository.FindPageByParams(simple.DB(), params)
}

func (this *tagService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.Tag, paging *simple.Paging) {
	return repositories.TagRepository.FindPageByCnd(simple.DB(), cnd)
}

func (this *tagService) Create(t *model.Tag) error {
	return repositories.TagRepository.Create(simple.DB(), t)
}

func (this *tagService) Update(t *model.Tag) error {
	return repositories.TagRepository.Update(simple.DB(), t)
}

func (this *tagService) Updates(id int64, columns map[string]interface{}) error {
	return repositories.TagRepository.Updates(simple.DB(), id, columns)
}

func (this *tagService) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.TagRepository.UpdateColumn(simple.DB(), id, name, value)
}

func (this *tagService) Delete(id int64) {
	repositories.TagRepository.Delete(simple.DB(), id)
}

// 自动完成
func (this *tagService) Autocomplete(input string) []model.Tag {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return nil
	}
	list, _ := repositories.TagRepository.Find(simple.DB(), simple.NewSqlCnd().Where("status = ? and name like ?",
		model.TagStatusOk, "%"+input+"%").Limit(6))
	return list
}

func (this *tagService) GetOrCreate(name string) (*model.Tag, error) {
	return repositories.TagRepository.GetOrCreate(simple.DB(), name)
}

func (this *tagService) GetByName(name string) *model.Tag {
	return repositories.TagRepository.FindByName(name)
}

func (this *tagService) GetTags() []model.TagResponse {
	list, err := repositories.TagRepository.Find(simple.DB(), simple.NewSqlCnd().Where("status = ?", model.TagStatusOk))
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
	return repositories.TagRepository.GetTagInIds(tagIds)
}
