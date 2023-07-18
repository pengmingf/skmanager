// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TOrderProcess is the golang structure for table t_order_process.
type TOrderProcess struct {
	Id          int         `json:"id"          description:"唯一自增ID"`
	OrderId     string      `json:"orderId"     description:"订单id"`
	ProcessType string      `json:"processType" description:"订单流程"`
	Enabled     int         `json:"enabled"     description:""`
	CreatedTime *gtime.Time `json:"createdTime" description:"创建时间"`
}
