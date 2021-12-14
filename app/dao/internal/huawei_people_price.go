// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// HuaweiPeoplePriceDao is the manager for logic model data accessing and custom defined data operations functions management.
type HuaweiPeoplePriceDao struct {
	Table   string                   // Table is the underlying table name of the DAO.
	Group   string                   // Group is the database configuration group name of current DAO.
	Columns HuaweiPeoplePriceColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// HuaweiPeoplePriceColumns defines and stores column names for table huawei_people_price.
type HuaweiPeoplePriceColumns struct {
	Id            string //
	Model         string // 型号
	Category      string // 分类  大类/小类
	ServiceMethod string // 服务方式
	ServiceCate   string // 服务分类
	Price         string // 人工费
}

//  huaweiPeoplePriceColumns holds the columns for table huawei_people_price.
var huaweiPeoplePriceColumns = HuaweiPeoplePriceColumns{
	Id:            "id",
	Model:         "model",
	Category:      "category",
	ServiceMethod: "service_method",
	ServiceCate:   "service_cate",
	Price:         "price",
}

// NewHuaweiPeoplePriceDao creates and returns a new DAO object for table data access.
func NewHuaweiPeoplePriceDao() *HuaweiPeoplePriceDao {
	return &HuaweiPeoplePriceDao{
		Group:   "default",
		Table:   "huawei_people_price",
		Columns: huaweiPeoplePriceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HuaweiPeoplePriceDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HuaweiPeoplePriceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HuaweiPeoplePriceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
