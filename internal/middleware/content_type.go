package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

// Common Header
const (
	HeaderAccessToken = "Authorization" // access token
	HeaderDeviceId    = "Device-Id"     // 设备id
	HeaderRequestId   = "Request-Id"    // 请求id
)

// Header Client Type，添加新类型记得去 common_util/client/client_type 中的数组中引入
const (
	HeaderClientType = "Client-Type" // 客户端类型
	ClientTypeWeb    = "web"         // 电脑网站
	ClientTypeApp    = "app"         // 手机app
	ClientTypeMWeb   = "mweb"        // 手机网站
	ClientTypeIoT    = "iot"         // 物联网设备
)

// Header Content Type
const (
	HeaderContentType  = "Content-Type"                   // Content-Type
	ContentTypeDefault = "application/json;charset=utf-8" // json utf-8
)

// Header logic id
const (
	HeaderUserId   = "User-Id"   // 用户id
	HeaderUserName = "User-Name" // 用户名称
	HeaderOrgId    = "Org-Id"    // 机构id
	HeaderRoleId   = "Role-Id"   // 角色id
	HeaderRoleName = "Role-Name" // 角色名称
	HeaderRoleTyp  = "Role-Typ"  // 角色类型
)

// Header wechat id
const (
	HeaderOpenId = "Open-Id" // 用户id
)

// ContentType 设置 Content-Type
func ContentType(r *ghttp.Request) {
	r.Middleware.Next()

	r.Response.Header().Set(HeaderContentType, ContentTypeDefault)
}
