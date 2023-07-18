// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TInviteAttentionDao is the data access object for table t_invite_attention.
type TInviteAttentionDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns TInviteAttentionColumns // columns contains all the column names of Table for convenient usage.
}

// TInviteAttentionColumns defines and stores column names for table t_invite_attention.
type TInviteAttentionColumns struct {
	Id           string // 唯一自增ID
	UserId       string // 邀请人id
	UserOpenid   string // 邀请人openid
	ShareCode    string // 邀请码（解码后的值，暂时无用）
	Openid       string // 被邀请人openid
	Event        string // 事件（关注subject、扫码scan）
	Enabled      string //
	CreatedTime  string // 创建时间
	ModifiedTime string // 更新时间
}

// tInviteAttentionColumns holds the columns for table t_invite_attention.
var tInviteAttentionColumns = TInviteAttentionColumns{
	Id:           "id",
	UserId:       "user_id",
	UserOpenid:   "user_openid",
	ShareCode:    "share_code",
	Openid:       "openid",
	Event:        "event",
	Enabled:      "enabled",
	CreatedTime:  "created_time",
	ModifiedTime: "modified_time",
}

// NewTInviteAttentionDao creates and returns a new DAO object for table data access.
func NewTInviteAttentionDao() *TInviteAttentionDao {
	return &TInviteAttentionDao{
		group:   "default",
		table:   "t_invite_attention",
		columns: tInviteAttentionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TInviteAttentionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TInviteAttentionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TInviteAttentionDao) Columns() TInviteAttentionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TInviteAttentionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TInviteAttentionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TInviteAttentionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}