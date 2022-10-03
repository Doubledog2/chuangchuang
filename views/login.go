package views

import (
	"go_boke/common"
	"go_boke/config"
	"net/http"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login

	login.WriteData(w, config.Cfg.Viewer)
}
