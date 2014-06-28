package actions

import "github.com/lunny/xweb"

type CategoryAction struct {
	BaseAction

	list xweb.Mapper
}

func (c *CategoryAction) List() {

}
