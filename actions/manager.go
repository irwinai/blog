package actions

import (
	"fmt"

	"github.com/lunny/xweb"
)

type ManagerAction struct {
	BaseAction

	home  xweb.Mapper `xweb:"/"`
	blogs xweb.Mapper `xweb:"/blog/list"`

	User User
}

func (m *ManagerAction) Home() {
	m.Render("manager/home.html")
}

func (m *ManagerAction) Blogs() {
	blogs := make([]*Blog, 0)
	err := m.Orm.Where("status=?", 1).Find(&blogs)
	if err != nil {
		fmt.Println(err)
		return
	}
	m.Render("blog-list.html", &xweb.T{
		"blogs": blogs,
	})
}
