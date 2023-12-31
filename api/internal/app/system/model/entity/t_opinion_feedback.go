// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TOpinionFeedback is the golang structure for table t_opinion_feedback.
type TOpinionFeedback struct {
	Id           int         `json:"id"           description:"唯一自增ID"`
	UserId       string      `json:"userId"       description:"用户openid"`
	Content      string      `json:"content"      description:"反馈内容"`
	Pic          string      `json:"pic"          description:"图片"`
	Phone        string      `json:"phone"        description:"手机号"`
	Enabled      int         `json:"enabled"      description:""`
	CreatedTime  *gtime.Time `json:"createdTime"  description:"创建时间"`
	ModifiedTime *gtime.Time `json:"modifiedTime" description:"更新时间"`
}
