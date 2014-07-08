package actions

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"code.google.com/p/go-uuid/uuid"
	"github.com/lunny/xweb"
)

type UploadAction struct {
	BaseAction

	home xweb.Mapper `xweb:"/"`
}

//处理文件上传
func (c *UploadAction) Home() {
	//根据操作系统获取不同的保存路径
	var path string
	if runtime.GOOS == "windows" {
		path = c.cfg.Map()["filepath_windows"]
	} else {
		path = c.cfg.Map()["filepath_linux"]
	}
	m := make(map[string]interface{})

	//验证图片类型,jpg,jpeg,png,bmp,gif
	h := c.Header("Content-Disposition")
	array := strings.Split(h, ";")
	filename := array[2][11 : len(array[2])-1]
	fileExt := filename[strings.LastIndex(filename, ".")+1 : len(filename)]
	ext := c.cfg.Map()["file_ext"]
	if strings.Index(ext, fileExt) == -1 {
		sendErr("错误的图片类型", c, m)
	}

	//验证图片大小
	size, _ := strconv.Atoi(c.cfg.Map()["file_size"])
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		sendErr("获取图片大小失败", c, m)
	}
	if len(b) > size {
		sendErr("图片不能超过10M", c, m)
	}

	//uuid version1创建保存图片名称
	random := uuid.NewUUID()
	//username := c.Session().Get(USERNAME) //获取当前用户名
	username := "irving" //先写假的
	now := time.Now()
	timeFolder := fmt.Sprintf("%v/%v/%v", now.Year(), int(now.Month()), now.Day()) //按照日创建目录
	path = fmt.Sprintf("%v/%v/%v", path, username, timeFolder)
	os.MkdirAll(path, 0666)
	savePath := fmt.Sprintf("%v/%v.%v", path, random, fileExt)
	fmt.Println("========save path==========", path)
	f, _ := os.OpenFile(savePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		sendErr("上传文件失败，系统异常", c, m)
	}

	//保存文件到数据库
	imagePath := fmt.Sprintf("/upload/%v/%v/%v.%v", username, timeFolder, random, fileExt)
	pic := new(Picture)
	pic.Name = fmt.Sprintf("%v.%v", random, fileExt)
	pic.OldName = filename
	pic.Path = imagePath
	pic.Size = strconv.Itoa(len(b))
	pic.Status = 1
	pic.UserId = 1 //这里先写死
	_, err = c.Orm.Insert(pic)
	if err != nil {
		sendErr("上传文件失败，系统异常", c, m)
	}
	m["err"] = ""
	m["msg"] = imagePath
	c.ServeJson(m)
}

//返回错误
func sendErr(msg string, c *UploadAction, m map[string]interface{}) {
	m["err"] = msg
	m["msg"] = msg
	c.ServeJson(m)
	return
}
