// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserRefund is the golang structure for table t_user_refund.
type TUserRefund struct {
	Id           int         `json:"id"           description:"唯一自增ID"`
	UserId       string      `json:"userId"       description:"用户openid"`
	Money        int         `json:"money"        description:"退款金额(分)"`
	Reason       string      `json:"reason"       description:"退款原因"`
	Operator     string      `json:"operator"     description:"操作人"`
	OperatorTime *gtime.Time `json:"operatorTime" description:"操作时间"`
	Status       string      `json:"status"       description:"退款状态"`
	Enabled      int         `json:"enabled"      description:""`
	CreatedTime  *gtime.Time `json:"createdTime"  description:"创建时间"`
	ModifiedTime *gtime.Time `json:"modifiedTime" description:"更新时间"`
}