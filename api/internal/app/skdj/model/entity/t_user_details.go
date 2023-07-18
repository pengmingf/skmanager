// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserDetails is the golang structure for table t_user_details.
type TUserDetails struct {
	Id           int         `json:"id"           description:"唯一自增ID"`
	UserId       string      `json:"userId"       description:"用户openid"`
	OrderId      string      `json:"orderId"      description:"订单号"`
	Status       string      `json:"status"       description:"状态，已完成finish，失败fail"`
	Money        int         `json:"money"        description:"金额(分)"`
	Enabled      int         `json:"enabled"      description:""`
	CreatedTime  *gtime.Time `json:"createdTime"  description:"创建时间"`
	ModifiedTime *gtime.Time `json:"modifiedTime" description:"更新时间"`
}