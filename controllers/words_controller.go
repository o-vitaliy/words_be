package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"net/http"
	"words/models"
)

func GetRandom(w http.ResponseWriter, r *http.Request) {
	o := orm.NewOrm()
	var tr models.Translation
	_ = o.Raw(`
SELECT * FROM translation
ORDER BY RAND() LIMIT 1
`).QueryRow(&tr)
	println(tr.En)
	json.NewEncoder(w).Encode(models.Word{Word: tr.En, Translation: tr.Ru})
}
