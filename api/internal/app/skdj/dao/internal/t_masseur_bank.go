// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TMasseurBankDao is the data access object for table t_masseur_bank.
type TMasseurBankDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns TMasseurBankColumns // columns contains all the column names of Table for convenient usage.
}

// TMasseurBankColumns defines and stores column names for table t_masseur_bank.
type TMasseurBankColumns struct {
	Id           string // 唯一自增ID
	MasseurId    string //
	CardId       string // 银行卡号
	CardName     string // 持卡人姓名
	BankName     string // 所属银行
	SubBankName  string // 分行
	Enabled      string //
	CreatedTime  string // 创建时间
	ModifiedTime string // 更新时间
}

// tMasseurBankColumns holds the columns for table t_masseur_bank.
var tMasseurBankColumns = TMasseurBankColumns{
	Id:           "id",
	MasseurId:    "masseur_id",
	CardId:       "card_id",
	CardName:     "card_name",
	BankName:     "bank_name",
	SubBankName:  "sub_bank_name",
	Enabled:      "enabled",
	CreatedTime:  "created_time",
	ModifiedTime: "modified_time",
}

// NewTMasseurBankDao creates and returns a new DAO object for table data access.
func NewTMasseurBankDao() *TMasseurBankDao {
	return &TMasseurBankDao{
		group:   "default",
		table:   "t_masseur_bank",
		columns: tMasseurBankColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TMasseurBankDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TMasseurBankDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TMasseurBankDao) Columns() TMasseurBankColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TMasseurBankDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TMasseurBankDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TMasseurBankDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
