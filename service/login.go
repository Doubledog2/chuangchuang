package service

import (
	"errors"
	"go_boke/dao"
	"go_boke/models"
	"go_boke/utils"
	"log"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	log.Println("账号密码old：", userName, passwd)
	passwd = utils.Md5Crypt(passwd, "mszlu")
	log.Println("账号密码new：", userName, passwd)
	user := dao.GetUser(userName, passwd)
	if user == nil {
		log.Println("账号密码不正确")
		return nil, errors.New("账号密码不正确")
	}
	uid := user.Uid
	//生成token  jwt技术进行生成 令牌  A.B.C
	token, err := utils.Award(&uid)
	if err != nil {
		log.Println("token未能生成")
		return nil, errors.New("token未能生成")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		token,
		userInfo,
	}

	return lr, nil
}
