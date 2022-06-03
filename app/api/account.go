package api

import (
	"easy-admin/app/model"
	"easy-admin/app/service"
	"easy-admin/plugin/resp"
	"github.com/gin-gonic/gin"
)

type account struct {
}

// Test 测试
func (*account) Test(c *gin.Context) {

	resp.New(c).ResponseOK("hello")
}

// Decrypt 消息解密
func (*account) Decrypt(c *gin.Context) {
	// 解析
	var requestParam model.AccountDecryptReq
	err := c.ShouldBind(&requestParam)
	if err != nil {
		resp.New(c).ResponseException(resp.NewException(resp.BadRequest, err.Error()))
	}

	// 解密
	respData, respExp := service.Account.Decrypt(c, &requestParam)
	if respExp != nil {
		resp.New(c).ResponseException(respExp)
		return
	}

	resp.New(c).ResponseOK(respData)
}
