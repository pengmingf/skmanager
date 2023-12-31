// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	system "github.com/tiger1103/gfast/v3/api/v1/skdj"
	"github.com/tiger1103/gfast/v3/internal/app/skdj/model/entity"
)

type (
	IService interface {
		GetList(ctx context.Context, req *system.ServiceSearchReq) (list []*entity.TService, err error)
		Edit(ctx context.Context, req *system.ServiceEditReq) (res *system.ServiceEditRes, err error)
	}
)

var (
	localService IService
)

func Service() IService {
	if localService == nil {
		panic("implement not found for interface IService, forgot register?")
	}
	return localService
}

func RegisterService(i IService) {
	localService = i
}
