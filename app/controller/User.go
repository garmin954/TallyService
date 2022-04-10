package controller

import (
	"fmt"
	"log"
	"net/http"
	"smg/app/middleware"
	"smg/app/model"
	"smg/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/idoubi/goz"
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
	engine.POST("/wxlogin", user.wxlogin)
	engine.POST("/register", user.register)
}

// 登录
func (user *User) login(ctx *gin.Context) {
	var params LoginParams
	if ctx.ShouldBind(&params) != nil {
		utils.Failed(ctx, "失败啊")
		return
	}

	var info model.User
	utils.DB.Table(model.UserTable).Where("username = ?", params.Username).Find(&info)

	if info.ID <= 0 {
		utils.Failed(ctx, "账号不存在")
		return
	}

	pw := utils.Md5(params.Password + info.Salt)
	if pw != info.Password {
		utils.Failed(ctx, "密码不正确")
		return
	}

	claims := middleware.MyClaims{}
	claims.Username = info.Username
	claims.Uid = info.ID
	claims.ExpiresAt = time.Now().Unix() + utils.Configs.Jwt.Expire

	token, err := middleware.NewJWT().CreateToken(claims)
	if err != nil {
		utils.Failed(ctx, "生成token失败")
		return
	}

	utils.Success(ctx, map[string]interface{}{"token": token, "expired": claims.ExpiresAt})
}

type wxloginParams struct {
	Code string `form:"code" json:"code"`
}

// 微信登录
func (user *User) wxlogin(ctx *gin.Context) {
	var params wxloginParams
	ctx.ShouldBind(&params)

	cli := goz.NewClient()

	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		"wx8ee1d2607ce3dd67",
		"5c7b99bbae37070f81c2ac2b1f1dac22",
		params.Code,
	)

	fmt.Println(url)

	resp, err := cli.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T", resp)
	// Output: *goz.Response
	res, _ := resp.GetBody()

	utils.Success(ctx, utils.JsonToMap(res.GetContents()))

}

// 注册
func (user *User) register(ctx *gin.Context) {
	var params RegisterParams
	ctx.ShouldBind(&params)

	salt := utils.RandomStr(10)
	params.Salt = salt
	params.Password = utils.Md5(params.Password + salt)

	result := utils.DB.Table(model.UserTable).Create(&params)

	//TODO 创建默认账本 以及和用户的对应关系

	if result.Error != nil {
		utils.Failed(ctx, result.Error.Error())
		return
	}
	utils.Success(ctx, params)
}

func (user *User) queryUsers(ctx *gin.Context) {
	var results []map[string]interface{}
	utils.DB.Table(model.UserTable).Select("nickname,id").Find(&results)

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"msg":  "成功！",
		"data": results,
	})
}
