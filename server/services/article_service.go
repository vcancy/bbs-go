package services

import (
	"errors"
	"math"
	"path"
	"strings"
	"time"

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/gorilla/feeds"
	"github.com/jinzhu/gorm"
	"github.com/mlogclub/simple"
	"github.com/sirupsen/logrus"

	"github.com/mlogclub/bbs-go/common"
	"github.com/mlogclub/bbs-go/common/config"
	"github.com/mlogclub/bbs-go/common/urls"
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/bbs-go/services/cache"
)

type ScanArticleCallback func(articles []model.Article) bool

var ArticleService = &articleService{}

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
	err := this.UpdateColumn(id, "status", model.ArticleStatusDeleted)
	if err == nil {
		// 删掉专栏文章
		SubjectContentService.DeleteByEntity(model.EntityTypeArticle, id)
		// 删掉标签文章
		ArticleTagService.DeleteByArticleId(id)
	}
	return err
}

// 根据文章编号批量获取文章
func (this *articleService) GetArticleInIds(articleIds []int64) []model.Article {
	if len(articleIds) == 0 {
		return nil
	}
	var articles []model.Article
	simple.DB().Where("id in (?)", articleIds).Find(&articles)
	return articles
}

// 获取文章对应的标签
func (this *articleService) GetArticleTags(articleId int64) []model.Tag {
	articleTags, err := ArticleTagService.QueryCnd(simple.NewSqlCnd("article_id = ?", articleId))
	if err != nil {
		return nil
	}
	var tagIds []int64
	for _, articleTag := range articleTags {
		tagIds = append(tagIds, articleTag.TagId)
	}
	return cache.TagCache.GetList(tagIds)
}

// 标签文章列表
func (this *articleService) GetTagArticles(tagId int64, page int) (articles []model.Article, paging *simple.Paging) {
	articleTags, paging := ArticleTagService.Query(simple.NewQueryParams(nil).
		Eq("tag_id", tagId).
		Eq("status", model.ArticleTagStatusOk).
		Page(page, 20).Desc("id"))
	if len(articleTags) > 0 {
		var articleIds []int64
		for _, articleTag := range articleTags {
			articleIds = append(articleIds, articleTag.ArticleId)
		}
		articles = this.GetArticleInIds(articleIds)
	}
	return
}

// 发布文章
func (this *articleService) Publish(userId int64, title, summary, content, contentType string, categoryId int64,
	tags []string, sourceUrl string, share bool) (article *model.Article, err error) {

	title = strings.TrimSpace(title)
	summary = strings.TrimSpace(summary)
	content = strings.TrimSpace(content)

	if len(title) == 0 {
		return nil, errors.New("标题不能为空")
	}
	if share { // 如果是分享的内容，必须有Summary和SourceUrl
		if len(summary) == 0 {
			return nil, errors.New("分享内容摘要不能为空")
		}
		if len(sourceUrl) == 0 {
			return nil, errors.New("分享内容原文链接不能为空")
		}
	} else {
		if len(content) == 0 {
			return nil, errors.New("内容不能为空")
		}
	}
	article = &model.Article{
		UserId:      userId,
		Title:       title,
		Summary:     summary,
		Content:     content,
		ContentType: contentType,
		CategoryId:  categoryId,
		Status:      model.ArticleStatusPublished,
		Share:       share,
		SourceUrl:   sourceUrl,
		CreateTime:  simple.NowTimestamp(),
		UpdateTime:  simple.NowTimestamp(),
	}

	err = simple.Tx(simple.DB(), func(tx *gorm.DB) error {
		if err := tx.Create(article).Error; err != nil {
			return err
		}
		tagIds := TagService.GetOrCreates(tx, tags)
		ArticleTagService.CreateArticleTags(tx, article.Id, tagIds)
		return nil
	})

	if err == nil {
		common.BaiduUrlPush([]string{urls.ArticleUrl(article.Id)})
		SubjectContentService.AnalyzeArticle(article)
	}
	return
}

// 修改文章
func (this *articleService) Edit(articleId int64, tags []string, title, content string) *simple.CodeError {
	if len(title) == 0 {
		return simple.NewErrorMsg("请输入标题")
	}
	if len(content) == 0 {
		return simple.NewErrorMsg("请填写文章内容")
	}

	err := simple.Tx(simple.DB(), func(tx *gorm.DB) error {
		if err := tx.Model(&model.Article{}).Where("id = ?", articleId).Updates(map[string]interface{}{
			"title":   title,
			"content": content,
		}).Error; err != nil {
			return err
		}
		tagIds := TagService.GetOrCreates(tx, tags)
		ArticleTagService.DeleteArticleTags(tx, articleId)         // 先删掉所有的标签
		ArticleTagService.CreateArticleTags(tx, articleId, tagIds) // 然后重新添加标签
		return nil
	})
	cache.ArticleTagCache.Invalidate(articleId)
	return simple.FromError(err)
}

