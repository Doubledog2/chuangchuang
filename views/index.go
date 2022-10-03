package views

import (
	"errors"
	"go_boke/common"
	"go_boke/service"
	"log"
	"net/http"
	"strconv"
)

//在 Go 语言中，客户端请求信息都封装到了 Request 对象，但是发送给客户端的响应并不是 Response 对象，而是 ResponseWriter
func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员!!"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	//每页显示的数量
	pageSize := 10
	hr, err := service.GetAllIndexInfo(page, pageSize)
	if err != nil {
		log.Println("Index获取数据出错：", err)
		index.WriteError(w, errors.New("系统出错，请练习管理员"))
	}
	index.WriteData(w, hr)
}
