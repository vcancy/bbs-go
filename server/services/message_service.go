package services

import (
	"github.com/mlogclub/simple"

	"github.com/mlogclub/bbs-go/common"
	"github.com/mlogclub/bbs-go/common/email"
	"github.com/mlogclub/bbs-go/common/urls"
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/bbs-go/repositories"
	"github.com/mlogclub/bbs-go/services/cache"
)

var MessageService = newMessageService()

func newMessageService() *messageService {
	return &messageService{}
}

type messageService struct {
}

func (this *messageService) Get(id int64) *model.Message {
	return repositories.MessageRepository.Get(simple.DB(), id)
}

func (this *messageService) Take(where ...interface{}) *model.Message {
	return repositories.MessageRepository.Take(simple.DB(), where...)
}

func (this *messageService) Find(cnd *simple.SqlCnd) (list []model.Message, err error) {
	return repositories.MessageRepository.Find(simple.DB(), cnd)
}

func (this *messageService) FindPageByParams(params *simple.QueryParams) (list []model.Message, paging *simple.Paging) {
	return repositories.MessageRepository.FindPageByParams(simple.DB(), params)
}

func (this *messageService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.Message, paging *simple.Paging) {
	return repositories.MessageRepository.FindPageByCnd(simple.DB(), cnd)
}

func (this *messageService) Create(t *model.Message) error {
	return repositories.MessageRepository.Create(simple.DB(), t)
}

func (this *messageService) Update(t *model.Message) error {
	return repositories.MessageRepository.Update(simple.DB(), t)
}

func (this *messageService) Updates(id int64, columns map[string]interface{}) error {
	return repositories.MessageRepository.Updates(simple.DB(), id, columns)
}

func (this *messageService) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.MessageRepository.UpdateColumn(simple.DB(), id, name, value)
}

func (this *messageService) Delete(id int64) {
	repositories.MessageRepository.Delete(simple.DB(), id)
}

func (this *messageService) GetUnReadCount(userId int64) (count int64) {
	simple.DB().Where("user_id = ? and status = ?", userId, model.MsgStatusUnread).Model(&model.Message{}).Count(&count)
	return
}

// 读消息
func (this *messageService) Read(id int64) *model.Message {
	msg := this.Get(id)
	if msg != nil && msg.Status == model.MsgStatusUnread {
		_ = this.UpdateColumn(id, "status", model.MsgStatusReaded) // 标记为已读
	}
	return msg
}

// 将所有消息标记为已读
func (this *messageService) MarkReadAll(userId int64) {
	simple.DB().Exec("update t_message set status = ? where user_id = ? and status = ?", model.MsgStatusReaded,
		userId, model.MsgStatusUnread)
}

// 发送消息
// fromId: 消息发送人
// toId: 消息接收人
func (this *messageService) Send(fromId, toId int64, content, quoteContent string, msgType int, extraData map[string]interface{}) {
	extraDataStr, _ := simple.FormatJson(extraData)
	message := &model.Message{
		FromId:       fromId,
		UserId:       toId,
		Content:      content,
		QuoteContent: quoteContent,
		Type:         msgType,
		ExtraData:    extraDataStr,
		Status:       model.MsgStatusUnread,
		CreateTime:   simple.NowTimestamp(),
	}
	err := this.Create(message)
	if err == nil {
		go func() {
			this.sendEmailNotice(message)
		}()
	}
}

func (this *messageService) sendEmailNotice(message *model.Message) {
	user := cache.UserCache.Get(message.UserId)
	if user == nil || len(user.Email.String) == 0 {
		return
	}
	email.SendTemplateEmail(user.Email.String, "M-LOG新消息提醒", "M-LOG新消息提醒", message.Content,
		message.QuoteContent, urls.AbsUrl("/user/messages"))
}

func (this *messageService) SendCommentMsg(comment *model.Comment) {
	commentUser := repositories.UserRepository.Get(simple.DB(), comment.UserId)
	commentSummary := common.GetMarkdownSummary(comment.Content)
	// 引用消息
	if comment.QuoteId > 0 {
		quote := repositories.CommentRepository.Get(simple.DB(), comment.QuoteId)
		if quote != nil && quote.UserId != comment.UserId {
			msgContent := commentUser.Nickname + " 回复了你的评论：" + commentSummary
			quoteContent := common.GetMarkdownSummary(quote.Content)
			this.Send(comment.UserId, quote.UserId, msgContent, quoteContent, model.MsgTypeComment, map[string]interface{}{
				"entityType": comment.EntityType,
				"entityId":   comment.EntityId,
				"commentId":  comment.Id,
				"quoteId":    comment.QuoteId,
			})
		}
	}

	// 文章评论消息
	{
		var userId int64 = 0
		var msgContent = ""
		var msgQuoteContent = ""
		if comment.EntityType == model.EntityTypeArticle {
			article := repositories.ArticleRepository.Get(simple.DB(), comment.EntityId)
			if article != nil && article.UserId != comment.UserId {
				userId = article.UserId
				msgContent = commentUser.Nickname + " 回复了你的文章：" + commentSummary
				msgQuoteContent = "《" + article.Title + "》"
			}
		} else if comment.EntityType == model.EntityTypeTopic {
			topic := repositories.TopicRepository.Get(simple.DB(), comment.EntityId)
			if topic != nil && topic.UserId != comment.UserId {
				userId = topic.UserId
				msgContent = commentUser.Nickname + " 回复了你的主题：" + commentSummary
				msgQuoteContent = "《" + topic.Title + "》"
			}
		}
		if userId > 0 {
			this.Send(comment.UserId, userId, msgContent, msgQuoteContent, model.MsgTypeComment, map[string]interface{}{
				"entityType": comment.EntityType,
				"entityId":   comment.EntityId,
				"commentId":  comment.Id,
				"quoteId":    comment.QuoteId,
			})
		}
	}
}
