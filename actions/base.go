package actions

import (
	"fmt"

	. "github.com/lunny/config"
	"github.com/lunny/xorm"
	"github.com/lunny/xweb"
)

type BaseAction struct {
	*xweb.Action
	Orm *xorm.Engine
	cfg *Config
}

func (c *BaseAction) Init() {
	c.Orm = c.App.GetConfig("Orm").(*xorm.Engine)
	c.cfg = c.App.GetConfig("cfg").(*Config)
	//初始化查询菜单栏category列表
	categories := make([]Category, 0)
	err := c.Orm.Where("status=?", 1).Find(&categories)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.AddTmplVars(&xweb.T{
		"categories": categories,
	})
}
