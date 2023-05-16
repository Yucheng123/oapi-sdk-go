// Package hire code generated by oapi sdk gen
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

// POST /open-apis/hire/v1/jobs/combined_create
func main() {
	// 创建 Client
	client := lark.NewClient("appID", "appSecret")
	// 创建请求对象
	req := larkhire.NewCombinedCreateJobReqBuilder().
		UserIdType("open_id").
		DepartmentIdType("open_department_id").
		CombinedJob(larkhire.NewCombinedJobBuilder().
			Code("R18").
			Experience(1).
			ExpiryTime(1622484739955).
			CustomizedDataList([]*larkhire.CombinedJobObjectValueMap{larkhire.NewCombinedJobObjectValueMapBuilder().Build()}).
			MinLevelId("6960663240925956547").
			MinSalary(1000).
			Title("后端研发").
			JobManagers(larkhire.NewJobManagerBuilder().Build()).
			JobProcessId("6960663240925956554").
			ProcessType(1).
			SubjectId("6960663240925956555").
			JobFunctionId("6960663240925956555").
			DepartmentId("6960663240925956549").
			HeadCount(100).
			IsNeverExpired(false).
			MaxSalary(2000).
			Requirement("熟悉后端研发").
			AddressId("6960663240925956553").
			Description("后端研发岗位描述").
			HighlightList([]string{}).
			JobTypeId("6960663240925956551").
			MaxLevelId("6960663240925956548").
			RecruitmentTypeId("6960663240925956552").
			RequiredDegree(20).
			JobCategoryId("6960663240925956550").
			AddressIdList([]string{}).
			JobAttribute(1).
			ExpiryTimestamp("1622484739955").
			InterviewRegistrationSchemaId("6930815272790114324").
			Build()).
		Build()
	// 发起请求
	resp, err := client.Hire.Job.CombinedCreate(context.Background(), req)

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
