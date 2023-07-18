/*
* @desc:部门管理参数
* @company:舒可到家健康管理有限公司
* @Author: pmf<291641454@qq.com>
* @Date:   2022/4/6 15:07
 */

package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/skdj/model/entity"
)

// ServiceSearchReq 服务项目列表
type ServiceSearchReq struct {
	g.Meta      `path:"/service/list" tags:"服务项目管理" method:"get" summary:"列表"`
	ServiceName string `p:"serviceName" dc:"服务项目名称"`
	Status      string `p:"status" dc:"项目状态"`
	Page        int    `json:"page" d:"1" dc:"第几页，1开始"`
}
type ServiceSearchRes struct {
	g.Meta      `mime:"application/json"`
	ServiceList []*entity.TService `json:"service_list"`
	Total       int                `json:"total"`
}

// ServiceAddReq 添加服务项目
type ServiceAddReq struct {
	g.Meta   `path:"/service/add" tags:"服务项目管理" method:"post" summary:"添加部门"`
	ParentID int    `p:"parentId"  v:"required#父级不能为空"`
	DeptName string `p:"deptName"  v:"required#部门名称不能为空"`
	OrderNum int    `p:"orderNum"  v:"required#排序不能为空"`
	Leader   string `p:"leader"`
	Phone    string `p:"phone"`
	Email    string `p:"email"  v:"email#邮箱格式不正确"`
	Status   uint   `p:"status"  v:"required#状态必须"`
}
type ServiceAddRes struct {
}

// ServiceEditReq 编辑服务项目
type ServiceEditReq struct {
	g.Meta `path:"/service/edit" tags:"服务项目管理" method:"put" summary:"编辑"`
}
type ServiceEditRes struct {
}

// ServiceDeleteReq 删除服务项目
type ServiceDeleteReq struct {
	g.Meta `path:"/service/delete" tags:"服务项目管理" method:"delete" summary:"删除部门"`
	Id     uint64 `p:"id" v:"required#id不能为空"`
}

type ServiceDeleteRes struct {
}
