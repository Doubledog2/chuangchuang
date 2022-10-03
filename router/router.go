package router

import (
	"go_boke/api"
	"go_boke/views"
	"net/http"
)

func Router() {
	//1. 页面  views 2. api 数据（json） 3. 静态资源
	http.HandleFunc("/", views.HTML.Index)
	//http://localhost:8080/c/1  1参数 分类的id
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/login", views.HTML.Login)
	//http://localhost:8080/p/7.html
	http.HandleFunc("/p/", views.HTML.Detail)
	http.HandleFunc("/c/delete/", views.HTML.DeleteCategoryItem)
	//写文章
	http.HandleFunc("/writing", views.HTML.Writing)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)

	http.HandleFunc("/api/v1/login", api.API.Login)
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost) //搜索
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
