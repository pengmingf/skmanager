// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)

// internalTUserDetailsDao is internal type for wrapping internal DAO implements.
type internalTUserDetailsDao = *internal.TUserDetailsDao

// tUserDetailsDao is the data access object for table t_user_details.
// You can define custom methods on it to extend its functionality as you wish.
type tUserDetailsDao struct {
	internalTUserDetailsDao
}

var (
	// TUserDetails is globally public accessible object for table t_user_details operations.
	TUserDetails = tUserDetailsDao{
		internal.NewTUserDetailsDao(),
	}
)

// Fill with you ideas below.
