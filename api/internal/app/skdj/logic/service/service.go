package service

import (
	"context"

	system "github.com/tiger1103/gfast/v3/api/v1/skdj"
	"github.com/tiger1103/gfast/v3/internal/app/skdj/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/skdj/service"
)

func init() {
	service.RegisterService(New())
}

func New() *sService {
	return &sService{}
}

type sService struct {
}

func (s *sService) GetList(ctx context.Context, req *system.ServiceSearchReq) (list []*entity.TService, err error) {
	// list, err = s.GetFromCache(ctx)
	// if err != nil {
	// 	return
	// }
	// rList := make([]*entity.SysDept, 0, len(list))
	// if req.DeptName != "" || req.Status != "" {
	// 	for _, v := range list {
	// 		if req.DeptName != "" && !gstr.ContainsI(v.DeptName, req.DeptName) {
	// 			continue
	// 		}
	// 		if req.Status != "" && v.Status != gconv.Uint(req.Status) {
	// 			continue
	// 		}
	// 		rList = append(rList, v)
	// 	}
	// 	list = rList
	// }
	return
}

func (s *sService) Edit(ctx context.Context, req *system.ServiceEditReq) (res *system.ServiceEditRes, err error) {
	return
}
