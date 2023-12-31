// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserDetailsDao is the data access object for table t_user_details.
type TUserDetailsDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns TUserDetailsColumns // columns contains all the column names of Table for convenient usage.
}

// TUserDetailsColumns defines and stores column names for table t_user_details.
type TUserDetailsColumns struct {
	Id           string // 唯一自增ID
	UserId       string // 用户openid
	OrderId      string // 订单号
	Status       string // 状态，已完成finish，失败fail
	Money        string // 金额(分)
	Enabled      string //
	CreatedTime  string // 创建时间
	ModifiedTime string // 更新时间
}

// tUserDetailsColumns holds the columns for table t_user_details.
var tUserDetailsColumns = TUserDetailsColumns{
	Id:           "id",
	UserId:       "user_id",
	OrderId:      "order_id",
	Status:       "status",
	Money:        "money",
	Enabled:      "enabled",
	CreatedTime:  "created_time",
	ModifiedTime: "modified_time",
}

// NewTUserDetailsDao creates and returns a new DAO object for table data access.
func NewTUserDetailsDao() *TUserDetailsDao {
	return &TUserDetailsDao{
		group:   "default",
		table:   "t_user_details",
		columns: tUserDetailsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUserDetailsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUserDetailsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUserDetailsDao) Columns() TUserDetailsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUserDetailsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUserDetailsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUserDetailsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
