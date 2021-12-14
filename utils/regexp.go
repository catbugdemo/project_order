package utils

import (
	"github.com/pkg/errors"
	"regexp"
)

const (
	// 手机号校验
	PHONE_REGEX = `^1[0-9]{10}$`

	// 身份证校验
	IDENTIFY_REGEX_18 = `^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`
	IDENTIFY_REGEX_15 = `^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}$`

	// 中文名称校验
	CHINESE_NAME_REGEX = "^[\u4E00-\u9FA5]{2,5}(?:·[\u4E00-\u9FA5]{2,5})*$"

	// 银行卡校验
	BANKCARD_REGEX = `^([1-9]{1})(\d{12,18})$`
)

// 手机校验是否正确
func PhoneRegexp(phone string) error {
	regex, _ := regexp.Compile(PHONE_REGEX)
	if !regex.MatchString(phone) {
		return errors.New("电话号码非法！")
	}
	return nil
}

// 身份证校验
func IdentifyRegexp(idCard string) error {
	regex18, _ := regexp.Compile(IDENTIFY_REGEX_18)
	regex15, _ := regexp.Compile(IDENTIFY_REGEX_15)
	if !regex18.MatchString(idCard) && !regex15.MatchString(idCard) {
		return errors.New("身份证号码非法!")
	}
	return nil
}

// 中文名校验
func NameRegexp(name string) error {
	regex, _ := regexp.Compile(CHINESE_NAME_REGEX)
	if !regex.MatchString(name) {
		return errors.New("名字非法!")
	}
	return nil
}

// 银行卡校验
func CardRegexp(bankCard string) error {
	regex, _ := regexp.Compile(BANKCARD_REGEX)
	if !regex.MatchString(bankCard) {
		return errors.New("银行卡号非法!")
	}
	return nil
}
