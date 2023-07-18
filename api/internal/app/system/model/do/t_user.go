// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TUser is the golang structure of table t_user for DAO operations like Where/Data.
type TUser struct {
	g.Meta        `orm:"table:t_user, do:true"`
	Id            interface{} // 唯一自增ID
	Openid        interface{} // 用户openid
	From          interface{} // 用户来源
	Nickname      interface{} // 昵称
	Sex           interface{} // 性别
	Province      interface{} // 省份
	City          interface{} // 城市
	Country       interface{} // 国家
	HeadImg       interface{} // 头像
	Balance       interface{} // 账户余额（分）
	Give          interface{} // 赠送金额（分）
	Privilege     interface{} // 用户特权信息
	UnionId       interface{} // 微信unionid
	ShareCode     interface{} // 分享码，唯一
	Phone         interface{} // 手机号
	LastLoginTime *gtime.Time // 最后登录时间
	CreatedTime   *gtime.Time // 创建时间
	ModifiedTime  *gtime.Time // 更新时间
}