package views

import (
	"errors"
	"fmt"
	"go_boke/common"
	"go_boke/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail
	//获取参数路径
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/p/")
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("不识别此请求路径"))
		return
	}
	postRes, err := service.GetPostDetail(pId)
	if err != nil {
		detail.WriteError(w, errors.New("查询出错"))
		return
	}
	detail.WriteData(w, postRes)
}
func (*HTMLApi) DeleteCategoryItem(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println("api deletepost success  " + path)
	pIdStr := strings.TrimPrefix(path, "/c/delete/")
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		log.Println("不识别此请求路径", err)
		common.Error(w, errors.New("不识别此请求路径"))
		return
	}
	service.DeletePost(pid)
	HTML.Category(w, r)
}
