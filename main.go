package main

import (
	"fmt"

	. "github.com/irwinai/blog/actions"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lunny/config"
	"github.com/lunny/xorm"
	"github.com/lunny/xweb"
)

func main() {
	//读取配置文件，创建数据库连接
	cfg, err := config.Load("config.ini")
	if err != nil {
		fmt.Println(err)
		return
	}

	cfgs := cfg.Map()

	// create Orm
	var orm *xorm.Engine
	orm, err = xorm.NewEngine("mysql", fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8",
		cfgs["dbuser"], cfgs["dbpasswd"], cfgs["dbhost"], cfgs["dbname"]))
	if err != nil {
		fmt.Println(err)
		return
	}
	orm.ShowSQL, _ = cfg.GetBool("showSql")
	orm.ShowDebug, _ = cfg.GetBool("showDebug")

	err = orm.Sync(&User{}, &Blog{}, &BlogCategory{}, &BlogTag{}, &Category{}, &Comment{}, &Tag{}, &Picture{})

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	server := xweb.MainServer()
	app := xweb.RootApp()
	app.SetConfig("Orm", orm)
	app.SetConfig("cfg", cfg)
	app.AppConfig.CheckXrsf = false //不验证xrsf

	if useCache, _ := cfg.GetBool("useCache"); useCache {
		server.Info("using orm cache system ...")
		cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
		orm.SetDefaultCacher(cacher)
	}

	//login filter
	// filter := xweb.NewLoginFilter(app, USER_ID, "/login")
	// filter.AddAnonymousUrls("/", "/login", "/about")
	// app.AddFilter(filter)

	// add actions
	xweb.AddAction(&HomeAction{})
	xweb.AutoAction(&BlogAction{}, &ManagerAction{}, &UploadAction{})
	xweb.Run(fmt.Sprintf("%v:%v", cfgs["address"], cfgs["port"]))
}
