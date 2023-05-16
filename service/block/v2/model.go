// Package block code generated by oapi sdk gen
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

package larkblock

import (
	"fmt"

	"github.com/Yucheng123/oapi-sdk-go/v3/core"
)

type Entity struct {
	BlockId     *string `json:"block_id,omitempty"`      // block 唯一标识
	Title       *string `json:"title,omitempty"`         // 标题
	BlockTypeId *string `json:"block_type_id,omitempty"` // block 类型ID
	SourceData  *string `json:"source_data,omitempty"`   // 业务数据 json
	SourceMeta  *string `json:"source_meta,omitempty"`   // 元数据 json
	Version     *string `json:"version,omitempty"`       // 版本号(自增)
	SourceLink  *string `json:"source_link,omitempty"`   // 链接
	Summary     *string `json:"summary,omitempty"`       // 总括
	Preview     *string `json:"preview,omitempty"`       // 预览
	I18nSummay  *string `json:"i18n_summay,omitempty"`   // 综述 json
	I18nPreview *string `json:"i18n_preview,omitempty"`  // 预览 json
	Owner       *string `json:"owner,omitempty"`         // 所有者
	Extra       *string `json:"extra,omitempty"`         // 扩展字段 json
}

type EntityBuilder struct {
	blockId         string // block 唯一标识
	blockIdFlag     bool
	title           string // 标题
	titleFlag       bool
	blockTypeId     string // block 类型ID
	blockTypeIdFlag bool
	sourceData      string // 业务数据 json
	sourceDataFlag  bool
	sourceMeta      string // 元数据 json
	sourceMetaFlag  bool
	version         string // 版本号(自增)
	versionFlag     bool
	sourceLink      string // 链接
	sourceLinkFlag  bool
	summary         string // 总括
	summaryFlag     bool
	preview         string // 预览
	previewFlag     bool
	i18nSummay      string // 综述 json
	i18nSummayFlag  bool
	i18nPreview     string // 预览 json
	i18nPreviewFlag bool
	owner           string // 所有者
	ownerFlag       bool
	extra           string // 扩展字段 json
	extraFlag       bool
}

func NewEntityBuilder() *EntityBuilder {
	builder := &EntityBuilder{}
	return builder
}

// block 唯一标识
//
// 示例值：7794641623571830467
func (builder *EntityBuilder) BlockId(blockId string) *EntityBuilder {
	builder.blockId = blockId
	builder.blockIdFlag = true
	return builder
}

// 标题
//
// 示例值：已阅 block
func (builder *EntityBuilder) Title(title string) *EntityBuilder {
	builder.title = title
	builder.titleFlag = true
	return builder
}

// block 类型ID
//
// 示例值：blk_6204893fee000013739f5359
func (builder *EntityBuilder) BlockTypeId(blockTypeId string) *EntityBuilder {
	builder.blockTypeId = blockTypeId
	builder.blockTypeIdFlag = true
	return builder
}

// 业务数据 json
//
// 示例值：{"data":"业务数据"}
func (builder *EntityBuilder) SourceData(sourceData string) *EntityBuilder {
	builder.sourceData = sourceData
	builder.sourceDataFlag = true
	return builder
}

// 元数据 json
//
// 示例值：{"bizId":"7094067849152430100"}
func (builder *EntityBuilder) SourceMeta(sourceMeta string) *EntityBuilder {
	builder.sourceMeta = sourceMeta
	builder.sourceMetaFlag = true
	return builder
}

// 版本号(自增)
//
// 示例值：1651716489253602
func (builder *EntityBuilder) Version(version string) *EntityBuilder {
	builder.version = version
	builder.versionFlag = true
	return builder
}

// 链接
//
// 示例值：{}
func (builder *EntityBuilder) SourceLink(sourceLink string) *EntityBuilder {
	builder.sourceLink = sourceLink
	builder.sourceLinkFlag = true
	return builder
}

// 总括
//
// 示例值：{"cn":"这是一个block"}
func (builder *EntityBuilder) Summary(summary string) *EntityBuilder {
	builder.summary = summary
	builder.summaryFlag = true
	return builder
}

