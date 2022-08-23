// Package translation code generated by oapi sdk gen
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

package larktranslation

import (
	"github.com/larksuite/oapi-sdk-go/v3/core"
)

type Term struct {
	From *string `json:"from,omitempty"` // 原文
	To   *string `json:"to,omitempty"`   // 译文
}

type TermBuilder struct {
	from     string // 原文
	fromFlag bool
	to       string // 译文
	toFlag   bool
}

func NewTermBuilder() *TermBuilder {
	builder := &TermBuilder{}
	return builder
}

// 原文
// 示例值：飞书
func (builder *TermBuilder) From(from string) *TermBuilder {
	builder.from = from
	builder.fromFlag = true
	return builder
}

// 译文
// 示例值：Lark
func (builder *TermBuilder) To(to string) *TermBuilder {
	builder.to = to
	builder.toFlag = true
	return builder
}

func (builder *TermBuilder) Build() *Term {
	req := &Term{}
	if builder.fromFlag {
		req.From = &builder.from

	}
	if builder.toFlag {
		req.To = &builder.to

	}
	return req
}

type Text struct {
}

type DetectTextReqBodyBuilder struct {
	text     string // 需要被识别语种的文本
	textFlag bool
}

func NewDetectTextReqBodyBuilder() *DetectTextReqBodyBuilder {
	builder := &DetectTextReqBodyBuilder{}
	return builder
}

// 需要被识别语种的文本
//
//示例值：你好
func (builder *DetectTextReqBodyBuilder) Text(text string) *DetectTextReqBodyBuilder {
	builder.text = text
	builder.textFlag = true
	return builder
}

func (builder *DetectTextReqBodyBuilder) Build() *DetectTextReqBody {
	req := &DetectTextReqBody{}
	if builder.textFlag {
		req.Text = &builder.text
	}
	return req
}

type DetectTextPathReqBodyBuilder struct {
	text     string // 需要被识别语种的文本
	textFlag bool
}

func NewDetectTextPathReqBodyBuilder() *DetectTextPathReqBodyBuilder {
	builder := &DetectTextPathReqBodyBuilder{}
	return builder
}

// 需要被识别语种的文本
//
// 示例值：你好
func (builder *DetectTextPathReqBodyBuilder) Text(text string) *DetectTextPathReqBodyBuilder {
	builder.text = text
	builder.textFlag = true
	return builder
}

func (builder *DetectTextPathReqBodyBuilder) Build() (*DetectTextReqBody, error) {
	req := &DetectTextReqBody{}
	if builder.textFlag {
		req.Text = &builder.text
	}
	return req, nil
}

type DetectTextReqBuilder struct {
	apiReq *larkcore.ApiReq
	body   *DetectTextReqBody
}

func NewDetectTextReqBuilder() *DetectTextReqBuilder {
	builder := &DetectTextReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 机器翻译 (MT)，支持 100 多种语言识别，返回符合 ISO 639-1 标准
func (builder *DetectTextReqBuilder) Body(body *DetectTextReqBody) *DetectTextReqBuilder {
	builder.body = body
	return builder
}

func (builder *DetectTextReqBuilder) Build() *DetectTextReq {
	req := &DetectTextReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.Body = builder.body
	return req
}

type DetectTextReqBody struct {
	Text *string `json:"text,omitempty"` // 需要被识别语种的文本
}

type DetectTextReq struct {
	apiReq *larkcore.ApiReq
	Body   *DetectTextReqBody `body:""`
}

type DetectTextRespData struct {
	Language *string `json:"language,omitempty"` // 识别的文本语种，返回符合 ISO 639-1 标准
}

type DetectTextResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *DetectTextRespData `json:"data"` // 业务数据
}

func (resp *DetectTextResp) Success() bool {
	return resp.Code == 0
}

type TranslateTextReqBodyBuilder struct {
	sourceLanguage     string // 源语言
	sourceLanguageFlag bool
	text               string // 源文本
	textFlag           bool
	targetLanguage     string // 目标语言
	targetLanguageFlag bool
	glossary           []*Term // 请求级术语表，携带术语，仅在本次翻译中生效（最多能携带 128个术语词）
	glossaryFlag       bool
}

func NewTranslateTextReqBodyBuilder() *TranslateTextReqBodyBuilder {
	builder := &TranslateTextReqBodyBuilder{}
	return builder
}

// 源语言
//
//示例值：zh
func (builder *TranslateTextReqBodyBuilder) SourceLanguage(sourceLanguage string) *TranslateTextReqBodyBuilder {
	builder.sourceLanguage = sourceLanguage
	builder.sourceLanguageFlag = true
	return builder
}

// 源文本
//
//示例值：尝试使用一下飞书吧
func (builder *TranslateTextReqBodyBuilder) Text(text string) *TranslateTextReqBodyBuilder {
	builder.text = text
	builder.textFlag = true
	return builder
}

// 目标语言
//
//示例值：en
func (builder *TranslateTextReqBodyBuilder) TargetLanguage(targetLanguage string) *TranslateTextReqBodyBuilder {
	builder.targetLanguage = targetLanguage
	builder.targetLanguageFlag = true
	return builder
}

