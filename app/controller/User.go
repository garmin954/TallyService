package controller

import (
	"net/http"
	"smg/app/middleware"
	"smg/app/model"
	"smg/app/service"
	"smg/utils"

	"github.com/gin-gonic/gin"
)

type User struct {
}

type LoginParams struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type RegisterParams struct {
	Username string `form:"username"`
	Mobile   string `form:"mobile"`
	Password string `form:"password"`
	Avatar   string `form:"avatar"`
	Unionid  string `form:"unionid"`
	Salt     string `form:"salt"`
}

// 路由
func (user *User) Router(engine *gin.RouterGroup) {
	group := engine.Group("/user")
	route := group.Use(middleware.JwtToken())
	route.GET("/query_users", user.queryUsers)

	engine.POST("/login", user.login)
	engine.POST("/wxlogin", user.wxLogin)
	engine.POST("/register", user.register)
	engine.POST("/wxregister", user.wxRegister)

}

// 登录
func (user *User) login(ctx *gin.Context) {
	var params LoginParams
	if ctx.ShouldBind(&params) != nil {
		utils.Failed(ctx, "失败啊")
		return
	}

	var SUser service.User
	r, err := SUser.CheckUser(params.Username, params.Password)
	if err != nil {
		utils.Failed(ctx, err.Error())
		return
	}

	utils.Success(ctx, r)
}

type wxLoginParams struct {
	Code string `form:"code" json:"code"`
}

type WXLoginResponse struct {
	UserInfo   *model.User `json:"userInfo"`
	IsRegister bool        `json:"isRegister"`
	service.TokenRes
}

// 微信登录
func (user *User) wxLogin(ctx *gin.Context) {
	var params wxLoginParams
	ctx.ShouldBind(&params)

	var SWechat service.Wechat
	res, err := SWechat.GetCode(params.Code)
	if err != nil {
		utils.Failed(ctx, err.Error())
		return
	}

	response := new(WXLoginResponse)
	r := utils.DB.Table(model.UserTable).Where("openid = ?", res.Openid).Find(&response.UserInfo)
	response.IsRegister = response.UserInfo.ID > 0
	if r.Error != nil {
		utils.Failed(ctx, r.Error.Error())
		return
	}

	if response.UserInfo.ID > 0 {
		var SUser *service.User
		tr, err := SUser.GenerateToken(response.UserInfo.ID, response.UserInfo.Username)
		if err != nil {
			utils.Failed(ctx, err.Error())
			return
		}
		response.Token = tr.Token
		response.Expired = tr.Expired
	}

	utils.Success(ctx, response)
}

type WXRegisterParams struct {
	EncryptedData string `json:"encryptedData" form:"encryptedData"`
	Iv            string `json:"iv" form:"iv"`
	Code          string `json:"code" form:"code"`
}

// 微信注册用户
func (user *User) wxRegister(ctx *gin.Context) {
	var params WXRegisterParams
	ctx.ShouldBind(&params)

	var SWechat service.Wechat
	res, err := SWechat.GetCode(params.Code)
	if err != nil {
		utils.Failed(ctx, err.Error())
		return
	}
	userInfo, err := SWechat.Decrypt(params.EncryptedData, params.Iv, res.SessionKey)
	if err != nil {
		utils.Failed(ctx, err.Error())
		return
	}

	info := model.User{
		Username: userInfo.Nickname,
		Nickname: userInfo.Nickname,
		Avatar:   userInfo.AvatarURL,
		Openid:   res.Openid,
		Password: "gm123654",
	}
	var SUser *service.User
	r, err := SUser.CreateUser(&info)
	if err != nil {
		utils.Failed(ctx, err.Error())
		return
	}

	utils.Success(ctx, r)
}

// 注册
func (user *User) register(ctx *gin.Context) {
	var params model.User
	ctx.ShouldBind(&params)

	var SUser *service.User
	if _, err := SUser.CreateUser(&params); err != nil {
		utils.Failed(ctx, err.Error())
		return
	}
	utils.Success(ctx, params)
}

func (user *User) queryUsers(ctx *gin.Context) {
	var results []map[string]interface{}
	utils.DB.Table(model.UserTable).Select("Nickname,id").Find(&results)

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"msg":  "成功！",
		"data": results,
	})
}