// 预览
//
// 示例值：{"cn":"这是一个block"}
func (builder *EntityBuilder) Preview(preview string) *EntityBuilder {
	builder.preview = preview
	builder.previewFlag = true
	return builder
}

// 综述 json
//
// 示例值：{"cn":"这是一个block","va": "this is a block"}
func (builder *EntityBuilder) I18nSummay(i18nSummay string) *EntityBuilder {
	builder.i18nSummay = i18nSummay
	builder.i18nSummayFlag = true
	return builder
}

// 预览 json
//
// 示例值：{"cn":"这是一个block","va": "this is a block"}
func (builder *EntityBuilder) I18nPreview(i18nPreview string) *EntityBuilder {
	builder.i18nPreview = i18nPreview
	builder.i18nPreviewFlag = true
	return builder
}

// 所有者
//
// 示例值：ou_fa7aa170f92d1615de63371ac425a767
func (builder *EntityBuilder) Owner(owner string) *EntityBuilder {
	builder.owner = owner
	builder.ownerFlag = true
	return builder
}

// 扩展字段 json
//
// 示例值：{}
func (builder *EntityBuilder) Extra(extra string) *EntityBuilder {
	builder.extra = extra
	builder.extraFlag = true
	return builder
}

func (builder *EntityBuilder) Build() *Entity {
	req := &Entity{}
	if builder.blockIdFlag {
		req.BlockId = &builder.blockId

	}
	if builder.titleFlag {
		req.Title = &builder.title

	}
	if builder.blockTypeIdFlag {
		req.BlockTypeId = &builder.blockTypeId

	}
	if builder.sourceDataFlag {
		req.SourceData = &builder.sourceData

	}
	if builder.sourceMetaFlag {
		req.SourceMeta = &builder.sourceMeta

	}
	if builder.versionFlag {
		req.Version = &builder.version

	}
	if builder.sourceLinkFlag {
		req.SourceLink = &builder.sourceLink

	}
	if builder.summaryFlag {
		req.Summary = &builder.summary

	}
	if builder.previewFlag {
		req.Preview = &builder.preview

	}
	if builder.i18nSummayFlag {
		req.I18nSummay = &builder.i18nSummay

	}
	if builder.i18nPreviewFlag {
		req.I18nPreview = &builder.i18nPreview

	}
	if builder.ownerFlag {
		req.Owner = &builder.owner

	}
	if builder.extraFlag {
		req.Extra = &builder.extra

	}
	return req
}

type Message struct {
	Body     *string  `json:"body,omitempty"`     // 协同数据内容
	Version  *string  `json:"version,omitempty"`  // 版本号(自增)
	BlockId  *string  `json:"block_id,omitempty"` // entity实体ID
	Resource *string  `json:"resource,omitempty"` // 业务来源
	OpenIds  []string `json:"open_ids,omitempty"` // 推送用户列表
}

type MessageBuilder struct {
	body         string // 协同数据内容
	bodyFlag     bool
	version      string // 版本号(自增)
	versionFlag  bool
	blockId      string // entity实体ID
	blockIdFlag  bool
	resource     string // 业务来源
	resourceFlag bool
	openIds      []string // 推送用户列表
	openIdsFlag  bool
}

func NewMessageBuilder() *MessageBuilder {
	builder := &MessageBuilder{}
	return builder
}

// 协同数据内容
//
// 示例值：{"id":"7094066727704592403","token":"test_123456789"}
func (builder *MessageBuilder) Body(body string) *MessageBuilder {
	builder.body = body
	builder.bodyFlag = true
	return builder
}

// 版本号(自增)
//
// 示例值：1637565292196
func (builder *MessageBuilder) Version(version string) *MessageBuilder {
	builder.version = version
	builder.versionFlag = true
	return builder
}

// entity实体ID
//
// 示例值：8116040162664047375
func (builder *MessageBuilder) BlockId(blockId string) *MessageBuilder {
	builder.blockId = blockId
	builder.blockIdFlag = true
	return builder
}

