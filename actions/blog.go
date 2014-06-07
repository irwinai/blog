package actions

import "github.com/lunny/xweb"

type BlogAction struct {
	BaseAction

	list xweb.Mapper
}

func (c *BlogAction) List() error {
	return c.Render("blog.html")
}
