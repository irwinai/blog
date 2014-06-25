package actions

import (
	"encoding/json"

	"github.com/lunny/xweb"
)

type HomeAction struct {
	BaseAction

	login  xweb.Mapper
	logout xweb.Mapper
	about  xweb.Mapper
	test   xweb.Mapper

	User User
}

func (c *HomeAction) Test() {
	c.Render("test.html")
}

func (c *HomeAction) About() {
	c.Render("about.html")
}

func (c *HomeAction) Logout() {
	c.Session().Invalidate(c.ResponseWriter)
	c.Go("login")
}

func (c *HomeAction) Login() {
	if c.Method() == "GET" {
		c.Render("login.html")
	} else if c.Method() == "POST" {
		user := c.GetString("user")
		json.Unmarshal([]byte(user), &c.User)
		m := make(map[string]string)
		if len(c.User.Username) <= 0 || len(c.User.Password) <= 0 {
			m[AJAX_STATUS] = "0"
			m[AJAX_MESSAGE] = "用户名或密码不能为空"
		} else {
			has, err := c.Orm.Get(&c.User)
			if err == nil {
				if has {
					session := c.Session()
					session.Set(USER_ID, c.User.Id)
					session.Set(USERNAME, c.User.Username)
					m[AJAX_STATUS] = "1"
					m[AJAX_MESSAGE] = "登入成功"
				} else {
					m[AJAX_STATUS] = "0"
					m[AJAX_MESSAGE] = "用户名或密码错误"
				}
			} else {
				m[AJAX_STATUS] = "0"
				m[AJAX_MESSAGE] = "服务器异常，请稍候重试"
			}
		}
		c.ServeJson(m)
	}
}
