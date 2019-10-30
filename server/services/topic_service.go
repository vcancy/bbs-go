package services

import (
	"math"
	"path"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/mlogclub/bbs-go/common"
	"github.com/mlogclub/bbs-go/common/config"
	"github.com/mlogclub/bbs-go/common/urls"
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/bbs-go/services/cache"

	"github.com/gorilla/feeds"
	"github.com/mlogclub/simple"
	"github.com/sirupsen/logrus"
)

type ScanTopicCallback func(topics []model.Topic) bool

var TopicService = &topicService{}

type topicService struct {
}

func (this *topicService) Get(id int64) *model.Topic {
	ret := &model.Topic{}
	if err := simple.DB().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *topicService) Take(where ...interface{}) *model.Topic {
	ret := &model.Topic{}
	if err := simple.DB().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *topicService) QueryCnd(cnd *simple.SqlCnd) (list []model.Topic, err error) {
	err = cnd.Exec(simple.DB()).Find(&list).Error
	return
}

func (this *topicService) Query(params *simple.QueryParams) (list []model.Topic, paging *simple.Paging) {
	params.StartQuery(simple.DB()).Find(&list)
	params.StartCount(simple.DB()).Model(&model.Topic{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *topicService) Create(t *model.Topic) (*model.Topic, error) {
	if err := simple.DB().Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (this *topicService) Update(t *model.Topic) error {
	return simple.DB().Save(t).Error
}

func (this *topicService) Updates(id int64, columns map[string]interface{}) error {
	return simple.DB().Model(&model.Topic{}).Where("id = ?", id).Updates(columns).Error
}

func (this *topicService) UpdateColumn(id int64, name string, value interface{}) error {
	return simple.DB().Model(&model.Topic{}).Where("id = ?", id).UpdateColumn(name, value).Error
}

func (this *topicService) Delete(id int64) error {
	err := this.UpdateColumn(id, "status", model.TopicStatusDeleted)
	if err == nil {
		// 删掉专栏文章
		SubjectContentService.DeleteByEntity(model.EntityTypeTopic, id)
		// 删掉标签文章
		TopicTagService.DeleteByTopicId(id)
	}
	return err
}

// 扫描
func (this *topicService) Scan(cb ScanTopicCallback) {
	var cursor int64
	for {
		list, err := this.QueryCnd(simple.NewSqlCnd("id > ?", cursor).Order("id asc").Size(100))
		if err != nil {
			break
		}
		if list == nil || len(list) == 0 {
			break
		}
		cursor = list[len(list)-1].Id
		if !cb(list) {
			break
		}
	}
}

// 倒序扫描
func (this *topicService) ScanDesc(cb ScanTopicCallback) {
	var cursor int64 = math.MaxInt64
	for {
		list, err := this.QueryCnd(simple.NewSqlCnd("id < ?", cursor).Order("id desc").Size(100))
		if err != nil {
			break
		}
		if list == nil || len(list) == 0 {
			break
		}
		cursor = list[len(list)-1].Id
		if !cb(list) {
			break
		}
	}
}

// 发表
func (this *topicService) Publish(userId int64, tags []string, title, content string, extra map[string]interface{}) (*model.Topic, *simple.CodeError) {
	if len(title) == 0 {
		return nil, simple.NewErrorMsg("标题不能为空")
	}

	if simple.RuneLen(title) > 128 {
		return nil, simple.NewErrorMsg("标题长度不能超过128")
	}

	now := simple.NowTimestamp()
	topic := &model.Topic{
		UserId:          userId,
		Title:           title,
		Content:         content,
		Status:          model.TopicStatusOk,
		LastCommentTime: now,
		CreateTime:      now,
	}
	if len(extra) > 0 {
		topic.ExtraData, _ = simple.FormatJson(extra)
	}

	err := simple.Tx(simple.DB(), func(tx *gorm.DB) error {
		if err := simple.DB().Create(topic).Error; err != nil {
			return err
		}
		tagIds := TagService.GetOrCreates(tx, tags)
		TopicTagService.CreateTopicTags(tx, topic.Id, tagIds)
		return nil
	})
	if err == nil {
		common.BaiduUrlPush([]string{urls.TopicUrl(topic.Id)})
	}
	return topic, simple.FromError(err)
}

// 更新
func (this *topicService) Edit(topicId int64, tags []string, title, content string) *simple.CodeError {
	if len(title) == 0 {
		return simple.NewErrorMsg("标题不能为空")
	}

	if simple.RuneLen(title) > 128 {
		return simple.NewErrorMsg("标题长度不能超过128")
	}

	err := simple.Tx(simple.DB(), func(tx *gorm.DB) error {
		tagIds := TagService.GetOrCreates(tx, tags)
		if err := tx.Model(&model.Topic{}).Where("id = ?", topicId).Updates(map[string]interface{}{
			"title":   title,
			"content": content,
		}).Error; err != nil {
			return err
		}
		TopicTagService.DeleteTopicTags(tx, topicId)         // 先删掉所有的标签
		TopicTagService.CreateTopicTags(tx, topicId, tagIds) // 然后重新添加标签
		return nil
	})
	return simple.FromError(err)
}

// 主题标签
func (this *topicService) GetTopicTags(topicId int64) []model.Tag {
	topicTags, err := TopicTagService.QueryCnd(simple.NewSqlCnd("topic_id = ?", topicId))
	if err != nil {
		return nil
	}
	var tagIds []int64
	for _, topicTag := range topicTags {
		tagIds = append(tagIds, topicTag.TagId)
	}
	return cache.TagCache.GetList(tagIds)
}

// 指定标签下的主题列表
func (this *topicService) GetTagTopics(tagId int64, page int) (topics []model.Topic, paging *simple.Paging) {
	topicTags, paging := TopicTagService.Query(simple.NewQueryParams(nil).
		Eq("tag_id", tagId).
		Eq("status", model.ArticleTagStatusOk).
		Page(page, 20).Desc("id"))
	if len(topicTags) > 0 {
		var topicIds []int64
		for _, topicTag := range topicTags {
			topicIds = append(topicIds, topicTag.TopicId)
		}
		topics = this.GetTopicInIds(topicIds)
	}
	return
}

// 根据编号批量获取主题
func (this *topicService) GetTopicInIds(topicIds []int64) []model.Topic {
	if len(topicIds) == 0 {
		return nil
	}
	var topics []model.Topic
	simple.DB().Where("id in (?)", topicIds).Find(&topics)
	return topics
}

// 浏览数+1
func (this *topicService) IncrViewCount(topicId int64) {
	simple.DB().Exec("update t_topic set view_count = view_count + 1 where id = ?", topicId)
}

// 当帖子被评论的时候，更新最后回复时间、回复数量+1
func (this *topicService) OnComment(topicId, lastCommentTime int64) {
	simple.DB().Exec("update t_topic set last_comment_time = ?, comment_count = comment_count + 1 where id = ?",
		lastCommentTime, topicId)
}

// rss
func (this *topicService) GenerateRss() {
	topics, err := this.QueryCnd(simple.NewSqlCnd("status = ?", model.TopicStatusOk).Order("id desc").Size(1000))
	if err != nil {
		logrus.Error(err)
		return
	}

	var items []*feeds.Item

	for _, topic := range topics {
		topicUrl := urls.TopicUrl(topic.Id)
		user := cache.UserCache.Get(topic.UserId)
		if user == nil {
			continue
		}
		item := &feeds.Item{
			Title:       topic.Title,
			Link:        &feeds.Link{Href: topicUrl},
			Description: common.GetMarkdownSummary(topic.Content),
			Author:      &feeds.Author{Name: user.Avatar, Email: user.Email.String},
			Created:     simple.TimeFromTimestamp(topic.CreateTime),
		}
		items = append(items, item)
	}
	siteTitle := cache.SysConfigCache.GetValue(model.SysConfigSiteTitle)
	siteDescription := cache.SysConfigCache.GetValue(model.SysConfigSiteDescription)
	feed := &feeds.Feed{
		Title:       siteTitle,
		Link:        &feeds.Link{Href: config.Conf.BaseUrl},
		Description: siteDescription,
		Author:      &feeds.Author{Name: siteTitle},
		Created:     time.Now(),
		Items:       items,
	}
	atom, err := feed.ToAtom()
	if err != nil {
		logrus.Error(err)
	} else {
		_ = simple.WriteString(path.Join(config.Conf.StaticPath, "topic_atom.xml"), atom, false)
	}

	rss, err := feed.ToRss()
	if err != nil {
		logrus.Error(err)
	} else {
		_ = simple.WriteString(path.Join(config.Conf.StaticPath, "topic_rss.xml"), rss, false)
	}
}
