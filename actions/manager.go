package actions

import (
	"fmt"
	"strings"
	"time"

	"github.com/lunny/xweb"
)

type ManagerAction struct {
	BaseAction

	home  xweb.Mapper `xweb:"/"`
	blogs xweb.Mapper `xweb:"/blog/list"`
	add   xweb.Mapper `xweb:"/blog/add"`

	User    User
	Blog    Blog
	Status  string
	Cat     []int
	Tags    string
	PicName string
}

func (m *ManagerAction) Home() {
	m.Render("manager/home.html")
}

func (m *ManagerAction) Blogs() {
	blogs := make([]Blog, 0)
	err := m.Orm.Where("status=?", 1).Find(&blogs)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(blogs) > 0 {
		result := make([]map[string]interface{}, len(blogs))
		for key, blog := range blogs {
			mp := make(map[string]interface{})
			mp["id"] = blog.Id
			mp["title"] = blog.Title
			mp["time"] = blog.CreateTime
			mp["status"] = blog.Status
			mp["commentCount"] = blog.CommentCount
			mp["views"] = blog.Views
			mp["like"] = blog.Like
			mp["createTime"] = blog.CreateTime
			//根据ID查询用户
			user := new(User)
			_, err = m.Orm.Id(blog.UserId).Get(user)
			if err != nil {
				return
			}
			if user != nil {
				mp["username"] = user.Username
			}
			//查询博客分类关联
			bs := make([]BlogCategory, 0)
			err = m.Orm.Where("blog_id=?", blog.Id).Find(&bs)
			if err != nil {
				return
			}
			if len(bs) > 0 {
				categories := make([]string, len(bs))
				for key2, value := range bs {
					category := new(Category)
					_, err = m.Orm.Id(value.CategoryId).Get(category)
					if err != nil {
						return
					}
					categories[key2] = category.Name
				}
				mp["categories"] = categories
			}
			//查询标签
			bt := make([]BlogTag, 0)
			err = m.Orm.Where("blog_id=?", blog.Id).Find(&bt)
			if err != nil {
				return
			}
			if len(bt) > 0 {
				tags := make([]string, len(bt))
				for key3, value := range bt {
					tag := new(Tag)
					m.Orm.Id(value.TagId).Get(tag)
					tags[key3] = tag.Name
				}
				mp["tags"] = tags
			}
			result[key] = mp
		}
		m.Render("manager/blog-list.html", &xweb.T{
			"result": result,
		})
	}
}

func (c *ManagerAction) Add() error {
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
			return err
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
					return err
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
						return err
					}
					bt.TagId = newTag.Id
				}
				tagList[key] = bt
			}
			session.Insert(&tagList)
		}

		//保存图片关联
		pic := new(Picture)
		_, err = c.Orm.Where("name=?", c.PicName).Get(pic)
		if err != nil {
			fmt.Println(err)
			return err
		}
		pic.BlogId = c.Blog.Id
		c.Orm.Update(pic)
		c.Go("blogs")
	}
	return xweb.NotSupported()
}
