// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserDao is the data access object for table t_user.
type TUserDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns TUserColumns // columns contains all the column names of Table for convenient usage.
}

// TUserColumns defines and stores column names for table t_user.
type TUserColumns struct {
	Id            string // 唯一自增ID
	Openid        string // 用户openid
	From          string // 用户来源
	Nickname      string // 昵称
	Sex           string // 性别
	Province      string // 省份
	City          string // 城市
	Country       string // 国家
	HeadImg       string // 头像
	Balance       string // 账户余额（分）
	Give          string // 赠送金额（分）
	Privilege     string // 用户特权信息
	UnionId       string // 微信unionid
	ShareCode     string // 分享码，唯一
	Phone         string // 手机号
	LastLoginTime string // 最后登录时间
	CreatedTime   string // 创建时间
	ModifiedTime  string // 更新时间
}

// tUserColumns holds the columns for table t_user.
var tUserColumns = TUserColumns{
	Id:            "id",
	Openid:        "openid",
	From:          "from",
	Nickname:      "nickname",
	Sex:           "sex",
	Province:      "province",
	City:          "city",
	Country:       "country",
	HeadImg:       "head_img",
	Balance:       "balance",
	Give:          "give",
	Privilege:     "privilege",
	UnionId:       "union_id",
	ShareCode:     "share_code",
	Phone:         "phone",
	LastLoginTime: "last_login_time",
	CreatedTime:   "created_time",
	ModifiedTime:  "modified_time",
}

// NewTUserDao creates and returns a new DAO object for table data access.
func NewTUserDao() *TUserDao {
	return &TUserDao{
		group:   "default",
		table:   "t_user",
		columns: tUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUserDao) Columns() TUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}