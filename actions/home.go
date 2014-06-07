package actions

import "github.com/lunny/xweb"

type HomeAction struct {
	BaseAction

	login xweb.Mapper

	User User
}

func (c *HomeAction) Login() error {
	if c.Method() == "GET" {
		return c.Render("login.html")
	} else if c.Method() == "POST" {
		if len(c.User.Username) <= 0 {
			return c.Go("login?message=登录名不能为空")
		}
		has, err := c.Orm.Get(&c.User)
		if err == nil {
			if has {
				session := c.Session()
				session.Set(USER_ID, c.User.Id)
				session.Set(USERNAME, c.User.Username)
				return c.Go("list", &BlogAction{})
			}
			return c.Go("login?message=账号或密码错误")
		}
		return err
	}
	return xweb.NotSupported()
}
