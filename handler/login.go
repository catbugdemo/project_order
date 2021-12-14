package handler

import (
	"github.com/catbugdemo/project_order/inits"
	"github.com/catbugdemo/project_order/service"
	"github.com/gin-gonic/gin"
	"time"
)

// 收录几种登录方式

// 微信扫码登入
type WechatLoginReq struct {
	Code string `json:"code" binding:"required"`
}

// @Summary 微信扫码登录
// @title 后台接口
// @Tags 登录
// @Router /login/wechat [post]
// @param param body WxAppLoginRequestParam true "用户请求参数"
// @Success 200 {object} JsonMsg
func WechatLogin(c *gin.Context) {
	now := time.Now()
	var req WechatLoginReq
	if err := c.BindJSON(&req); err != nil {
		JSONErr(c, Fail, err, req, now)
		return
	}
	m := c.MustGet("secret").(map[string]interface{})
	data, err := service.WechatLogin(req.Code, m["secretId"].(string), m["secretKey"].(string), inits.DB())
	if err != nil {
		JSONErr(c, Fail, err, req, now)
		return
	}
	JSONMsg(c, Success, req, data, now)
	return
}
