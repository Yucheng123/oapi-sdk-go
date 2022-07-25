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

package larkext

import (
	"context"
	"net/http"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
)

func NewService(config *larkcore.Config) *ExtService {
	s := &ExtService{config: config}
	s.DriveExplorer = &driveExplorer{service: s}
	return s
}

// 业务域服务定义
type ExtService struct {
	config        *larkcore.Config
	DriveExplorer *driveExplorer
}

// 资源服务定义
type driveExplorer struct {
	service *ExtService
}

func (d *driveExplorer) CreateFile(ctx context.Context, req *CreateFileReq, options ...larkcore.RequestOptionFunc) (*CreateFileResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/drive/explorer/v2/file/:folderToken"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, d.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateFileResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
