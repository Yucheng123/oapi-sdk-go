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

// Package dispatcher code generated by oapi sdk gen
package dispatcher

import (
	"context"

	larkevent "github.com/Yucheng123/oapi-sdk-go/v3/event"
	larkapplication "github.com/Yucheng123/oapi-sdk-go/v3/service/application/v6"
	larkapproval "github.com/Yucheng123/oapi-sdk-go/v3/service/approval/v4"
	larkcontact "github.com/Yucheng123/oapi-sdk-go/v3/service/contact/v3"
	larkim "github.com/Yucheng123/oapi-sdk-go/v3/service/im/v1"
	larkmeeting_room "github.com/Yucheng123/oapi-sdk-go/v3/service/meeting_room/v1"
)

// v1消息协议：用户购买应用商店付费应用成功后发送给应用ISV的通知事件。
//
// 事件描述文档：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/event/public-app-purchase
func (dispatcher *EventDispatcher) OnP1OrderPaidV6(handler func(ctx context.Context, event *larkapplication.P1OrderPaidV6) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["order_paid"]
	if existed {
		panic("event: multiple handler registrations for " + "order_paid")
	}
	dispatcher.eventType2EventHandler["order_paid"] = larkapplication.NewP1OrderPaidV6Handler(handler)
	return dispatcher
}

// v1消息协议：企业解散后会推送此事件。商店应用开发者可在收到此事件后进行相应的账户注销、数据清理等处理。自建应用无此事件。
//
// 事件描述文档：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/event/app-uninstalled
func (dispatcher *EventDispatcher) OnP1AppUninstalledV6(handler func(ctx context.Context, event *larkapplication.P1AppUninstalledV6) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["app_uninstalled"]
	if existed {
		panic("event: multiple handler registrations for " + "app_uninstalled")
	}
	dispatcher.eventType2EventHandler["app_uninstalled"] = larkapplication.NewP1AppUninstalledV6Handler(handler)
	return dispatcher
}

// v1消息协议：当企业管理员在管理员后台启用、停用应用，或应用被平台停用时，开放平台推送 app_status_change 事件到请求网址。
//
// 事件描述文档：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/event/app-enabled-or-disabled
func (dispatcher *EventDispatcher) OnP1AppStatusChangedV6(handler func(ctx context.Context, event *larkapplication.P1AppStatusChangedV6) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["app_status_change"]
	if existed {
		panic("event: multiple handler registrations for " + "app_status_change")
	}
	dispatcher.eventType2EventHandler["app_status_change"] = larkapplication.NewP1AppStatusChangedV6Handler(handler)
	return dispatcher
}

