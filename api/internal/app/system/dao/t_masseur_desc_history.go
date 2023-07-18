// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)

// internalTMasseurDescHistoryDao is internal type for wrapping internal DAO implements.
type internalTMasseurDescHistoryDao = *internal.TMasseurDescHistoryDao

// tMasseurDescHistoryDao is the data access object for table t_masseur_desc_history.
// You can define custom methods on it to extend its functionality as you wish.
type tMasseurDescHistoryDao struct {
	internalTMasseurDescHistoryDao
}

var (
	// TMasseurDescHistory is globally public accessible object for table t_masseur_desc_history operations.
	TMasseurDescHistory = tMasseurDescHistoryDao{
		internal.NewTMasseurDescHistoryDao(),
	}
)

// Fill with you ideas below.
