package models

import "github.com/astaxie/beego/orm"

type Word struct {
	Word        string `json:"word"`
	Translation string `json:"translation"`
}

type Translation struct {
	Id int
	En string
	Ru string
}

func init() {
	orm.RegisterModel(new(Translation))
}
