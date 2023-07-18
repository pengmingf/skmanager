// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TMasseurBank is the golang structure for table t_masseur_bank.
type TMasseurBank struct {
	Id           int         `json:"id"           description:"唯一自增ID"`
	MasseurId    string      `json:"masseurId"    description:""`
	CardId       string      `json:"cardId"       description:"银行卡号"`
	CardName     string      `json:"cardName"     description:"持卡人姓名"`
	BankName     string      `json:"bankName"     description:"所属银行"`
	SubBankName  string      `json:"subBankName"  description:"分行"`
	Enabled      int         `json:"enabled"      description:""`
	CreatedTime  *gtime.Time `json:"createdTime"  description:"创建时间"`
	ModifiedTime *gtime.Time `json:"modifiedTime" description:"更新时间"`
}
