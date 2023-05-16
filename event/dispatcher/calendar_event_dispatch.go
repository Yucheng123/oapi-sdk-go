// Package dispatcher code generated by oapi sdk gen
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

package dispatcher

import (
	"context"
	"github.com/Yucheng123/oapi-sdk-go/v3/service/calendar/v4"
)

// 日历变更
//
// - 当订阅用户的日历列表有日历变动时触发此事件。
//
// - 应用首先需要调用上述接口建立订阅关系。应用收到该事件后，使用事件的 user_list 字段中的用户对应的 user_access_token 调用[获取日历列表](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/list)接口拉取增量的变更数据
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/events/changed
func (dispatcher *EventDispatcher) OnP2CalendarChangedV4(handler func(ctx context.Context, event *larkcalendar.P2CalendarChangedV4) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["calendar.calendar.changed_v4"]
	if existed {
		panic("event: multiple handler registrations for " + "calendar.calendar.changed_v4")
	}
	dispatcher.eventType2EventHandler["calendar.calendar.changed_v4"] = larkcalendar.NewP2CalendarChangedV4Handler(handler)
	return dispatcher
}

// ACL新建
//
// - 当被订阅的日历上有ACL被创建时触发此事件。
//
// - 特殊说明：应用首先需要调用上述接口建立订阅关系。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/events/created
func (dispatcher *EventDispatcher) OnP2CalendarAclCreatedV4(handler func(ctx context.Context, event *larkcalendar.P2CalendarAclCreatedV4) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["calendar.calendar.acl.created_v4"]
	if existed {
		panic("event: multiple handler registrations for " + "calendar.calendar.acl.created_v4")
	}
	dispatcher.eventType2EventHandler["calendar.calendar.acl.created_v4"] = larkcalendar.NewP2CalendarAclCreatedV4Handler(handler)
	return dispatcher
}

// ACL移除
//
// - 当被订阅的日历上有ACL被删除时触发此事件。
//
// - 特殊说明：应用首先需要调用上述接口建立订阅关系。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/events/deleted
func (dispatcher *EventDispatcher) OnP2CalendarAclDeletedV4(handler func(ctx context.Context, event *larkcalendar.P2CalendarAclDeletedV4) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["calendar.calendar.acl.deleted_v4"]
	if existed {
		panic("event: multiple handler registrations for " + "calendar.calendar.acl.deleted_v4")
	}
	dispatcher.eventType2EventHandler["calendar.calendar.acl.deleted_v4"] = larkcalendar.NewP2CalendarAclDeletedV4Handler(handler)
	return dispatcher
}

// 日程变更
//
// - 当被订阅的用户日历下有日程变更时触发此事件。
//
// - 应用首先需要调用[订阅日程变更事件接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/subscription)建立订阅关系。应用收到该事件后，使用事件的 user_list 字段中的用户对应的 user_access_token 调用[获取日程列表](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/list)接口拉取事件中 calendar_id 字段对应的日历下的日程数据
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/events/changed
func (dispatcher *EventDispatcher) OnP2CalendarEventChangedV4(handler func(ctx context.Context, event *larkcalendar.P2CalendarEventChangedV4) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["calendar.calendar.event.changed_v4"]
	if existed {
		panic("event: multiple handler registrations for " + "calendar.calendar.event.changed_v4")
	}
	dispatcher.eventType2EventHandler["calendar.calendar.event.changed_v4"] = larkcalendar.NewP2CalendarEventChangedV4Handler(handler)
	return dispatcher
}
