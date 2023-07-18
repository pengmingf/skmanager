// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)

// internalTOrderComplaintDao is internal type for wrapping internal DAO implements.
type internalTOrderComplaintDao = *internal.TOrderComplaintDao

// tOrderComplaintDao is the data access object for table t_order_complaint.
// You can define custom methods on it to extend its functionality as you wish.
type tOrderComplaintDao struct {
	internalTOrderComplaintDao
}

var (
	// TOrderComplaint is globally public accessible object for table t_order_complaint operations.
	TOrderComplaint = tOrderComplaintDao{
		internal.NewTOrderComplaintDao(),
	}
)

// Fill with you ideas below.
