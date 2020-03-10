package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"words/models"
)

type WordsController struct {
	beego.Controller
}

func (this *WordsController) GetAll() []*models.Word {
	data, err := ioutil.ReadFile("./sample_words.json")
	if err != nil {
		fmt.Print(err)
	}

	// json data
	var obj []*models.Word
	// unmarshall it
	_ = json.Unmarshal(data, &obj)
	return obj
}

func (this *WordsController) Get() {
	o := orm.NewOrm()
	var tr models.Translation
	_ = o.Raw(`
SELECT * FROM translation
ORDER BY RAND() LIMIT 1
`).QueryRow(&tr)

	this.Data["json"] = models.Word{Word: tr.En, Translation: tr.Ru}
	this.ServeJSON()
}
func (this *WordsController) Post() {
	res := models.Word{Word: "hello2", Translation: "post"}
	this.Data["json"] = res
	this.ServeJSON()
}
