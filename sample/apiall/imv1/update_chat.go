// Package im code generated by oapi sdk gen
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

package main

import (
	"context"
	"fmt"
)

// PUT /open-apis/im/v1/chats/:chat_id
func main() {
	// 创建 Client
	client := lark.NewClient("appID", "appSecret")
	// 创建请求对象
	req := larkim.NewUpdateChatReqBuilder().
		ChatId("oc_a0553eda9014c201e6969b478895c230").
		UserIdType("user_id").
		Body(larkim.NewUpdateChatReqBodyBuilder().
			Avatar("default-avatar_44ae0ca3-e140-494b-956f-78091e348435").
			Name("群聊").
			Description("测试群描述").
			I18nNames(larkim.NewI18nNamesBuilder().Build()).
			AddMemberPermission("all_members").
			ShareCardPermission("allowed").
			AtAllPermission("all_members").
			EditPermission("all_members").
			OwnerId("4d7a3c6g").
			JoinMessageVisibility("only_owner").
			LeaveMessageVisibility("only_owner").
			MembershipApproval("no_approval_required").
			RestrictedModeSetting(larkim.NewRestrictedModeSettingBuilder().Build()).
			ChatType("private").
			Build()).
		Build()
	// 发起请求
	resp, err := client.Im.Chat.Update(context.Background(), req)

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return
	}

	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
}
