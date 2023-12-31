// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TOrderBell is the golang structure of table t_order_bell for DAO operations like Where/Data.
type TOrderBell struct {
	g.Meta       `orm:"table:t_order_bell, do:true"`
	Id           interface{} // 唯一自增ID
	OrderId      interface{} // 主订单id
	SubOrderId   interface{} // 订单id
	OldServiceId interface{} // 老项目id,升级单专用
	ServiceMoney interface{} // 项目总价
	PayType      interface{} // 支付方式0余额支付1汇付宝支付
	PayMoney     interface{} // 订单金额（分）
	OrderType    interface{} // 订单状态,up为升级单，add为加钟单
	Status       interface{} // 订单状态
	Enabled      interface{} //
	CreatedTime  *gtime.Time // 创建时间
	ModifiedTime *gtime.Time // 更新时间
}
