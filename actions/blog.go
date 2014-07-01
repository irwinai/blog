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
	paging  xweb.Mapper

	Blog   Blog
	Status string
	Cat    []int
	Tags   string
}

func (c *BlogAction) Paging() {
	key := c.GetString("type")
	id, _ := c.GetInt("id")
	sql := ""
	if key == "pre" {
		sql = "select max(id) mid from blog where id<? and status=1"
	} else if key == "next" {
		sql = "select min(id) mid from blog where id>? and status=1"
	}
	fmt.Println(id)
	fmt.Println(sql)
	result, err := c.Orm.Query(sql, id)
	if err != nil {
		return
	}
	fmt.Println(result)
	if result != nil && len(result) > 0 && len(result[0]) > 0 {
		mid := (string)(result[0]["mid"][0])
		c.Go(fmt.Sprintf("detail?id=%v", mid))
	} else {
		c.Go(fmt.Sprintf("detail?id=%v", id))
	}
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

func (c *BlogAction) List() {
	blogs := make([]Blog, 0)
	err := c.Orm.Where("status=?", 1).Find(&blogs)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Render("list.html", &xweb.T{
		"blogs": blogs,
	})
}

func (c *BlogAction) Detail() {
	id, _ := c.GetInt("id")
	blog := new(Blog)
	_, err := c.Orm.Id(id).Get(blog)
	if err != nil {
		return
	}
	//根据ID查询用户
	user := new(User)
	_, err = c.Orm.Id(blog.UserId).Get(user)
	if err != nil {
		return
	}
	//查询标签
	bt := make([]BlogTag, 0)
	err = c.Orm.Where("blog_id=?", blog.Id).Find(&bt)
	if err != nil {
		return
	}
	tags := make([]*Tag, len(bt))
	if len(bt) > 0 {
		for key, value := range bt {
			tag := new(Tag)
			c.Orm.Id(value.TagId).Get(tag)
			fmt.Println(tag.Name)
			tags[key] = tag
		}
	}
	fmt.Println(user)
	fmt.Println(tags)
	c.Render("blog.html", &xweb.T{
		"blog": blog,
		"user": user,
		"tags": tags,
	})
}

func (c *BlogAction) Archive() error {
	return c.Render("archive.html")
}
