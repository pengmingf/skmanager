// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TOrderProcess is the golang structure of table t_order_process for DAO operations like Where/Data.
type TOrderProcess struct {
	g.Meta      `orm:"table:t_order_process, do:true"`
	Id          interface{} // 唯一自增ID
	OrderId     interface{} // 订单id
	ProcessType interface{} // 订单流程
	Enabled     interface{} //
	CreatedTime *gtime.Time // 创建时间
}
