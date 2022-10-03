package main

import (
	"go_boke/common"
	"go_boke/router"
	"log"
	"net/http"
)

func init() {
	//模板加载
	common.LoadTemplate()
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8083",
	}
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
