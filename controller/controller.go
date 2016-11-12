package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Unknwon/macaron"
)

type GetResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Body string `json:"body"`
}

type PostRequest struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

const (
	CODE_SUCCESS = 10000
	MSG_SUCCESS  = "OK"
)

func GetMethod(ctx *macaron.Context) string {
	response := new(GetResponse)
	response.Code = CODE_SUCCESS
	response.Msg = MSG_SUCCESS
	response.Body = "when you see this msg,the result is ok"
	ret_str, _ := json.Marshal(response)
	return string(ret_str)
}

func PostMethod(ctx *macaron.Context) string {
	req := new(PostRequest)
	body, _ := ctx.Req.Body().String()
	json.Unmarshal([]byte(body), req)
	response := new(GetResponse)
	if req.UserName == "ricky" && req.PassWord == "123456" {
		response.Code = CODE_SUCCESS
		response.Msg = MSG_SUCCESS
		response.Body = "success"
	} else {
		response.Code = 10001
		response.Msg = "error"
		response.Body = "wrong user name or password"
	}
	ret_str, _ := json.Marshal(response)
	return string(ret_str)
}

type UserName struct {
	UserName string
	PassWord string
}

func (user *UserName) InitInfo() UserName {
	u := new(UserName)
	u.UserName = "zhangsan"
	u.PassWord = "123456"
	return *u
}
func (user *UserName) Auth(password string) int {
	if user.PassWord == "123456" {
		return http.StatusOK
	}
	return http.StatusBadRequest
}
