// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/skdj/dao/internal"
)

// internalTUserRefundDao is internal type for wrapping internal DAO implements.
type internalTUserRefundDao = *internal.TUserRefundDao

// tUserRefundDao is the data access object for table t_user_refund.
// You can define custom methods on it to extend its functionality as you wish.
type tUserRefundDao struct {
	internalTUserRefundDao
}

var (
	// TUserRefund is globally public accessible object for table t_user_refund operations.
	TUserRefund = tUserRefundDao{
		internal.NewTUserRefundDao(),
	}
)

// Fill with you ideas below.
