// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TMasseurPicExamine is the golang structure for table t_masseur_pic_examine.
type TMasseurPicExamine struct {
	Id           int         `json:"id"           description:"唯一自增ID"`
	MasseurId    string      `json:"masseurId"    description:""`
	Pic          string      `json:"pic"          description:"图片"`
	PicType      string      `json:"picType"      description:"图片类型,profession职业照，life生活照"`
	Approval     int         `json:"approval"     description:"0待审批1审批通过2审批不通过"`
	ApprovalUser string      `json:"approvalUser" description:"审批人"`
	Enabled      int         `json:"enabled"      description:""`
	CreatedTime  *gtime.Time `json:"createdTime"  description:"创建时间"`
	ModifiedTime *gtime.Time `json:"modifiedTime" description:"更新时间"`
}
