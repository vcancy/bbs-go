package services

import (
	"github.com/kataras/iris/context"

	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/bbs-go/services/cache"

	"time"

	"github.com/mlogclub/simple"
)

var UserTokenService = &userTokenService{}

type userTokenService struct {
}

func (this *userTokenService) Get(id int64) *model.UserToken {
	ret := &model.UserToken{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *userTokenService) Take(where ...interface{}) *model.UserToken {
	ret := &model.UserToken{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *userTokenService) QueryCnd(cnd *simple.SqlCnd) (list []model.UserToken, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *userTokenService) Query(params *simple.QueryParams) (list []model.UserToken, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.UserToken{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *userTokenService) Create(t *model.UserToken) (*model.UserToken, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *userTokenService) Update(t *model.UserToken) error {
	return simple.DB().Save(t).Error
}

func (this *userTokenService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.UserToken{}).Where("id = ?", id).Updates(columns).Error
}

func (this *userTokenService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.UserToken{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *userTokenService) Delete(id int64) error {
	return simple.DB().Delete(&model.UserToken{}, "id = ?", id).Error
}

// 根据Token查询
func (this *userTokenService) GetByToken(token string) *model.UserToken {
	if len(token) == 0 {
		return nil
	}
	return this.Take("token = ?", token)
}

// 获取当前登录用户
func (this *userTokenService) GetCurrent(ctx context.Context) *model.User {
	token := this.GetUserToken(ctx)
	userToken := cache.UserTokenCache.Get(token)
	// 没找到授权
	if userToken == nil || userToken.Status == model.UserTokenStatusDisabled {
		return nil
	}
	// 授权过期
	if userToken.ExpiredAt <= simple.NowTimestamp() {
		return nil
	}
	return cache.UserCache.Get(userToken.UserId)
}

// 退出登录
func (this *userTokenService) Signout(ctx context.Context) error {
	token := this.GetUserToken(ctx)
	userToken := this.GetByToken(token)
	if userToken == nil {
		return nil
	}
	return this.UpdateColumn(userToken.Id, "status", model.UserTokenStatusDisabled)
}

// 从请求体中获取UserToken
func (this *userTokenService) GetUserToken(ctx context.Context) string {
	userToken := ctx.FormValue("userToken")
	if len(userToken) > 0 {
		return userToken
	}
	return ctx.GetHeader("X-User-Token")
}

// 生成
func (this *userTokenService) Generate(userId int64) (string, error) {
	expiredAt := time.Now().Add(time.Hour * 24 * 7) // 7天后过期
	userToken, err := this.Create(&model.UserToken{
		Token:      simple.Uuid(),
		UserId:     userId,
		ExpiredAt:  simple.Timestamp(expiredAt),
		Status:     model.UserTokenStatusOk,
		CreateTime: simple.NowTimestamp(),
	})
	if err != nil {
		return "", err
	}
	return userToken.Token, nil
}

// 禁用
func (this *userTokenService) Disable(token string) error {
	t := this.GetByToken(token)
	if t == nil {
		return nil
	}
	err := this.UpdateColumn(t.Id, "status", model.UserTokenStatusDisabled)
	if err != nil {
		cache.UserTokenCache.Invalidate(token)
	}
	return err
}
