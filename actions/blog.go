package actions

import "github.com/lunny/xweb"

type BlogAction struct {
	BaseAction

	list   xweb.Mapper
	detail xweb.Mapper
}

func (c *BlogAction) List() error {
	return c.Render("list.html")
}

func (c *BlogAction) Detail() error {
	return c.Render("blog.html")
}
