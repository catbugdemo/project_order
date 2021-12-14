package utils

import (
	"fmt"
	"github.com/catbugdemo/project_order/log"
	"github.com/pkg/errors"
)

// 微信返回数据
type WechatResp struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int32  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
}

const (
	// 微信校验地址
	wechat_url = "https://api.weixin.qq.com/sns/oauth2/access_token?code=%s&grant_type=authorization_code&appid=%s&secret=%s"
)

// 获取微信信息
func GetWechatResp(code, appid, secret string) (WechatResp, error) {
	var wr WechatResp
	if err := HttpGet(fmt.Sprintf(wechat_url, code, appid, secret), &wr); err != nil {
		log.Info(err)
		return WechatResp{}, errors.WithStack(err)
	}
	return wr, nil
}
