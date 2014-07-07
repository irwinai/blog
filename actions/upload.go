package actions

import (
	"fmt"
	"io"
	"os"
	"strings"

	"code.google.com/p/go-uuid/uuid"
	"github.com/lunny/xweb"
)

type UploadAction struct {
	BaseAction

	home xweb.Mapper `xweb:"/"`
}

//处理文件上传
func (c *UploadAction) Home() {
	path := c.cfg.Map()["filepath"]
	fmt.Println("===============file path========================", path)
	fmt.Println("===============file size=======================", c.Request.ContentLength/1024/1024)

	h := c.Header("Content-Disposition")
	array := strings.Split(h, ";")
	filename := array[2][11 : len(array[2])-1]
	fmt.Println("===============file ext=======================", filename)

	random := uuid.NewUUID()
	os.MkdirAll(path, 0666)
	f, err := os.OpenFile(fmt.Sprintf("%v\\%v.jpg", path, random), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	io.Copy(f, c.Request.Body)
	io.Copy(c.ResponseWriter, f)
	c.ServeJson(map[string]interface{}{"err": "", "msg": "200906030521128703.gif"})
}
