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
	fmt.Println("=======================================")
	file, head, err := c.GetFile("filedata")
	if err != nil {
		fmt.Println(err)
		return
	}
	//sroot := c.cfg.Map()["filepath"]
	fmt.Println("========================================================")
	fmt.Println(head.Filename, "----------------")
	fmt.Println(file, "----------------")
	c.ServeJson(map[string]interface{}{"err": "", "msg": "200906030521128703.gif"})
	// f, err := os.OpenFile(, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	// if err != nil {
	// 	return err
	// }
	// defer f.Close()
	// _, err = io.Copy(f, file)
	// return err
	//	return xweb.NotSupported()
}
