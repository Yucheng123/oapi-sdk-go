// Package human_authentication code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package larkhuman_authentication

import (
	"context"
	"net/http"

	"github.com/Yucheng123/oapi-sdk-go/v3/core"
)

func NewService(config *larkcore.Config) *HumanAuthenticationService {
	h := &HumanAuthenticationService{config: config}
	h.Identity = &identity{service: h}
	return h
}

type HumanAuthenticationService struct {
	config   *larkcore.Config
	Identity *identity // 实名认证
}

type identity struct {
	service *HumanAuthenticationService
}

// 录入身份信息
//
// - 该接口用于录入实名认证的身份信息，在唤起有源活体认证前，需要使用该接口进行实名认证。
//
// - 实名认证接口会有计费管理，接入前请联系飞书开放平台工作人员，邮箱：openplatform@bytedance.com。;;仅通过计费申请的应用，才能在[开发者后台](https://open.feishu.cn/app)查找并申请该接口的权限。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/human_authentication-v1/identity/create
//
// - 使用Demo链接:https://github.com/Yucheng123/oapi-sdk-go/tree/v3_main/sample/apiall/human_authenticationv1/create_identity.go
func (i *identity) Create(ctx context.Context, req *CreateIdentityReq, options ...larkcore.RequestOptionFunc) (*CreateIdentityResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/human_authentication/v1/identities"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateIdentityResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
