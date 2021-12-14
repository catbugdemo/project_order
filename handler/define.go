package handler

import (
	"github.com/catbugdemo/project_order/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ErrCode int

//go:generate stringer -type ErrCode -linecomment
const (
	Success ErrCode = iota // success
	Fail                   // fail
)

// JSONErr return err info
func JSONErr(c *gin.Context, errNo ErrCode, err error, req interface{}, timeNow time.Time) {
	log.Infof("request: url:%s client_ip:%s body:%v cost:%v err:%+v",
		c.Request.RequestURI, c.ClientIP(), req, time.Since(timeNow), err)
	c.JSON(http.StatusOK, gin.H{
		"error_no":  errNo,
		"error_msg": errNo.String(),
	})
}

// JSONMsg return info
func JSONMsg(c *gin.Context, errNo ErrCode, req, resp interface{}, timeNow time.Time) {
	log.Infof("request: url:%s client_ip:%s body:%v cost:%v resp:%v",
		c.Request.RequestURI, c.ClientIP(), req, time.Since(timeNow), resp)
	c.JSON(http.StatusOK, gin.H{
		"error_no":  errNo,
		"error_msg": errNo.String(),
		"data":      resp,
	})
}
