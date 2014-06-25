package actions

import "github.com/lunny/xweb"

type BlogAction struct {
	BaseAction

	list    xweb.Mapper
	detail  xweb.Mapper
	archive xweb.Mapper
	add     xweb.Mapper

	Blog Blog
}

func (c *BlogAction) Add() {
	if c.Method() == "GET" {
		c.Render("manager/article_write.html")
	} else if c.Method() == "POST" {
		//c.Orm.Insert(Blog)
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
