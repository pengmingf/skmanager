// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TMasseurPicExamineDao is the data access object for table t_masseur_pic_examine.
type TMasseurPicExamineDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns TMasseurPicExamineColumns // columns contains all the column names of Table for convenient usage.
}

// TMasseurPicExamineColumns defines and stores column names for table t_masseur_pic_examine.
type TMasseurPicExamineColumns struct {
	Id           string // 唯一自增ID
	MasseurId    string //
	Pic          string // 图片
	PicType      string // 图片类型,profession职业照，life生活照
	Approval     string // 0待审批1审批通过2审批不通过
	ApprovalUser string // 审批人
	Enabled      string //
	CreatedTime  string // 创建时间
	ModifiedTime string // 更新时间
}

// tMasseurPicExamineColumns holds the columns for table t_masseur_pic_examine.
var tMasseurPicExamineColumns = TMasseurPicExamineColumns{
	Id:           "id",
	MasseurId:    "masseur_id",
	Pic:          "pic",
	PicType:      "pic_type",
	Approval:     "approval",
	ApprovalUser: "approval_user",
	Enabled:      "enabled",
	CreatedTime:  "created_time",
	ModifiedTime: "modified_time",
}

// NewTMasseurPicExamineDao creates and returns a new DAO object for table data access.
func NewTMasseurPicExamineDao() *TMasseurPicExamineDao {
	return &TMasseurPicExamineDao{
		group:   "default",
		table:   "t_masseur_pic_examine",
		columns: tMasseurPicExamineColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TMasseurPicExamineDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TMasseurPicExamineDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TMasseurPicExamineDao) Columns() TMasseurPicExamineColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TMasseurPicExamineDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TMasseurPicExamineDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TMasseurPicExamineDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