// 请求级术语表，携带术语，仅在本次翻译中生效（最多能携带 128个术语词）
//
//示例值：
func (builder *TranslateTextReqBodyBuilder) Glossary(glossary []*Term) *TranslateTextReqBodyBuilder {
	builder.glossary = glossary
	builder.glossaryFlag = true
	return builder
}

func (builder *TranslateTextReqBodyBuilder) Build() *TranslateTextReqBody {
	req := &TranslateTextReqBody{}
	if builder.sourceLanguageFlag {
		req.SourceLanguage = &builder.sourceLanguage
	}
	if builder.textFlag {
		req.Text = &builder.text
	}
	if builder.targetLanguageFlag {
		req.TargetLanguage = &builder.targetLanguage
	}
	if builder.glossaryFlag {
		req.Glossary = builder.glossary
	}
	return req
}

type TranslateTextPathReqBodyBuilder struct {
	sourceLanguage     string // 源语言
	sourceLanguageFlag bool
	text               string // 源文本
	textFlag           bool
	targetLanguage     string // 目标语言
	targetLanguageFlag bool
	glossary           []*Term // 请求级术语表，携带术语，仅在本次翻译中生效（最多能携带 128个术语词）
	glossaryFlag       bool
}

func NewTranslateTextPathReqBodyBuilder() *TranslateTextPathReqBodyBuilder {
	builder := &TranslateTextPathReqBodyBuilder{}
	return builder
}

// 源语言
//
// 示例值：zh
func (builder *TranslateTextPathReqBodyBuilder) SourceLanguage(sourceLanguage string) *TranslateTextPathReqBodyBuilder {
	builder.sourceLanguage = sourceLanguage
	builder.sourceLanguageFlag = true
	return builder
}

// 源文本
//
// 示例值：尝试使用一下飞书吧
func (builder *TranslateTextPathReqBodyBuilder) Text(text string) *TranslateTextPathReqBodyBuilder {
	builder.text = text
	builder.textFlag = true
	return builder
}

// 目标语言
//
// 示例值：en
func (builder *TranslateTextPathReqBodyBuilder) TargetLanguage(targetLanguage string) *TranslateTextPathReqBodyBuilder {
	builder.targetLanguage = targetLanguage
	builder.targetLanguageFlag = true
	return builder
}

// 请求级术语表，携带术语，仅在本次翻译中生效（最多能携带 128个术语词）
//
// 示例值：
func (builder *TranslateTextPathReqBodyBuilder) Glossary(glossary []*Term) *TranslateTextPathReqBodyBuilder {
	builder.glossary = glossary
	builder.glossaryFlag = true
	return builder
}

func (builder *TranslateTextPathReqBodyBuilder) Build() (*TranslateTextReqBody, error) {
	req := &TranslateTextReqBody{}
	if builder.sourceLanguageFlag {
		req.SourceLanguage = &builder.sourceLanguage
	}
	if builder.textFlag {
		req.Text = &builder.text
	}
	if builder.targetLanguageFlag {
		req.TargetLanguage = &builder.targetLanguage
	}
	if builder.glossaryFlag {
		req.Glossary = builder.glossary
	}
	return req, nil
}

type TranslateTextReqBuilder struct {
	apiReq *larkcore.ApiReq
	body   *TranslateTextReqBody
}

func NewTranslateTextReqBuilder() *TranslateTextReqBuilder {
	builder := &TranslateTextReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 机器翻译 (MT)，支持以下语种互译：;"zh": 汉语；;"zh-Hant": 繁体汉语；;"en": 英语；;"ja": 日语；;"ru": 俄语；;"de": 德语；;"fr": 法语；;"it": 意大利语；;"pl": 波兰语；;"th": 泰语；;"hi": 印地语；;"id": 印尼语；;"es": 西班牙语；;"pt": 葡萄牙语；;"ko": 朝鲜语；;"vi": 越南语；
func (builder *TranslateTextReqBuilder) Body(body *TranslateTextReqBody) *TranslateTextReqBuilder {
	builder.body = body
	return builder
}

func (builder *TranslateTextReqBuilder) Build() *TranslateTextReq {
	req := &TranslateTextReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.Body = builder.body
	return req
}

type TranslateTextReqBody struct {
	SourceLanguage *string `json:"source_language,omitempty"` // 源语言
	Text           *string `json:"text,omitempty"`            // 源文本
	TargetLanguage *string `json:"target_language,omitempty"` // 目标语言
	Glossary       []*Term `json:"glossary,omitempty"`        // 请求级术语表，携带术语，仅在本次翻译中生效（最多能携带 128个术语词）
}

type TranslateTextReq struct {
	apiReq *larkcore.ApiReq
	Body   *TranslateTextReqBody `body:""`
}

type TranslateTextRespData struct {
	Text *string `json:"text,omitempty"` // 翻译后的文本
}

type TranslateTextResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *TranslateTextRespData `json:"data"` // 业务数据
}

func (resp *TranslateTextResp) Success() bool {
	return resp.Code == 0
}