// 业务来源
//
// 示例值：read_block
func (builder *MessageBuilder) Resource(resource string) *MessageBuilder {
	builder.resource = resource
	builder.resourceFlag = true
	return builder
}

// 推送用户列表
//
// 示例值：["ou_fa7aa170f92d1615de63371ac425a767"]
func (builder *MessageBuilder) OpenIds(openIds []string) *MessageBuilder {
	builder.openIds = openIds
	builder.openIdsFlag = true
	return builder
}

func (builder *MessageBuilder) Build() *Message {
	req := &Message{}
	if builder.bodyFlag {
		req.Body = &builder.body

	}
	if builder.versionFlag {
		req.Version = &builder.version

	}
	if builder.blockIdFlag {
		req.BlockId = &builder.blockId

	}
	if builder.resourceFlag {
		req.Resource = &builder.resource

	}
	if builder.openIdsFlag {
		req.OpenIds = builder.openIds
	}
	return req
}

type CreateEntityReqBodyBuilder struct {
	title           string // 标题
	titleFlag       bool
	blockTypeId     string // block 类型ID
	blockTypeIdFlag bool
	sourceData      string // 内容
	sourceDataFlag  bool
	sourceMeta      string // 元数据
	sourceMetaFlag  bool
	version         string // 版本号(自增值)
	versionFlag     bool
	sourceLink      string // block原链接
	sourceLinkFlag  bool
	owner           string // 所有者
	ownerFlag       bool
	extra           string // 扩展字段
	extraFlag       bool
	i18nSummary     string // 国际化概括
	i18nSummaryFlag bool
	i18nPreview     string // 国际化预览
	i18nPreviewFlag bool
	summary         string // 概括
	summaryFlag     bool
	preview         string // 预览
	previewFlag     bool
}

func NewCreateEntityReqBodyBuilder() *CreateEntityReqBodyBuilder {
	builder := &CreateEntityReqBodyBuilder{}
	return builder
}

// 标题
//
//示例值：已阅block
func (builder *CreateEntityReqBodyBuilder) Title(title string) *CreateEntityReqBodyBuilder {
	builder.title = title
	builder.titleFlag = true
	return builder
}

// block 类型ID
//
//示例值：blk_614c1c952f800014b27f87d6
func (builder *CreateEntityReqBodyBuilder) BlockTypeId(blockTypeId string) *CreateEntityReqBodyBuilder {
	builder.blockTypeId = blockTypeId
	builder.blockTypeIdFlag = true
	return builder
}

// 内容
//
//示例值：{"data":"业务数据"}
func (builder *CreateEntityReqBodyBuilder) SourceData(sourceData string) *CreateEntityReqBodyBuilder {
	builder.sourceData = sourceData
	builder.sourceDataFlag = true
	return builder
}

// 元数据
//
//示例值：{"id":7090084015725608979}
func (builder *CreateEntityReqBodyBuilder) SourceMeta(sourceMeta string) *CreateEntityReqBodyBuilder {
	builder.sourceMeta = sourceMeta
	builder.sourceMetaFlag = true
	return builder
}

// 版本号(自增值)
//
//示例值：1
func (builder *CreateEntityReqBodyBuilder) Version(version string) *CreateEntityReqBodyBuilder {
	builder.version = version
	builder.versionFlag = true
	return builder
}

// block原链接
//
//示例值：{"_data":"https://docs.feishu.cn/block/78","_version":1}
func (builder *CreateEntityReqBodyBuilder) SourceLink(sourceLink string) *CreateEntityReqBodyBuilder {
	builder.sourceLink = sourceLink
	builder.sourceLinkFlag = true
	return builder
}

// 所有者
//
//示例值：ou_fa7aa170f92d1615de63371ac425a767
func (builder *CreateEntityReqBodyBuilder) Owner(owner string) *CreateEntityReqBodyBuilder {
	builder.owner = owner
	builder.ownerFlag = true
	return builder
}

// 扩展字段
//
//示例值：{}
func (builder *CreateEntityReqBodyBuilder) Extra(extra string) *CreateEntityReqBodyBuilder {
	builder.extra = extra
	builder.extraFlag = true
	return builder
}

