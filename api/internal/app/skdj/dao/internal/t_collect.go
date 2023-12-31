// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TCollectDao is the data access object for table t_collect.
type TCollectDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns TCollectColumns // columns contains all the column names of Table for convenient usage.
}

// TCollectColumns defines and stores column names for table t_collect.
type TCollectColumns struct {
	Id          string // 唯一自增ID
	UserId      string // 用户id
	MasseurId   string // 技师id
	CreatedTime string // 创建时间
}

// tCollectColumns holds the columns for table t_collect.
var tCollectColumns = TCollectColumns{
	Id:          "id",
	UserId:      "user_id",
	MasseurId:   "masseur_id",
	CreatedTime: "created_time",
}

// NewTCollectDao creates and returns a new DAO object for table data access.
func NewTCollectDao() *TCollectDao {
	return &TCollectDao{
		group:   "default",
		table:   "t_collect",
		columns: tCollectColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TCollectDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TCollectDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TCollectDao) Columns() TCollectColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TCollectDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TCollectDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TCollectDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
