package controller

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/skdj"
	"github.com/tiger1103/gfast/v3/internal/app/skdj/service"
)

var Service = serviceController{}

type serviceController struct {
	BaseController
}

// List 部门列表
func (c *serviceController) List(ctx context.Context, req *system.ServiceSearchReq) (res *system.ServiceSearchRes, err error) {
	res = new(system.ServiceSearchRes)
	res.ServiceList, err = service.Service().GetList(ctx, req)
	return
}

// // Add 添加部门
// func (c *serviceController) Add(ctx context.Context, req *system.DeptAddReq) (res *system.DeptAddRes, err error) {
// 	err = service.SysDept().Add(ctx, req)
// 	return
// }
//
// // Edit 修改部门
// func (c *serviceController) Edit(ctx context.Context, req *system.DeptEditReq) (res *system.DeptEditRes, err error) {
// 	err = service.SysDept().Edit(ctx, req)
// 	return
// }
//
// // Delete 删除部门
// func (c *serviceController) Delete(ctx context.Context, req *system.DeptDeleteReq) (res *system.DeptDeleteRes, err error) {
// 	err = service.SysDept().Delete(ctx, req.Id)
// 	return
// }
