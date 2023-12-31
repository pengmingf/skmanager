// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserTopUp is the golang structure for table t_user_top_up.
type TUserTopUp struct {
	Id            int         `json:"id"            description:"唯一自增ID"`
	UserId        string      `json:"userId"        description:"用户openid"`
	UpId          string      `json:"upId"          description:"充值订单号"`
	Status        string      `json:"status"        description:"状态，已完成finish，失败fail"`
	BdMasseurId   int         `json:"bdMasseurId"   description:"绑定的技师ID,自增ID"`
	BdMasseurName string      `json:"bdMasseurName" description:"绑定的技师名称"`
	Give          int         `json:"give"          description:"赠送的金额"`
	Money         int         `json:"money"         description:"总金额，包含赠送的(分)"`
	Enabled       int         `json:"enabled"       description:""`
	CreatedTime   *gtime.Time `json:"createdTime"   description:"创建时间"`
	ModifiedTime  *gtime.Time `json:"modifiedTime"  description:"更新时间"`
}
