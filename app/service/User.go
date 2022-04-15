package service

import (
	"errors"
	"smg/app/middleware"
	"smg/app/model"
	"smg/utils"
	"time"
)

type User struct{}

// 创建用户
func (user *User) CreateUser(params *model.User) (*model.User, error) {

	var info model.User
	utils.DB.Table(model.UserTable).Where("username = ?", params.Username).Find(&info)

	if info.ID > 0 {
		return nil, errors.New("account already exists")
	}

	salt := utils.RandomStr(10)
	params.Salt = salt
	params.Password = utils.Md5(params.Password + salt)

	tx := utils.DB.Begin()
	result := tx.Table(model.UserTable).Create(&params)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	var MUser model.MUser
	err := MUser.InitUserLedger(tx, params.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return params, nil
}

type CheckUserResponse struct {
	UserInfo *model.User `json:"userInfo"`
	Expired  int64       `json:"expired"`
	Token    string      `json:"token"`
}

// 验证登录并返回token
func (user *User) CheckUser(username string, password string) (*CheckUserResponse, error) {
	rep := new(CheckUserResponse)
	r := utils.DB.Table(model.UserTable).Where("username = ?", username).Find(&rep.UserInfo)
	if r.Error != nil {
		return nil, r.Error
	}

	if rep.UserInfo.ID <= 0 {
		return nil, errors.New("account not exist")
	}

	pw := utils.Md5(password + rep.UserInfo.Salt)
	if pw != rep.UserInfo.Password {
		return nil, errors.New("incorrect password")
	}

	tr, err := user.GenerateToken(rep.UserInfo.ID, rep.UserInfo.Username)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	rep.Token = tr.Token
	rep.Expired = tr.Expired

	return rep, nil
}

type TokenRes struct {
	Expired int64  `json:"expired"`
	Token   string `json:"token"`
}

// 生成token
func (user *User) GenerateToken(id int, username string) (*TokenRes, error) {
	claims := middleware.MyClaims{}
	claims.Username = username
	claims.Uid = id
	claims.ExpiresAt = time.Now().Unix() + utils.Configs.Jwt.Expire

	token, err := middleware.NewJWT().CreateToken(claims)
	if err != nil {
		return nil, err
	}

	r := new(TokenRes)
	r.Token = token
	r.Expired = claims.ExpiresAt
	return r, nil
}

// func (user *User) EncryptPwd(pwd string, salt string) string {

// }
