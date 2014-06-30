package actions

import (
	"fmt"
	"strings"
	"time"

	"github.com/lunny/xweb"
)

type BlogAction struct {
	BaseAction

	list    xweb.Mapper
	detail  xweb.Mapper
	archive xweb.Mapper
	add     xweb.Mapper

	Blog   Blog
	Status string
	Cat    []int
	Tags   string
}

func (c *BlogAction) Add() {
	if c.Method() == "GET" {
		c.Render("manager/add-blog.html")
	} else if c.Method() == "POST" {
		session := c.Orm.NewSession()
		defer session.Close()
		defer session.Commit()
		session.Begin()
		if c.Status == "on" {
			c.Blog.Status = 1
		} else {
			c.Blog.Status = 0
		}
		c.Blog.CreateTime = time.Now()
		c.Blog.UpdateTime = time.Now()
		//保存博客
		_, err := session.Insert(&c.Blog)
		if err != nil {
			fmt.Println(err)
			session.Rollback()
			return
		}
		//保存博客分类关联
		if c.Cat != nil {
			list := make([]*BlogCategory, len(c.Cat))
			for key, value := range c.Cat {
				bc := new(BlogCategory)
				bc.BlogId = c.Blog.Id
				bc.CategoryId = value
				list[key] = bc
			}
			session.Insert(&list)
		}
		//保存博客标签关联
		if c.Tags != "" {
			tags := strings.Split(c.Tags, ",")
			tagList := make([]*BlogTag, len(tags))
			for key, value := range tags {
				tag := new(Tag)
				has, err := session.Where("name=? and status=?", value, 1).Get(tag)
				if err != nil {
					fmt.Println(err)
					return
				}
				//如果该标签已经存在
				bt := new(BlogTag)
				bt.BlogId = c.Blog.Id
				if has {
					bt.TagId = tag.Id
				} else {
					newTag := new(Tag)
					newTag.Name = value
					newTag.Status = 1
					_, err := session.Insert(newTag)
					if err != nil {
						fmt.Println(err)
						session.Rollback()
						return
					}
					bt.TagId = newTag.Id
				}
				tagList[key] = bt
			}
			session.Insert(&tagList)
		}
		fmt.Println("插入成功======================")
	}
}

func (c *BlogAction) List() error {
	return c.Render("list.html")
}

func (c *BlogAction) Detail() error {
	return c.Render("blog.html")
}

func (c *BlogAction) Archive() error {
	return c.Render("archive.html")
}
