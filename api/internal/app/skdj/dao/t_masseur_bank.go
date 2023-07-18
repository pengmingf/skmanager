// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/skdj/dao/internal"
)

// internalTMasseurBankDao is internal type for wrapping internal DAO implements.
type internalTMasseurBankDao = *internal.TMasseurBankDao

// tMasseurBankDao is the data access object for table t_masseur_bank.
// You can define custom methods on it to extend its functionality as you wish.
type tMasseurBankDao struct {
	internalTMasseurBankDao
}

var (
	// TMasseurBank is globally public accessible object for table t_masseur_bank operations.
	TMasseurBank = tMasseurBankDao{
		internal.NewTMasseurBankDao(),
	}
)

// Fill with you ideas below.