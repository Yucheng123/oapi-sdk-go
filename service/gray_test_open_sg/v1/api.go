// Package gray_test_open_sg code generated by oapi sdk gen
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

package larkgray_test_open_sg

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

func NewService(config *larkcore.Config) *GrayTestOpenSgService {
	g := &GrayTestOpenSgService{config: config}
	g.Moto = &moto{service: g}
	return g
}

type GrayTestOpenSgService struct {
	config *larkcore.Config
	Moto   *moto // moto
}

type moto struct {
	service *GrayTestOpenSgService
}

//
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer/cli_a24d6887e8b9d00c?from=op_doc_tab&apiName=create&project=gray_test_open_sg&resource=moto&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/gray_test_open_sgv1//create_moto.go
func (m *moto) Create(ctx context.Context, req *CreateMotoReq, options ...larkcore.RequestOptionFunc) (*CreateMotoResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/gray_test_open_sg/v1/motos"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateMotoResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

//
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer/cli_a24d6887e8b9d00c?from=op_doc_tab&apiName=get&project=gray_test_open_sg&resource=moto&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/gray_test_open_sgv1//get_moto.go
func (m *moto) Get(ctx context.Context, req *GetMotoReq, options ...larkcore.RequestOptionFunc) (*GetMotoResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/gray_test_open_sg/v1/motos/:moto_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetMotoResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

//
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer/cli_a24d6887e8b9d00c?from=op_doc_tab&apiName=list&project=gray_test_open_sg&resource=moto&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/gray_test_open_sgv1//list_moto.go
func (m *moto) List(ctx context.Context, req *ListMotoReq, options ...larkcore.RequestOptionFunc) (*ListMotoResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/gray_test_open_sg/v1/motos"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListMotoResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *moto) ListByIterator(ctx context.Context, req *ListMotoReq, options ...larkcore.RequestOptionFunc) (*ListMotoIterator, error) {
	return &ListMotoIterator{
		ctx:      ctx,
		req:      req,
		listFunc: m.List,
		options:  options,
		limit:    req.Limit}, nil
}
