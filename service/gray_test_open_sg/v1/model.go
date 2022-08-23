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
	"fmt"

	"context"
	"errors"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

type Level struct {
	Level *string `json:"level,omitempty"` // string
	Body  *string `json:"body,omitempty"`  // body
	Type  *string `json:"type,omitempty"`  // type
}

type LevelBuilder struct {
	level     string // string
	levelFlag bool
	body      string // body
	bodyFlag  bool
	type_     string // type
	typeFlag  bool
}

func NewLevelBuilder() *LevelBuilder {
	builder := &LevelBuilder{}
	return builder
}

// string
// 示例值：
func (builder *LevelBuilder) Level(level string) *LevelBuilder {
	builder.level = level
	builder.levelFlag = true
	return builder
}

// body
// 示例值：
func (builder *LevelBuilder) Body(body string) *LevelBuilder {
	builder.body = body
	builder.bodyFlag = true
	return builder
}

// type
// 示例值：
func (builder *LevelBuilder) Type(type_ string) *LevelBuilder {
	builder.type_ = type_
	builder.typeFlag = true
	return builder
}

func (builder *LevelBuilder) Build() *Level {
	req := &Level{}
	if builder.levelFlag {
		req.Level = &builder.level

	}
	if builder.bodyFlag {
		req.Body = &builder.body

	}
	if builder.typeFlag {
		req.Type = &builder.type_

	}
	return req
}

type Moto struct {
	MotoId   *string `json:"moto_id,omitempty"`   // desc
	Id       *string `json:"id,omitempty"`        // desc2
	UserName *string `json:"user_name,omitempty"` // name
	Type     *string `json:"type,omitempty"`      // type
}

type MotoBuilder struct {
	motoId       string // desc
	motoIdFlag   bool
	id           string // desc2
	idFlag       bool
	userName     string // name
	userNameFlag bool
	type_        string // type
	typeFlag     bool
}

func NewMotoBuilder() *MotoBuilder {
	builder := &MotoBuilder{}
	return builder
}

// desc
// 示例值：
func (builder *MotoBuilder) MotoId(motoId string) *MotoBuilder {
	builder.motoId = motoId
	builder.motoIdFlag = true
	return builder
}

// desc2
// 示例值：
func (builder *MotoBuilder) Id(id string) *MotoBuilder {
	builder.id = id
	builder.idFlag = true
	return builder
}

// name
// 示例值：
func (builder *MotoBuilder) UserName(userName string) *MotoBuilder {
	builder.userName = userName
	builder.userNameFlag = true
	return builder
}

// type
// 示例值：
func (builder *MotoBuilder) Type(type_ string) *MotoBuilder {
	builder.type_ = type_
	builder.typeFlag = true
	return builder
}

func (builder *MotoBuilder) Build() *Moto {
	req := &Moto{}
	if builder.motoIdFlag {
		req.MotoId = &builder.motoId

	}
	if builder.idFlag {
		req.Id = &builder.id

	}
	if builder.userNameFlag {
		req.UserName = &builder.userName

	}
	if builder.typeFlag {
		req.Type = &builder.type_

	}
	return req
}

type CreateMotoReqBuilder struct {
	apiReq *larkcore.ApiReq
	level  *Level
}

func NewCreateMotoReqBuilder() *CreateMotoReqBuilder {
	builder := &CreateMotoReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// desc
//
// 示例值：
func (builder *CreateMotoReqBuilder) DepartmentIdType(departmentIdType string) *CreateMotoReqBuilder {
	builder.apiReq.QueryParams.Set("department_id_type", fmt.Sprint(departmentIdType))
	return builder
}

//
func (builder *CreateMotoReqBuilder) Level(level *Level) *CreateMotoReqBuilder {
	builder.level = level
	return builder
}

func (builder *CreateMotoReqBuilder) Build() *CreateMotoReq {
	req := &CreateMotoReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.QueryParams = builder.apiReq.QueryParams
	req.apiReq.Body = builder.level
	return req
}

type CreateMotoReq struct {
	apiReq *larkcore.ApiReq
	Level  *Level `body:""`
}

type CreateMotoRespData struct {
	Moto *Moto `json:"moto,omitempty"` // desc22222
}

type CreateMotoResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *CreateMotoRespData `json:"data"` // 业务数据
}

func (resp *CreateMotoResp) Success() bool {
	return resp.Code == 0
}

type GetMotoReqBuilder struct {
	apiReq *larkcore.ApiReq
}

