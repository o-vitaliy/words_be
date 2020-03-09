package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"math/rand"
	"words/models"
)

type WordsController struct {
	beego.Controller
}

func (this *WordsController) Get() {
	data, err := ioutil.ReadFile("./sample_words.json")
	if err != nil {
		fmt.Print(err)
	}

	// json data
	var obj []*models.Word
	fmt.Println(string(data))
	// unmarshall it
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}
	r := rand.Intn(len(obj))
	this.Data["json"] = obj[r]
	this.ServeJSON()
}
func (this *WordsController) Post() {
	res := models.Word{Word: "hello2", Translation: "post"}
	this.Data["json"] = res
	this.ServeJSON()
}
