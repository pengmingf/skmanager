// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TServiceDao is the data access object for table t_service.
type TServiceDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns TServiceColumns // columns contains all the column names of Table for convenient usage.
}

// TServiceColumns defines and stores column names for table t_service.
type TServiceColumns struct {
	Id           string // 唯一自增ID
	ServiceId    string // 服务id
	Name         string // 服务项目名称
	Desc         string // 服务项目描述
	Label        string // 标签（显示在名称旁边）
	Price        string // 价格(分)
	SPic         string // 首页缩略图
	BPic         string // 详情页大图片
	ServiceTime  string // 服务时长
	Cnt          string // 服务被预定次数
	Content      string // 详情页内容（图片）
	IsShow       string // 是否首页展示
	Enabled      string // 是否有效0否1是
	CreatedTime  string // 创建时间
	ModifiedTime string // 更新时间
}

// tServiceColumns holds the columns for table t_service.
var tServiceColumns = TServiceColumns{
	Id:           "id",
	ServiceId:    "service_id",
	Name:         "name",
	Desc:         "desc",
	Label:        "label",
	Price:        "price",
	SPic:         "s_pic",
	BPic:         "b_pic",
	ServiceTime:  "service_time",
	Cnt:          "cnt",
	Content:      "content",
	IsShow:       "is_show",
	Enabled:      "enabled",
	CreatedTime:  "created_time",
	ModifiedTime: "modified_time",
}

// NewTServiceDao creates and returns a new DAO object for table data access.
func NewTServiceDao() *TServiceDao {
	return &TServiceDao{
		group:   "default",
		table:   "t_service",
		columns: tServiceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TServiceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TServiceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TServiceDao) Columns() TServiceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TServiceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TServiceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TServiceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
