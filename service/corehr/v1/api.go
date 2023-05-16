// Package corehr code generated by oapi sdk gen
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

package larkcorehr

import (
	"context"
	"net/http"

	"github.com/Yucheng123/oapi-sdk-go/v3/core"
)

func NewService(config *larkcore.Config) *CorehrService {
	c := &CorehrService{config: config}
	c.Leave = &leave{service: c}
	return c
}

type CorehrService struct {
	config *larkcore.Config
	Leave  *leave // 休假管理
}

type leave struct {
	service *CorehrService
}

// 批量查询员工请假记录
//
// - 批量获取员工的请假记录数据
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/leave/leave_request_history
//
// - 使用Demo链接:https://github.com/Yucheng123/oapi-sdk-go/tree/v3_main/sample/apiall/corehrv1/leaveRequestHistory_leave.go
func (l *leave) LeaveRequestHistory(ctx context.Context, req *LeaveRequestHistoryLeaveReq, options ...larkcore.RequestOptionFunc) (*LeaveRequestHistoryLeaveResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/corehr/v1/leaves/leave_request_history"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, l.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &LeaveRequestHistoryLeaveResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, l.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