// 国际化概括
//
//示例值：{"cn":"这是一个block","va": "this is a block"}
func (builder *CreateEntityReqBodyBuilder) I18nSummary(i18nSummary string) *CreateEntityReqBodyBuilder {
	builder.i18nSummary = i18nSummary
	builder.i18nSummaryFlag = true
	return builder
}

// 国际化预览
//
//示例值：{"cn":"这是一个block","va": "this is a block"}
func (builder *CreateEntityReqBodyBuilder) I18nPreview(i18nPreview string) *CreateEntityReqBodyBuilder {
	builder.i18nPreview = i18nPreview
	builder.i18nPreviewFlag = true
	return builder
}

// 概括
//
//示例值：{"cn":"这是一个block"}
func (builder *CreateEntityReqBodyBuilder) Summary(summary string) *CreateEntityReqBodyBuilder {
	builder.summary = summary
	builder.summaryFlag = true
	return builder
}

// 预览
//
//示例值：{"cn":"这是一个block"}
func (builder *CreateEntityReqBodyBuilder) Preview(preview string) *CreateEntityReqBodyBuilder {
	builder.preview = preview
	builder.previewFlag = true
	return builder
}

func (builder *CreateEntityReqBodyBuilder) Build() *CreateEntityReqBody {
	req := &CreateEntityReqBody{}
	if builder.titleFlag {
		req.Title = &builder.title
	}
	if builder.blockTypeIdFlag {
		req.BlockTypeId = &builder.blockTypeId
	}
	if builder.sourceDataFlag {
		req.SourceData = &builder.sourceData
	}
	if builder.sourceMetaFlag {
		req.SourceMeta = &builder.sourceMeta
	}
	if builder.versionFlag {
		req.Version = &builder.version
	}
	if builder.sourceLinkFlag {
		req.SourceLink = &builder.sourceLink
	}
	if builder.ownerFlag {
		req.Owner = &builder.owner
	}
	if builder.extraFlag {
		req.Extra = &builder.extra
	}
	if builder.i18nSummaryFlag {
		req.I18nSummary = &builder.i18nSummary
	}
	if builder.i18nPreviewFlag {
		req.I18nPreview = &builder.i18nPreview
	}
	if builder.summaryFlag {
		req.Summary = &builder.summary
	}
	if builder.previewFlag {
		req.Preview = &builder.preview
	}
	return req
}

type CreateEntityPathReqBodyBuilder struct {
	title           string // 标题
	titleFlag       bool
	blockTypeId     string // block 类型ID
	blockTypeIdFlag bool
	sourceData      string // 内容
	sourceDataFlag  bool
	sourceMeta      string // 元数据
	sourceMetaFlag  bool
	version         string // 版本号(自增值)
	versionFlag     bool
	sourceLink      string // block原链接
	sourceLinkFlag  bool
	owner           string // 所有者
	ownerFlag       bool
	extra           string // 扩展字段
	extraFlag       bool
	i18nSummary     string // 国际化概括
	i18nSummaryFlag bool
	i18nPreview     string // 国际化预览
	i18nPreviewFlag bool
	summary         string // 概括
	summaryFlag     bool
	preview         string // 预览
	previewFlag     bool
}

func NewCreateEntityPathReqBodyBuilder() *CreateEntityPathReqBodyBuilder {
	builder := &CreateEntityPathReqBodyBuilder{}
	return builder
}

// 标题
//
// 示例值：已阅block
func (builder *CreateEntityPathReqBodyBuilder) Title(title string) *CreateEntityPathReqBodyBuilder {
	builder.title = title
	builder.titleFlag = true
	return builder
}

// block 类型ID
//
// 示例值：blk_614c1c952f800014b27f87d6
func (builder *CreateEntityPathReqBodyBuilder) BlockTypeId(blockTypeId string) *CreateEntityPathReqBodyBuilder {
	builder.blockTypeId = blockTypeId
	builder.blockTypeIdFlag = true
	return builder
}

