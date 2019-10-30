package main

import (
	"github.com/mlogclub/simple"

	"github.com/mlogclub/bbs-go/model"
)

func main() {
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.User{}))
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.UserToken{}))
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.Category{}))
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.Tag{}))
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.Article{}))
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.ArticleTag{}))
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.Comment{}))
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.Favorite{}))
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.Topic{}))
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.TopicTag{}))
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.TopicLike{}))
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.Message{}))
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.SysConfig{}))
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.Project{}))
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.Subject{}))
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.SubjectContent{}))
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.Link{}))
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.CollectRule{}))
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.CollectArticle{}))
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.ThirdAccount{}))
}
