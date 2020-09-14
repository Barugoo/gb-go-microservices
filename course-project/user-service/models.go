package main

type User struct {
	ID       int32  `xorm:"id"`
	Email    string `xorm:"email"`
	Password string `xorm:"password"`
}
