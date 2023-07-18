// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TAddressDao is the data access object for table t_address.
type TAddressDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns TAddressColumns // columns contains all the column names of Table for convenient usage.
}

// TAddressColumns defines and stores column names for table t_address.
type TAddressColumns struct {
	Id           string // 唯一自增ID
	UserId       string // 用户id
	ContactName  string // 联系人
	Phone        string // 手机号
	Path         string // 地址
	Doorplate    string // 门牌号
	Lat          string // 纬度
	Lnt          string // 经度
	IsDefault    string // 是否默认
	CreatedTime  string // 创建时间
	ModifiedTime string // 更新时间
}

// tAddressColumns holds the columns for table t_address.
var tAddressColumns = TAddressColumns{
	Id:           "id",
	UserId:       "user_id",
	ContactName:  "contact_name",
	Phone:        "phone",
	Path:         "path",
	Doorplate:    "doorplate",
	Lat:          "lat",
	Lnt:          "lnt",
	IsDefault:    "is_default",
	CreatedTime:  "created_time",
	ModifiedTime: "modified_time",
}

// NewTAddressDao creates and returns a new DAO object for table data access.
func NewTAddressDao() *TAddressDao {
	return &TAddressDao{
		group:   "default",
		table:   "t_address",
		columns: tAddressColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TAddressDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TAddressDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TAddressDao) Columns() TAddressColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TAddressDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TAddressDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TAddressDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}