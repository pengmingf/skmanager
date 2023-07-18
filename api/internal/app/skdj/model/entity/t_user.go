// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUser is the golang structure for table t_user.
type TUser struct {
	Id            int         `json:"id"            description:"唯一自增ID"`
	Openid        string      `json:"openid"        description:"用户openid"`
	From          string      `json:"from"          description:"用户来源"`
	Nickname      string      `json:"nickname"      description:"昵称"`
	Sex           string      `json:"sex"           description:"性别"`
	Province      string      `json:"province"      description:"省份"`
	City          string      `json:"city"          description:"城市"`
	Country       string      `json:"country"       description:"国家"`
	HeadImg       string      `json:"headImg"       description:"头像"`
	Balance       int         `json:"balance"       description:"账户余额（分）"`
	Give          int         `json:"give"          description:"赠送金额（分）"`
	Privilege     string      `json:"privilege"     description:"用户特权信息"`
	UnionId       string      `json:"unionId"       description:"微信unionid"`
	ShareCode     string      `json:"shareCode"     description:"分享码，唯一"`
	Phone         string      `json:"phone"         description:"手机号"`
	LastLoginTime *gtime.Time `json:"lastLoginTime" description:"最后登录时间"`
	CreatedTime   *gtime.Time `json:"createdTime"   description:"创建时间"`
	ModifiedTime  *gtime.Time `json:"modifiedTime"  description:"更新时间"`
}