// v1消息协议：用户阅读机器人发送的消息后触发此事件。
//
// 事件描述文档：https://open.feishu.cn/document/ukTMukTMukTM/ugzMugzMugzM/event/message-read
func (dispatcher *EventDispatcher) OnP1MessageReadV1(handler func(ctx context.Context, event *larkim.P1MessageReadV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["message_read"]
	if existed {
		panic("event: multiple handler registrations for " + "message_read")
	}
	dispatcher.eventType2EventHandler["message_read"] = larkim.NewP1MessageReadV1Handler(handler)
	return dispatcher
}

// v1消息协议：当用户发送消息给机器人或在群聊中@机器人时触发此事件。
//
// 事件描述文档：https://open.feishu.cn/document/ukTMukTMukTM/ugzMugzMugzM/event/receive-message
func (dispatcher *EventDispatcher) OnP1MessageReceiveV1(handler func(ctx context.Context, event *larkim.P1MessageReceiveV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["message"]
	if existed {
		panic("event: multiple handler registrations for " + "message")
	}
	dispatcher.eventType2EventHandler["message"] = larkim.NewP1MessageReceiveV1Handler(handler)
	return dispatcher
}

// v1消息协议：当员工的激活、暂停账号/恢复账号、操作离职时会触发此事件。此事件不依赖于任何权限。
//
// 事件描述文档：https://open.feishu.cn/document/ukTMukTMukTM/uETNz4SM1MjLxUzM//event/user-status-changed
func (dispatcher *EventDispatcher) OnP1UserStatusChangedV3(handler func(ctx context.Context, event *larkcontact.P1UserStatusChangedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["user_status_change"]
	if existed {
		panic("event: multiple handler registrations for " + "user_status_change")
	}
	dispatcher.eventType2EventHandler["user_status_change"] = larkcontact.NewP1UserStatusChangedV3Handler(handler)
	return dispatcher
}

// v1消息协议：当员工加入企业（user_add）、离职（user_leave）、个人信息发生变化（user_update）时，推送此事件。
//
// 事件描述文档：https://open.feishu.cn/document/ukTMukTMukTM/uETNz4SM1MjLxUzM//event/employee-change
func (dispatcher *EventDispatcher) OnP1UserChangedV3(handler func(ctx context.Context, event *larkcontact.P1UserChangedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["user_add"]
	if existed {
		panic("event: multiple handler registrations for " + "user_add")
	}
	dispatcher.eventType2EventHandler["user_add"] = larkcontact.NewP1UserChangedV3Handler(handler)

	_, existed = dispatcher.eventType2EventHandler["user_leave"]
	if existed {
		panic("event: multiple handler registrations for " + "user_add")
	}
	dispatcher.eventType2EventHandler["user_leave"] = larkcontact.NewP1UserChangedV3Handler(handler)

	_, existed = dispatcher.eventType2EventHandler["user_update"]
	if existed {
		panic("event: multiple handler registrations for " + "user_update")
	}
	dispatcher.eventType2EventHandler["user_update"] = larkcontact.NewP1UserChangedV3Handler(handler)
	return dispatcher
}

// v1消息协议：当新建部门（dept_add）、删除部门（dept_delete）、修改部门（dept_update）时，推送此事件。
//
// 事件描述文档：https://open.feishu.cn/document/ukTMukTMukTM/uETNz4SM1MjLxUzM//event/department-update
func (dispatcher *EventDispatcher) OnP1DepartmentChangedV3(handler func(ctx context.Context, event *larkcontact.P1DepartmentChangedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["dept_add"]
	if existed {
		panic("event: multiple handler registrations for " + "dept_add")
	}
	dispatcher.eventType2EventHandler["dept_add"] = larkcontact.NewP1DepartmentChangedV3Handler(handler)

	_, existed = dispatcher.eventType2EventHandler["dept_delete"]
	if existed {
		panic("event: multiple handler registrations for " + "dept_delete")
	}
	dispatcher.eventType2EventHandler["dept_delete"] = larkcontact.NewP1DepartmentChangedV3Handler(handler)

	_, existed = dispatcher.eventType2EventHandler["dept_update"]
	if existed {
		panic("event: multiple handler registrations for " + "dept_update")
	}
	dispatcher.eventType2EventHandler["dept_update"] = larkcontact.NewP1DepartmentChangedV3Handler(handler)
	return dispatcher
}

// v1消息协议：当应用申请了 以应用身份访问通讯录 权限后，管理员可以配置应用的通讯录授权范围,当此范围变化时，就会触发授权范围变化事件。
//
// 事件描述文档：https://open.feishu.cn/document/ukTMukTMukTM/uETNz4SM1MjLxUzM//event/scope-change
func (dispatcher *EventDispatcher) OnP1ContactScopeChangedV3(handler func(ctx context.Context, event *larkcontact.P1ContactScopeChangedV3) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact_scope_change"]
	if existed {
		panic("event: multiple handler registrations for " + "contact_scope_change")
	}
	dispatcher.eventType2EventHandler["contact_scope_change"] = larkcontact.NewP1ContactScopeChangedV3Handler(handler)
	return dispatcher
}

// v1消息协议：机器人被邀请加入群聊时触发此事件。
//
// 事件描述文档：https://open.feishu.cn/document/ukTMukTMukTM/ugzMugzMugzM/event/bot-added-to-group
func (dispatcher *EventDispatcher) OnP1AddBotV1(handler func(ctx context.Context, event *larkim.P1AddBotV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["add_bot"]
	if existed {
		panic("event: multiple handler registrations for " + "add_bot")
	}
	dispatcher.eventType2EventHandler["add_bot"] = larkim.NewP1AddBotV1Handler(handler)
	return dispatcher
}

// v1消息协议：机器人被从群聊中移除时触发此事件。
//
// 事件描述文档：https://open.feishu.cn/document/ukTMukTMukTM/ugzMugzMugzM/event/bot-removed-from-group
func (dispatcher *EventDispatcher) OnP1RemoveAddBotV1(handler func(ctx context.Context, event *larkim.P1RemoveBotV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["remove_bot"]
	if existed {
		panic("event: multiple handler registrations for " + "remove_bot")
	}
	dispatcher.eventType2EventHandler["remove_bot"] = larkim.NewP1RemoveBotV1Handler(handler)
	return dispatcher
}

// v1消息协议：用户进群、出群后触发此事件
//
// 事件描述文档：https://open.feishu.cn/document/ukTMukTMukTM/uQDOwUjL0gDM14CN4ATN/event/user-joins-or-leave-group
func (dispatcher *EventDispatcher) OnP1UserInOutChatV1(handler func(ctx context.Context, event *larkim.P1UserInOutChatV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["add_user_to_chat"]
	if existed {
		panic("event: multiple handler registrations for " + "add_user_to_chat")
	}
	dispatcher.eventType2EventHandler["add_user_to_chat"] = larkim.NewP1UserInOutChatV1Handler(handler)

	_, existed = dispatcher.eventType2EventHandler["remove_user_from_chat"]
	if existed {
		panic("event: multiple handler registrations for " + "remove_user_from_chat")
	}
	dispatcher.eventType2EventHandler["remove_user_from_chat"] = larkim.NewP1UserInOutChatV1Handler(handler)

	_, existed = dispatcher.eventType2EventHandler["revoke_add_user_from_chat"]
	if existed {
		panic("event: multiple handler registrations for " + "revoke_add_user_from_chat")
	}
	dispatcher.eventType2EventHandler["revoke_add_user_from_chat"] = larkim.NewP1UserInOutChatV1Handler(handler)
	return dispatcher
}

// v1消息协议：群聊被解散后触发此事件。
//
// 事件描述文档：https://open.feishu.cn/document/ukTMukTMukTM/uQDOwUjL0gDM14CN4ATN/event/group-closed
func (dispatcher *EventDispatcher) OnP1ChatDisbandV1(handler func(ctx context.Context, event *larkim.P1ChatDisbandV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["chat_disband"]
	if existed {
		panic("event: multiple handler registrations for " + "chat_disband")
	}
	dispatcher.eventType2EventHandler["chat_disband"] = larkim.NewP1DisbandChatV1Handler(handler)
	return dispatcher
}

// v1消息协议：群配置修改后触发此事件。
//
// 事件描述文档：https://open.feishu.cn/document/ukTMukTMukTM/uQDOwUjL0gDM14CN4ATN/event/group-configuration-changes
func (dispatcher *EventDispatcher) OnP1GroupSettingUpdatedV1(handler func(ctx context.Context, event *larkim.P1GroupSettingUpdatedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["group_setting_update"]
	if existed {
		panic("event: multiple handler registrations for " + "group_setting_update")
	}
	dispatcher.eventType2EventHandler["group_setting_update"] = larkim.NewP1GroupSettingUpdatedV1Handler(handler)
	return dispatcher
}

// v1消息协议：当租户第一次安装并启用此应用时触发此事件，启用应用包含以下场景：
//  当租户管理员后台首次开通应用
//  租户内的普通成员首次安装此应用
//
// 事件描述文档：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/event/app-first-enabled
func (dispatcher *EventDispatcher) OnP1AppOpenV6(handler func(ctx context.Context, event *larkapplication.P1AppOpenV6) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["app_open"]
	if existed {
		panic("event: multiple handler registrations for " + "app_open")
	}
	dispatcher.eventType2EventHandler["app_open"] = larkapplication.NewP1AppOpenV6Handler(handler)
	return dispatcher
}

// v1消息协议：首次会话是用户了解应用的重要机会，你可以发送操作说明、配置地址来指导用户开始使用你的应用。
//
// 事件描述文档：https://open.feishu.cn/document/ukTMukTMukTM/uYDNxYjL2QTM24iN0EjN/bot-events
func (dispatcher *EventDispatcher) OnP1P2PChatCreatedV1(handler func(ctx context.Context, event *larkim.P1P2PChatCreatedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["p2p_chat_create"]
	if existed {
		panic("event: multiple handler registrations for " + "p2p_chat_create")
	}
	dispatcher.eventType2EventHandler["p2p_chat_create"] = larkim.NewP1P2PChatCreatedV1Handler(handler)
	return dispatcher
}

// v1消息协议：当添加了第三方会议室的日程发生变动时（创建/更新/删除）触发此事件。
//
// 事件描述文档：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/meeting_room-v1/event/third-room-event-changes
func (dispatcher *EventDispatcher) OnP1ThirdPartyMeetingRoomChangedV1(handler func(ctx context.Context, event *larkmeeting_room.P1ThirdPartyMeetingRoomChangedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["third_party_meeting_room_event_created"]
	if existed {
		panic("event: multiple handler registrations for " + "third_party_meeting_room_event_created")
	}
	dispatcher.eventType2EventHandler["third_party_meeting_room_event_created"] = larkmeeting_room.NewP1ThirdPartyMeetingRoomChangedV1Handler(handler)

	_, existed = dispatcher.eventType2EventHandler["third_party_meeting_room_event_updated"]
	if existed {
		panic("event: multiple handler registrations for " + "third_party_meeting_room_event_updated")
	}
	dispatcher.eventType2EventHandler["third_party_meeting_room_event_updated"] = larkmeeting_room.NewP1ThirdPartyMeetingRoomChangedV1Handler(handler)

	_, existed = dispatcher.eventType2EventHandler["third_party_meeting_room_event_deleted"]
	if existed {
		panic("event: multiple handler registrations for " + "third_party_meeting_room_event_deleted")
	}
	dispatcher.eventType2EventHandler["third_party_meeting_room_event_deleted"] = larkmeeting_room.NewP1ThirdPartyMeetingRoomChangedV1Handler(handler)
	return dispatcher
}

// v1消息协议：审批」应用的表单里如果包含 请假控件组，则在此表单审批通过后触发此事件。
//
// 事件描述文档：https://open.feishu.cn/document/ukTMukTMukTM/uIDO24iM4YjLygjN/event/leave
func (dispatcher *EventDispatcher) OnP1LeaveApprovalV4(handler func(ctx context.Context, event *larkapproval.P1LeaveApprovalV4) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["leave_approvalV2"]
	if existed {
		panic("event: multiple handler registrations for " + "leave_approvalV2")
	}
	dispatcher.eventType2EventHandler["leave_approvalV2"] = larkapproval.NewP1LeaveApprovalV4Handler(handler)
	return dispatcher
}

// v1消息协议：「审批」应用的表单里如果包含 加班控件组，则在此表单审批通过后触发此事件。
//
// 事件描述文档：https://open.feishu.cn/document/ukTMukTMukTM/uIDO24iM4YjLygjN/event/overtime
func (dispatcher *EventDispatcher) OnP1WorkApprovalV4(handler func(ctx context.Context, event *larkapproval.P1WorkApprovalV4) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["work_approval"]
	if existed {
		panic("event: multiple handler registrations for " + "work_approval")
	}
	dispatcher.eventType2EventHandler["work_approval"] = larkapproval.NewP1WorkApprovalV4Handler(handler)
	return dispatcher
}

// v1消息协议：换班申请审批通过后触发此事件。 你可以在「打卡」应用里提交换班申请。
//
// 事件描述文档：https://open.feishu.cn/document/ukTMukTMukTM/uIDO24iM4YjLygjN/event/shift-change
func (dispatcher *EventDispatcher) OnP1ShiftApprovalV4(handler func(ctx context.Context, event *larkapproval.P1ShiftApprovalV4) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["shift_approval"]
	if existed {
		panic("event: multiple handler registrations for " + "shift_approval")
	}
	dispatcher.eventType2EventHandler["shift_approval"] = larkapproval.NewP1ShiftApprovalV4Handler(handler)
	return dispatcher
}

// v1消息协议：补卡申请审批通过后触发此事件。 你可以在「打卡」应用里提交补卡申请。
//
// 事件描述文档：https://open.feishu.cn/document/ukTMukTMukTM/uIDO24iM4YjLygjN/event/attendance-record-correction
func (dispatcher *EventDispatcher) OnP1RemedyApprovalV4(handler func(ctx context.Context, event *larkapproval.P1RemedyApprovalV4) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["remedy_approval"]
	if existed {
		panic("event: multiple handler registrations for " + "remedy_approval")
	}
	dispatcher.eventType2EventHandler["remedy_approval"] = larkapproval.NewP1RemedyApprovalV4Handler(handler)
	return dispatcher
}

// v1消息协议：「审批」应用的表单里如果包含 出差控件组，则在此表单审批通过后触发此事件。
//
// 事件描述文档：https://open.feishu.cn/document/ukTMukTMukTM/uIDO24iM4YjLygjN/event/business-trip
func (dispatcher *EventDispatcher) OnP1TripApprovalV4(handler func(ctx context.Context, event *larkapproval.P1TripApprovalV4) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["trip_approval"]
	if existed {
		panic("event: multiple handler registrations for " + "trip_approval")
	}
	dispatcher.eventType2EventHandler["trip_approval"] = larkapproval.NewP1TripApprovalV4Handler(handler)
	return dispatcher
}

// v1消息协议：「审批」应用的表单里如果包含 外出控件组，则在此表单审批通过后触发此事件。
//
// 事件描述文档：https://open.feishu.cn/document/ukTMukTMukTM/uIDO24iM4YjLygjN/event/out-of-office
func (dispatcher *EventDispatcher) OnP1OutApprovalV4(handler func(ctx context.Context, event *larkapproval.P1OutApprovalV4) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["out_approval"]
	if existed {
		panic("event: multiple handler registrations for " + "out_approval")
	}
	dispatcher.eventType2EventHandler["out_approval"] = larkapproval.NewP1OutApprovalV4Handler(handler)
	return dispatcher
}

// 订阅事件扩展：开发者可自己传递事件类型，并传递对应事件类型的处理器
func (dispatcher *EventDispatcher) OnCustomizedEvent(eventType string, handler func(ctx context.Context, event *larkevent.EventReq) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler[eventType]
	if existed {
		panic("event: multiple handler registrations for " + eventType)
	}
	dispatcher.eventType2EventHandler[eventType] = &defaultHandler{handler: handler}
	return dispatcher
}

// 当 ISV 需要自己管理 app_ticket 和 token 时，需要注册该处理器,以获取app_ticket。这时 SDK 内将不能帮开发者自动获取和缓存token。
//
// 事件描述文档：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/event/app_ticket-events
func (dispatcher *EventDispatcher) OnAppTicketEvent(handler func(ctx context.Context, event *AppTicketEvent) error) *EventDispatcher {
	dispatcher.eventType2EventHandler["app_ticket"] = &CustomAppTicketEventHandler{handler: handler}
	return dispatcher
}
