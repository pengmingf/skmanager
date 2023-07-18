// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/skdj/dao/internal"
)

// internalSysOperLogDao is internal type for wrapping internal DAO implements.
type internalSysOperLogDao = *internal.SysOperLogDao

// sysOperLogDao is the data access object for table sys_oper_log.
// You can define custom methods on it to extend its functionality as you wish.
type sysOperLogDao struct {
	internalSysOperLogDao
}

var (
	// SysOperLog is globally public accessible object for table sys_oper_log operations.
	SysOperLog = sysOperLogDao{
		internal.NewSysOperLogDao(),
	}
)

// Fill with you ideas below.
