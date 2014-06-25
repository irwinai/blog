package actions

import "github.com/lunny/xweb"

type ManagerAction struct {
	BaseAction

	home xweb.Mapper `xweb:"/"`

	User User
}

func (c *ManagerAction) Home() {
	c.Render("manager/home.html")
}
