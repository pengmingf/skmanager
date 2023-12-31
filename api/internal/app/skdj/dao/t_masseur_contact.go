// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/skdj/dao/internal"
)

// internalTMasseurContactDao is internal type for wrapping internal DAO implements.
type internalTMasseurContactDao = *internal.TMasseurContactDao

// tMasseurContactDao is the data access object for table t_masseur_contact.
// You can define custom methods on it to extend its functionality as you wish.
type tMasseurContactDao struct {
	internalTMasseurContactDao
}

var (
	// TMasseurContact is globally public accessible object for table t_masseur_contact operations.
	TMasseurContact = tMasseurContactDao{
		internal.NewTMasseurContactDao(),
	}
)

// Fill with you ideas below.
