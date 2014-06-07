package actions

import (
	"github.com/lunny/xorm"
	"github.com/lunny/xweb"
)

type BaseAction struct {
	*xweb.Action
	Orm *xorm.Engine
}

func (c *BaseAction) Init() {
	c.Orm = c.App.GetConfig("Orm").(*xorm.Engine)
}
