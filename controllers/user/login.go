package user

import (
	"encoding/json"

	"github.com/kataras/iris"
	"myiris/models/response"
	"myiris/library/logger"
)

type User struct {
	Name  string `json:"name"`
	Uid   int64  `json:"uid"`
	Token string `json:"token"`
}

func NewUserInstance() *User {
	return &User{
		Name:  "xuxd",
		Uid:   60000808,
		Token: "XSA1243-DFMDNKD-BBGO",
	}
}

func (u *User) GetInfo(ctx iris.Context) {
	data, _ := json.Marshal(u)
	if _, err := ctx.Write(data); err != nil {
		logger.Error("baseInfo", err)
	}
}

func (u *User) GetToken(ctx iris.Context) {
	if _, err := ctx.WriteString(u.Token); err != nil {
		logger.Error("baseInfo", err)
	}
}

func (u *User) PostUser(ctx iris.Context) {
	var (
		err  error
		code int32
	)
	defer func() {
		res := response.NewResponse(code, err, u)
		if _, err = ctx.JSON(res); err != nil {
			logger.Error("response failed; ", err)
			return
		}
	}()

	if err = ctx.ReadJSON(u); err != nil {
		code = iris.StatusBadRequest
		return
	}
	code = iris.StatusOK
}

func (u *User) Change(ctx iris.Context) {
	var (
		err  error
		code int32
	)

	defer func() {
		res := response.NewResponse(code, err, u)
		if _, err = ctx.JSON(res); err != nil {
			logger.Error("response failed; ", err)
			return
		}
	}()

	// 读取到User
	if err = ctx.ReadJSON(u); err != nil {
		code = iris.StatusBadRequest
		return
	}
	code = iris.StatusOK
	u.Name = ctx.URLParam("name")
}
