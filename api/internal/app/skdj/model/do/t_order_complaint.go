// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TOrderComplaint is the golang structure of table t_order_complaint for DAO operations like Where/Data.
type TOrderComplaint struct {
	g.Meta         `orm:"table:t_order_complaint, do:true"`
	Id             interface{} // 唯一自增ID
	OrderId        interface{} // 订单id
	MasseurId      interface{} // 服务人id
	UserId         interface{} // 用户id
	MasseurName    interface{} // 技师名称
	ComplaintLabel interface{} // 投诉类型
	Content        interface{} // 投诉内容
	Pic            interface{} // 上传凭证
	Phone          interface{} // 联系方式
	Status         interface{} // 状态
	CreatedTime    *gtime.Time // 创建时间
	ModifiedTime   *gtime.Time // 更新时间
}
