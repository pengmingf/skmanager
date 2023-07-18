// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TMasseurService is the golang structure of table t_masseur_service for DAO operations like Where/Data.
type TMasseurService struct {
	g.Meta       `orm:"table:t_masseur_service, do:true"`
	Id           interface{} // 唯一自增ID
	ServiceId    interface{} // 服务id
	MasseurId    interface{} // 技师id
	Status       interface{} // 状态 0软删除 1使用
	CreatedTime  *gtime.Time // 创建时间
	ModifiedTime *gtime.Time // 更新时间
}