// 内容
//
// 示例值：{"data":"业务数据"}
func (builder *CreateEntityPathReqBodyBuilder) SourceData(sourceData string) *CreateEntityPathReqBodyBuilder {
	builder.sourceData = sourceData
	builder.sourceDataFlag = true
	return builder
}

// 元数据
//
// 示例值：{"id":7090084015725608979}
func (builder *CreateEntityPathReqBodyBuilder) SourceMeta(sourceMeta string) *CreateEntityPathReqBodyBuilder {
	builder.sourceMeta = sourceMeta
	builder.sourceMetaFlag = true
	return builder
}

// 版本号(自增值)
//
// 示例值：1
func (builder *CreateEntityPathReqBodyBuilder) Version(version string) *CreateEntityPathReqBodyBuilder {
	builder.version = version
	builder.versionFlag = true
	return builder
}

// block原链接
//
// 示例值：{"_data":"https://docs.feishu.cn/block/78","_version":1}
func (builder *CreateEntityPathReqBodyBuilder) SourceLink(sourceLink string) *CreateEntityPathReqBodyBuilder {
	builder.sourceLink = sourceLink
	builder.sourceLinkFlag = true
	return builder
}

// 所有者
//
// 示例值：ou_fa7aa170f92d1615de63371ac425a767
func (builder *CreateEntityPathReqBodyBuilder) Owner(owner string) *CreateEntityPathReqBodyBuilder {
	builder.owner = owner
	builder.ownerFlag = true
	return builder
}

// 扩展字段
//
// 示例值：{}
func (builder *CreateEntityPathReqBodyBuilder) Extra(extra string) *CreateEntityPathReqBodyBuilder {
	builder.extra = extra
	builder.extraFlag = true
	return builder
}

// 国际化概括
//
// 示例值：{"cn":"这是一个block","va": "this is a block"}
func (builder *CreateEntityPathReqBodyBuilder) I18nSummary(i18nSummary string) *CreateEntityPathReqBodyBuilder {
	builder.i18nSummary = i18nSummary
	builder.i18nSummaryFlag = true
	return builder
}

// 国际化预览
//
// 示例值：{"cn":"这是一个block","va": "this is a block"}
func (builder *CreateEntityPathReqBodyBuilder) I18nPreview(i18nPreview string) *CreateEntityPathReqBodyBuilder {
	builder.i18nPreview = i18nPreview
	builder.i18nPreviewFlag = true
	return builder
}

// 概括
//
// 示例值：{"cn":"这是一个block"}
func (builder *CreateEntityPathReqBodyBuilder) Summary(summary string) *CreateEntityPathReqBodyBuilder {
	builder.summary = summary
	builder.summaryFlag = true
	return builder
}

// 预览
//
// 示例值：{"cn":"这是一个block"}
func (builder *CreateEntityPathReqBodyBuilder) Preview(preview string) *CreateEntityPathReqBodyBuilder {
	builder.preview = preview
	builder.previewFlag = true
	return builder
}

func (builder *CreateEntityPathReqBodyBuilder) Build() (*CreateEntityReqBody, error) {
	req := &CreateEntityReqBody{}
	if builder.titleFlag {
		req.Title = &builder.title
	}
	if builder.blockTypeIdFlag {
		req.BlockTypeId = &builder.blockTypeId
	}
	if builder.sourceDataFlag {
		req.SourceData = &builder.sourceData
	}
	if builder.sourceMetaFlag {
		req.SourceMeta = &builder.sourceMeta
	}
	if builder.versionFlag {
		req.Version = &builder.version
	}
	if builder.sourceLinkFlag {
		req.SourceLink = &builder.sourceLink
	}
	if builder.ownerFlag {
		req.Owner = &builder.owner
	}
	if builder.extraFlag {
		req.Extra = &builder.extra
	}
	if builder.i18nSummaryFlag {
		req.I18nSummary = &builder.i18nSummary
	}
	if builder.i18nPreviewFlag {
		req.I18nPreview = &builder.i18nPreview
	}
	if builder.summaryFlag {
		req.Summary = &builder.summary
	}
	if builder.previewFlag {
		req.Preview = &builder.preview
	}
	return req, nil
}

