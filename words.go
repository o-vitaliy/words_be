package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"words/controllers"
	"words/models"
)

const MysqlUsername = "root"
const MysqlPassword = "12345678"
const MysqlDbName = "words"

func init() {
	print(orm.RegisterDriver("mysql", orm.DRMySQL))
	path := fmt.Sprintf("%s:%s@/%s?charset=utf8", MysqlUsername, MysqlPassword, MysqlDbName)
	print(path)
	print(orm.RegisterDataBase("default", "mysql", path))
}

func main() {
	createTableIfNeed()
	beego.Router("/", &controllers.WordsController{}, "get:Get;post:Post")
	beego.Run()
}

func createTableIfNeed() {
	o := orm.NewOrm()
	orm.Debug = true
	print(o.Using("default"))
	// Database alias.
	name := "default"

	// Drop table and re-create.
	force := false

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	if force {
		c := &controllers.WordsController{}
		for _, v := range c.GetAll() {
			translation := new(models.Translation)
			translation.En = v.Word
			translation.Ru = v.Translation
			o.Insert(translation)
		}
	}

}
