package common

import (
	"encoding/json"
	"go_boke/config"
	"go_boke/models"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	//优化
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		//耗时
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			log.Println("模板加载错误")
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}

//获取json参数
func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}
func Error(w http.ResponseWriter, err error) {
	var result models.Result
	result.Code = -999
	result.Error = err.Error()
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
		log.Println(2)
	}
}
func Success(w http.ResponseWriter, data interface{}) {
	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(resultJson)
	if err != nil {
		log.Println(err)

	}
}
