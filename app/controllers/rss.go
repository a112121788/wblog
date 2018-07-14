package controllers

import (
	"fmt"

	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/feeds"
	"github.com/a112121788/wblog/app/helpers"
	"github.com/a112121788/wblog/app/models"
	"github.com/a112121788/wblog/config"
)

func RssGet(c *gin.Context) {
	now := helpers.GetCurrentTime()
	domain := config.GetConfiguration().Domain
	feed := &feeds.Feed{
		Title:       "Wblog",
		Link:        &feeds.Link{Href: domain},
		Description: "Wblog,talk about golang,java and so on.",
		Author:      &feeds.Author{Name: "a112121788", Email: "a112121788love@163.com"},
		Created:     now,
	}

	feed.Items = make([]*feeds.Item, 0)
	posts, err := models.ListPublishedPost("")
	if err == nil {
		for _, post := range posts {
			item := &feeds.Item{
				Id:          fmt.Sprintf("%s/post/%d", domain, post.ID),
				Title:       post.Title,
				Link:        &feeds.Link{Href: fmt.Sprintf("%s/post/%d", domain, post.ID)},
				Description: string(post.Excerpt()),
				Created:     now,
			}
			feed.Items = append(feed.Items, item)
		}
	}
	rss, err := feed.ToRss()
	if err == nil {
		c.Writer.WriteString(rss)
	} else {
		seelog.Error(err)
	}
}
