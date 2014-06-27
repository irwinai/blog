package actions

import (
	"fmt"

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
		if c.Status == "on" {
			c.Blog.Status = 1
		} else {
			c.Blog.Status = 0
		}
		_, err := c.Orm.Insert(c.Blog)
		if err != nil {
			fmt.Println(err)
			return
		}
		if c.Cat != nil {
			list := make([]*BlogCategory, 10)
			for _, value := range c.Cat {
				bc := new(BlogCategory)
				bc.BlogId = c.Blog.Id
				bc.CategoryId = value
				list = append(list, bc)
			}
			c.Orm.Insert(list)
		}
		fmt.Print("status=", c.Status)
		fmt.Println("cat = ", c.Cat)
		fmt.Println("tags = ", c.Tags)
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
