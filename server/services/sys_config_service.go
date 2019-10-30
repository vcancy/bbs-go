
package services

import (
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/simple"
)

var SysConfigService = &sysConfigService {}

type sysConfigService struct {
}

func (this *sysConfigService) Get(id int64) *model.SysConfig {
	ret := &model.SysConfig{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *sysConfigService) Take(where ...interface{}) *model.SysConfig {
	ret := &model.SysConfig{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *sysConfigService) QueryCnd(cnd *simple.SqlCnd) (list []model.SysConfig, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *sysConfigService) Query(params *simple.QueryParams) (list []model.SysConfig, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.SysConfig{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *sysConfigService) Create(t *model.SysConfig) (*model.SysConfig, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *sysConfigService) Update(t *model.SysConfig) error {
	return simple.DB().Save(t).Error
}

func (this *sysConfigService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.SysConfig{}).Where("id = ?", id).Updates(columns).Error
}

func (this *sysConfigService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.SysConfig{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *sysConfigService) Delete(id int64) error {
	return simple.DB().Delete(&model.SysConfig{}, "id = ?", id).Error
}

