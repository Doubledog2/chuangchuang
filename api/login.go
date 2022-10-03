package api

import (
	"go_boke/common"
	"go_boke/service"
	"log"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		log.Println("api.Login() fail")
		log.Println(err)
		log.Println(loginRes)
		common.Error(w, err)
		return
	}
	common.Success(w, loginRes)

}
