// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TOrderBell is the golang structure for table t_order_bell.
type TOrderBell struct {
	Id           int         `json:"id"           description:"唯一自增ID"`
	OrderId      string      `json:"orderId"      description:"主订单id"`
	SubOrderId   string      `json:"subOrderId"   description:"订单id"`
	OldServiceId string      `json:"oldServiceId" description:"老项目id,升级单专用"`
	ServiceMoney int         `json:"serviceMoney" description:"项目总价"`
	PayType      int         `json:"payType"      description:"支付方式0余额支付1汇付宝支付"`
	PayMoney     int         `json:"payMoney"     description:"订单金额（分）"`
	OrderType    string      `json:"orderType"    description:"订单状态,up为升级单，add为加钟单"`
	Status       string      `json:"status"       description:"订单状态"`
	Enabled      int         `json:"enabled"      description:""`
	CreatedTime  *gtime.Time `json:"createdTime"  description:"创建时间"`
	ModifiedTime *gtime.Time `json:"modifiedTime" description:"更新时间"`
}
