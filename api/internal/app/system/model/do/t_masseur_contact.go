// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TMasseurContact is the golang structure of table t_masseur_contact for DAO operations like Where/Data.
type TMasseurContact struct {
	g.Meta       `orm:"table:t_masseur_contact, do:true"`
	Id           interface{} // 唯一自增ID
	MasseurId    interface{} //
	Name         interface{} // 姓名
	Phone        interface{} // 手机号
	Enabled      interface{} //
	CreatedTime  *gtime.Time // 创建时间
	ModifiedTime *gtime.Time // 更新时间
}