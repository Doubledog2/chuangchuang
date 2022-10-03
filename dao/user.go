package dao

import (
	"fmt"
	"go_boke/models"
	"log"
)

func GetUserNameById(uId int) string {
	row := DB.QueryRow("select user_name from blog_users where uid=?", uId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var userName string
	_ = row.Scan(&userName)
	return userName
}

func GetUser(userName, passwd string) *models.User {
	row := DB.QueryRow(
		"select * from  blog_users where user_name=? and passwd=? limit 1",
		userName,
		passwd,
	)
	fmt.Println("row:", row)
	if row.Err() != nil {
		log.Println("blog_user获取失败")
		log.Println(row.Err())
		return nil
	}
	var user = &models.User{}
	err := row.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println("blog_user赋值失败") //bug
		log.Println(err)

		return nil
	}
	return user
}