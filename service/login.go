package service

import (
	"fmt"
	"github.com/catbugdemo/project_order/utils"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// 微信扫码登入逻辑
func WechatLogin(code, appid, secret string, db *gorm.DB) (string, error) {
	// 获取微信信息
	wechatResp, err := utils.GetWechatResp(code, appid, secret)
	if err != nil {
		return "", errors.WithStack(err)
	}
	fmt.Println(wechatResp)

	// TODO 接下来的身份校验
	// ...

	// 生成 token
	return "", nil
}
