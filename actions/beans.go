package actions

import "time"

type User struct {
	Id            int       `xorm:"not null pk autoincr INT(10)"`
	Username      string    `xorm:"not null VARCHAR(32)"`
	Nickname      string    `xorm:"VARCHAR(32)"`
	Avatar        string    `xorm:"VARCHAR(100)"`
	Password      string    `xorm:"not null VARCHAR(32)"`
	Sex           int       `xorm:"default 1 TINYINT(4)"`
	Age           int       `xorm:"TINYINT(4)"`
	Colleage      string    `xorm:"VARCHAR(50)"`
	Hobby         string    `xorm:"VARCHAR(200)"`
	LastLoginTime time.Time `xorm:"DATETIME"`
	LastLoginIp   string    `xorm:"VARCHAR(100)"`
	RegisterTime  time.Time `xorm:"DATETIME"`
	Status        int       `xorm:"not null default 1 TINYINT(4)"`
	Sign          string    `xorm:"VARCHAR(200)"`
	Skin          int       `xorm:"INT(11)"`
}

type Blog struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	Title        string    `xorm:"not null VARCHAR(50)"`
	Content      string    `xorm:"not null TEXT"`
	UserId       int       `xorm:"not null INT(11)"`
	CreateTime   time.Time `xorm:"not null DATETIME"`
	UpdateTime   time.Time `xorm:"DATETIME"`
	DeleteTime   time.Time `xorm:"DATETIME"`
	Views        int       `xorm:"default 0 INT(11)"`
	CommentCount int       `xorm:"default 0 INT(11)"`
	Like         int       `xorm:"default 0 INT(11)"`
	Status       int       `xorm:"not null default 1 TINYINT(4)"`
}

type BlogCategory struct {
	Id         int `xorm:"not null pk autoincr INT(10)"`
	BlogId     int `xorm:"not null INT(11)"`
	CategoryId int `xorm:"not null INT(11)"`
}

type BlogTag struct {
	Id     int `xorm:"not null pk autoincr INT(10)"`
	BlogId int `xorm:"not null INT(11)"`
	TagId  int `xorm:"not null INT(11)"`
}

type Category struct {
	Id     int    `xorm:"not null pk autoincr INT(11)"`
	Name   string `xorm:"not null VARCHAR(10)"`
	Status int    `xorm:"not null default 1 TINYINT(4)"`
}

type Comment struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Content    string    `xorm:"not null VARCHAR(1000)"`
	CreateTime time.Time `xorm:"not null DATETIME"`
	UpdateTime time.Time `xorm:"DATETIME"`
	DeleteTime time.Time `xorm:"DATETIME"`
	Username   string    `xorm:"not null VARCHAR(32)"`
	Email      string    `xorm:"VARCHAR(100)"`
	Website    string    `xorm:"VARCHAR(200)"`
	ParentId   int       `xorm:"default 0 INT(11)"`
	BlogId     int       `xorm:"not null INT(11)"`
	UserId     int       `xorm:"INT(11)"`
	Ip         string    `xorm:"VARCHAR(50)"`
	Status     int       `xorm:"not null TINYINT(4)"`
}

type Tag struct {
	Id     int    `xorm:"not null pk autoincr INT(11)"`
	Name   string `xorm:"not null VARCHAR(10)"`
	Status int    `xorm:"not null TINYINT(4)"`
}
