package actions

type User struct {
	Id       int64
	Username string `xorm:"varchar(25) not null unique 'username'"`
	Password string `xorm:"varchar(32) not null 'password'"`
}
