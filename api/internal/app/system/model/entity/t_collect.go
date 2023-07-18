// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TCollect is the golang structure for table t_collect.
type TCollect struct {
	Id          int         `json:"id"          description:"唯一自增ID"`
	UserId      string      `json:"userId"      description:"用户id"`
	MasseurId   string      `json:"masseurId"   description:"技师id"`
	CreatedTime *gtime.Time `json:"createdTime" description:"创建时间"`
}