type CreateEntityReqBuilder struct {
	apiReq *larkcore.ApiReq
	body   *CreateEntityReqBody
}

func NewCreateEntityReqBuilder() *CreateEntityReqBuilder {
	builder := &CreateEntityReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 开发者可以通过该接口将部分或全部数据存放于 BlockEntity。
func (builder *CreateEntityReqBuilder) Body(body *CreateEntityReqBody) *CreateEntityReqBuilder {
	builder.body = body
	return builder
}

func (builder *CreateEntityReqBuilder) Build() *CreateEntityReq {
	req := &CreateEntityReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.Body = builder.body
	return req
}

type CreateEntityReqBody struct {
	Title       *string `json:"title,omitempty"`         // 标题
	BlockTypeId *string `json:"block_type_id,omitempty"` // block 类型ID
	SourceData  *string `json:"source_data,omitempty"`   // 内容
	SourceMeta  *string `json:"source_meta,omitempty"`   // 元数据
	Version     *string `json:"version,omitempty"`       // 版本号(自增值)
	SourceLink  *string `json:"source_link,omitempty"`   // block原链接
	Owner       *string `json:"owner,omitempty"`         // 所有者
	Extra       *string `json:"extra,omitempty"`         // 扩展字段
	I18nSummary *string `json:"i18n_summary,omitempty"`  // 国际化概括
	I18nPreview *string `json:"i18n_preview,omitempty"`  // 国际化预览
	Summary     *string `json:"summary,omitempty"`       // 概括
	Preview     *string `json:"preview,omitempty"`       // 预览
}

type CreateEntityReq struct {
	apiReq *larkcore.ApiReq
	Body   *CreateEntityReqBody `body:""`
}

type CreateEntityRespData struct {
	Entity *Entity `json:"entity,omitempty"` // 返回对象实体
}

type CreateEntityResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *CreateEntityRespData `json:"data"` // 业务数据
}

func (resp *CreateEntityResp) Success() bool {
	return resp.Code == 0
}

type UpdateEntityReqBuilder struct {
	apiReq *larkcore.ApiReq
	entity *Entity
}

func NewUpdateEntityReqBuilder() *UpdateEntityReqBuilder {
	builder := &UpdateEntityReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// block唯一标识
//
// 示例值：7794641623571830467
func (builder *UpdateEntityReqBuilder) BlockId(blockId string) *UpdateEntityReqBuilder {
	builder.apiReq.PathParams.Set("block_id", fmt.Sprint(blockId))
	return builder
}

// 开发者通过该接口可以更新存储在BlockEntity中的数据，并实时推送到端侧。
func (builder *UpdateEntityReqBuilder) Entity(entity *Entity) *UpdateEntityReqBuilder {
	builder.entity = entity
	return builder
}

func (builder *UpdateEntityReqBuilder) Build() *UpdateEntityReq {
	req := &UpdateEntityReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.PathParams = builder.apiReq.PathParams
	req.apiReq.Body = builder.entity
	return req
}

type UpdateEntityReq struct {
	apiReq *larkcore.ApiReq
	Entity *Entity `body:""`
}

type UpdateEntityResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
}

func (resp *UpdateEntityResp) Success() bool {
	return resp.Code == 0
}

type CreateMessageReqBuilder struct {
	apiReq  *larkcore.ApiReq
	message *Message
}

func NewCreateMessageReqBuilder() *CreateMessageReqBuilder {
	builder := &CreateMessageReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 根据BlockID向指定用户列表推送协同数据。
func (builder *CreateMessageReqBuilder) Message(message *Message) *CreateMessageReqBuilder {
	builder.message = message
	return builder
}

func (builder *CreateMessageReqBuilder) Build() *CreateMessageReq {
	req := &CreateMessageReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.Body = builder.message
	return req
}

type CreateMessageReq struct {
	apiReq  *larkcore.ApiReq
	Message *Message `body:""`
}

type CreateMessageResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
}

func (resp *CreateMessageResp) Success() bool {
	return resp.Code == 0
}
