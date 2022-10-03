package views

import (
	"go_boke/common"
	"go_boke/service"
	"log"
	"net/http"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr := service.Writing()
	log.Println(wr.CdnURL)
	writing.WriteData(w, wr)
}
