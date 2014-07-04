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
	paging  xweb.Mapper
	like    xweb.Mapper
}

//赞功能
func (c *BlogAction) Like() {

}

//博客详情页左右翻页功能
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
	if result != nil && len(result) > 0 && len(result[0]) > 0 {
		mid := (string)(result[0]["mid"][0])
		c.Go(fmt.Sprintf("detail?id=%v", mid))
	} else {
		c.Go(fmt.Sprintf("detail?id=%v", id))
	}
}

func (c *BlogAction) List() {
	blogs := make([]Blog, 0)
	err := c.Orm.Where("status=?", 1).Desc("create_time").Find(&blogs)
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