func NewGetMotoReqBuilder() *GetMotoReqBuilder {
	builder := &GetMotoReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// desc33333
//
// 示例值：
func (builder *GetMotoReqBuilder) MotoId(motoId string) *GetMotoReqBuilder {
	builder.apiReq.PathParams.Set("moto_id", fmt.Sprint(motoId))
	return builder
}

// desc
//
// 示例值：
func (builder *GetMotoReqBuilder) BodyLevel(bodyLevel string) *GetMotoReqBuilder {
	builder.apiReq.QueryParams.Set("body_level", fmt.Sprint(bodyLevel))
	return builder
}

func (builder *GetMotoReqBuilder) Build() *GetMotoReq {
	req := &GetMotoReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.PathParams = builder.apiReq.PathParams
	req.apiReq.QueryParams = builder.apiReq.QueryParams
	return req
}

type GetMotoReq struct {
	apiReq *larkcore.ApiReq
}

type GetMotoRespData struct {
	Moto *Moto `json:"moto,omitempty"` // desc
}

type GetMotoResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *GetMotoRespData `json:"data"` // 业务数据
}

func (resp *GetMotoResp) Success() bool {
	return resp.Code == 0
}

type ListMotoReqBuilder struct {
	apiReq *larkcore.ApiReq
	limit  int // 最大返回多少记录，当使用迭代器访问时才有效
}

func NewListMotoReqBuilder() *ListMotoReqBuilder {
	builder := &ListMotoReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 最大返回多少记录，当使用迭代器访问时才有效
func (builder *ListMotoReqBuilder) Limit(limit int) *ListMotoReqBuilder {
	builder.limit = limit
	return builder
}

// 分页大小
//
// 示例值：
func (builder *ListMotoReqBuilder) PageSize(pageSize int) *ListMotoReqBuilder {
	builder.apiReq.QueryParams.Set("page_size", fmt.Sprint(pageSize))
	return builder
}

// 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果
//
// 示例值：
func (builder *ListMotoReqBuilder) PageToken(pageToken string) *ListMotoReqBuilder {
	builder.apiReq.QueryParams.Set("page_token", fmt.Sprint(pageToken))
	return builder
}

// level
//
// 示例值：
func (builder *ListMotoReqBuilder) Level(level int) *ListMotoReqBuilder {
	builder.apiReq.QueryParams.Set("level", fmt.Sprint(level))
	return builder
}

func (builder *ListMotoReqBuilder) Build() *ListMotoReq {
	req := &ListMotoReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.Limit = builder.limit
	req.apiReq.QueryParams = builder.apiReq.QueryParams
	return req
}

type ListMotoReq struct {
	apiReq *larkcore.ApiReq
	Limit  int // 最多返回多少记录，只有在使用迭代器访问时，才有效

}

type ListMotoRespData struct {
	Items     []string `json:"items,omitempty"`      // desc
	PageToken *string  `json:"page_token,omitempty"` //
	HasMore   *bool    `json:"has_more,omitempty"`   //
}

type ListMotoResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *ListMotoRespData `json:"data"` // 业务数据
}

func (resp *ListMotoResp) Success() bool {
	return resp.Code == 0
}

type ListMotoIterator struct {
	nextPageToken *string
	items         []string
	index         int
	limit         int
	ctx           context.Context
	req           *ListMotoReq
	listFunc      func(ctx context.Context, req *ListMotoReq, options ...larkcore.RequestOptionFunc) (*ListMotoResp, error)
	options       []larkcore.RequestOptionFunc
	curlNum       int
}

func (iterator *ListMotoIterator) Next() (bool, string, error) {
	// 达到最大量，则返回
	if iterator.limit > 0 && iterator.curlNum >= iterator.limit {
		return false, "", nil
	}

	// 为0则拉取数据
	if iterator.index == 0 || iterator.index >= len(iterator.items) {
		if iterator.index != 0 && iterator.nextPageToken == nil {
			return false, "", nil
		}
		if iterator.nextPageToken != nil {
			iterator.req.apiReq.QueryParams.Set("page_token", *iterator.nextPageToken)
		}
		resp, err := iterator.listFunc(iterator.ctx, iterator.req, iterator.options...)
		if err != nil {
			return false, "", err
		}

		if resp.Code != 0 {
			return false, "", errors.New(fmt.Sprintf("Code:%d,Msg:%s", resp.Code, resp.Msg))
		}

		if len(resp.Data.Items) == 0 {
			return false, "", nil
		}

		iterator.nextPageToken = resp.Data.PageToken
		iterator.items = resp.Data.Items
		iterator.index = 0
	}

	block := iterator.items[iterator.index]
	iterator.index++
	iterator.curlNum++
	return true, block, nil
}

func (iterator *ListMotoIterator) NextPageToken() *string {
	return iterator.nextPageToken
}
