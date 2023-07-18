// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/system/dao/internal"
)

// internalToolsGenTableDao is internal type for wrapping internal DAO implements.
type internalToolsGenTableDao = *internal.ToolsGenTableDao

// toolsGenTableDao is the data access object for table tools_gen_table.
// You can define custom methods on it to extend its functionality as you wish.
type toolsGenTableDao struct {
	internalToolsGenTableDao
}

var (
	// ToolsGenTable is globally public accessible object for table tools_gen_table operations.
	ToolsGenTable = toolsGenTableDao{
		internal.NewToolsGenTableDao(),
	}
)

// Fill with you ideas below.