// 相关文章
func (this *articleService) GetRelatedArticles(articleId int64) []model.Article {
	tagIds := cache.ArticleTagCache.Get(articleId)
	if len(tagIds) == 0 {
		return nil
	}
	var articleTags []model.ArticleTag
	simple.DB().Where("tag_id in (?)", tagIds).Limit(30).Find(&articleTags)

	set := hashset.New()
	if len(articleTags) > 0 {
		for _, articleTag := range articleTags {
			set.Add(articleTag.ArticleId)
		}
	}

	var articleIds []int64
	for i, articleId := range set.Values() {
		if i < 10 {
			articleIds = append(articleIds, articleId.(int64))
		}
	}

	return this.GetArticleInIds(articleIds)
}

// 最新文章
func (this *articleService) GetUserNewestArticles(userId int64) []model.Article {
	articles, err := this.QueryCnd(simple.NewSqlCnd("user_id = ? and status = ?", userId, model.ArticleStatusPublished).Order("id desc").Size(10))
	if err != nil {
		return nil
	}
	return articles
}

// 扫描
func (this *articleService) Scan(cb ScanArticleCallback) {
	var cursor int64
	for {
		list, err := this.QueryCnd(simple.NewSqlCnd("id > ? ", cursor).Order("id asc").Size(100))
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

// 从新往旧扫描
func (this *articleService) ScanDesc(cb ScanArticleCallback) {
	var cursor int64 = math.MaxInt64
	for {
		list, err := this.QueryCnd(simple.NewSqlCnd("id < ? ", cursor).Order("id desc").Size(100))
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

// 扫描
func (this *articleService) ScanWithDate(dateFrom, dateTo int64, cb ScanArticleCallback) {
	var cursor int64
	for {
		list, err := this.QueryCnd(simple.NewSqlCnd("id > ? and status = ? and create_time >= ? and create_time < ?",
			cursor, model.ArticleStatusPublished, dateFrom, dateTo).Order("id asc").Size(300))
		if err != nil {
			break
		}
		if list == nil || len(list) == 0 {
			break
		}
		cursor = list[len(list)-1].Id
		cb(list)
	}
}

// rss
func (this *articleService) GenerateRss() {
	articles, err := this.QueryCnd(simple.NewSqlCnd("status = ?", model.ArticleStatusPublished).Order("id desc").Size(1000))
	if err != nil {
		logrus.Error(err)
		return
	}

	var items []*feeds.Item

	for _, article := range articles {
		articleUrl := urls.ArticleUrl(article.Id)
		user := cache.UserCache.Get(article.UserId)
		if user == nil {
			continue
		}
		description := ""
		if article.ContentType == model.ContentTypeMarkdown {
			description = common.GetMarkdownSummary(article.Content)
		} else {
			description = common.GetHtmlSummary(article.Content)
		}
		item := &feeds.Item{
			Title:       article.Title,
			Link:        &feeds.Link{Href: articleUrl},
			Description: description,
			Author:      &feeds.Author{Name: user.Avatar, Email: user.Email.String},
			Created:     simple.TimeFromTimestamp(article.CreateTime),
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
		_ = simple.WriteString(path.Join(config.Conf.StaticPath, "atom.xml"), atom, false)
	}

	rss, err := feed.ToRss()
	if err != nil {
		logrus.Error(err)
	} else {
		_ = simple.WriteString(path.Join(config.Conf.StaticPath, "rss.xml"), rss, false)
	}
}

// 生成码农日报内容
func (this *articleService) GetDailyContent(userIds []int64) string {
	if userIds == nil || len(userIds) == 0 {
		return ""
	}

	content := "\n"

	dateFromTemp := time.Now().Add(-time.Hour * 24)
	dateToTemp := time.Now()
	dateFrom := time.Date(dateFromTemp.Year(), dateFromTemp.Month(), dateFromTemp.Day(), 0, 0, 0, 0, dateFromTemp.Location())
	dateTo := time.Date(dateToTemp.Year(), dateToTemp.Month(), dateToTemp.Day(), 0, 0, 0, 0, dateToTemp.Location())

	this.ScanWithDate(simple.Timestamp(dateFrom), simple.Timestamp(dateTo), func(articles []model.Article) bool {
		for _, article := range articles {
			if common.IndexOf(userIds, article.UserId) != -1 {
				content += "## " + article.Title + "\n\n"
				if len(strings.TrimSpace(article.Summary)) > 0 {
					content += strings.TrimSpace(article.Summary) + "\n\n"
				}
				content += "[点击查看原文>>](" + urls.ArticleUrl(article.Id) + ")\n\n"
			}
		}
		return true
	})
	return content
}
