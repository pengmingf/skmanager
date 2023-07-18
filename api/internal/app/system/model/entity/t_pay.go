// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TPay is the golang structure for table t_pay.
type TPay struct {
	Id           int         `json:"id"           description:"唯一自增ID"`
	PreId        string      `json:"preId"        description:"预支付id"`
	BusinessType string      `json:"businessType" description:"业务类型,order订单，ord_add加钟，ord_up升级，cz充值"`
	BusinessId   string      `json:"businessId"   description:"对应类型的业务id"`
	PayStatus    string      `json:"payStatus"    description:"支付状态 create创建订单，支付失败fail，支付成功success"`
	BackTime     *gtime.Time `json:"backTime"     description:"回调时间"`
	BackContent  string      `json:"backContent"  description:"回调内容"`
	BackType     string      `json:"backType"     description:"回调方式，search主动查询，callback三方回调"`
	Enabled      int         `json:"enabled"      description:""`
	CreatedTime  *gtime.Time `json:"createdTime"  description:"创建时间"`
	ModifiedTime *gtime.Time `json:"modifiedTime" description:"更新时间"`
}
