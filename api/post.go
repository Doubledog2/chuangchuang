package api

import (
	"errors"
	"go_boke/common"
	"go_boke/dao"
	"go_boke/models"
	"go_boke/service"
	"go_boke/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/api/v1/post/")
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		common.Error(w, errors.New("不识别此请求路径"))
		return
	}
	post, err := dao.GetPostById(pid)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, post)
}

func (*Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	//获取用户id，判断用户是否登录
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登录已过期"))
		return
	}
	uid := claim.Uid
	//POST  save
	method := r.Method
	switch method {
	case http.MethodPost: //save
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		post := &models.Post{
			-1,
			title,
			content,
			markdown,
			categoryId,
			uid,
			0,
			pType,
			slug,
			time.Now(),
			time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut: //update
		// update
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pidFloat := params["pid"].(float64)
		pType := int(postType)
		pid := int(pidFloat)
		post := &models.Post{
			pid,
			title,
			content,
			markdown,
			categoryId,
			uid,
			0,
			pType,
			slug,
			time.Now(),
			time.Now(),
		}
		service.UpdatePost(post)
		common.Success(w, post)
	}
}

//搜索post文章
func (*Api) SearchPost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	condition := r.Form.Get("val")
	searchResp := service.SearchPost(condition)
	common.Success(w, searchResp)
}
