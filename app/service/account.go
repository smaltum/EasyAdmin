package service

import (
	"context"
	"easy-admin/app/model"
	"easy-admin/plugin/resp"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var Account = &account{}

type account struct {
}

// Decrypt 消息解密
func (*account) Decrypt(ctx context.Context, requestParam *model.AccountDecryptReq) (*model.AccountPlainDataResp, *resp.Exception) {

	// 发起请求
	var wxBaseUrl strings.Builder
	//wxBaseUrl.WriteString(fmt.Sprintf("%s?%s=", "appid", boot.Config.MinProgram.AppId))
	//wxBaseUrl.WriteString(fmt.Sprintf("%s&%s=", "secret", boot.Config.MinProgram.Secret))
	wxBaseUrl.WriteString(fmt.Sprintf("%s&%s=", "js_code", requestParam.Code))
	wxBaseUrl.WriteString(fmt.Sprintf("%s&%s=", "grant_type", "authorization_code"))
	respData, err := http.Get(wxBaseUrl.String())
	if err != nil {
		return nil, resp.NewException(resp.ServerError, "请求sessionKey失败")
	}
	defer func() {
		_ = respData.Body.Close()
	}()

	body, err := ioutil.ReadAll(respData.Body)
	if err != nil {
		return nil, resp.NewException(resp.ServerError, "读取sessionKey失败")
	}

	// 解析
	var accountPlainDataResp model.AccountPlainDataResp
	_ = json.Unmarshal([]byte(body), &accountPlainDataResp)

	if accountPlainDataResp.OpenID == "" {
		return nil, resp.NewException(resp.ServerError, "code已失效,请重新登录")
	}

	// todo 记录

	return &accountPlainDataResp, nil
}
