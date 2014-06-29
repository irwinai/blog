package actions

import (
	"fmt"

	"github.com/lunny/xweb"
)

type UploadAction struct {
	BaseAction

	home xweb.Mapper `xweb:"/"`
}

func (c *UploadAction) Home() {
	fmt.Println("--------------")

	fmt.Println("=============", c.GetForm())

	// switch action {
	// case "config":
	// 	c.Render("ueeditor/config.json")

	// case "uploadimage":
	// 	fmt.Println("upload image")
	// }

}
