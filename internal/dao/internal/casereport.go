// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CasereportDao is the data access object for table casereport.
type CasereportDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns CasereportColumns // columns contains all the column names of Table for convenient usage.
}

// CasereportColumns defines and stores column names for table casereport.
type CasereportColumns struct {
	Id       string //
	Message  string //
	CaseId   string //
	CaseDesc string //
	RunTime  string //
	Status   string //
	RunUser  string //
	RunIp    string //
}

// casereportColumns holds the columns for table casereport.
var casereportColumns = CasereportColumns{
	Id:       "id",
	Message:  "message",
	CaseId:   "case_id",
	CaseDesc: "case_desc",
	RunTime:  "run_time",
	Status:   "status",
	RunUser:  "run_user",
	RunIp:    "run_ip",
}

// NewCasereportDao creates and returns a new DAO object for table data access.
func NewCasereportDao() *CasereportDao {
	return &CasereportDao{
		group:   "default",
		table:   "casereport",
		columns: casereportColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CasereportDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CasereportDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CasereportDao) Columns() CasereportColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CasereportDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CasereportDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CasereportDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
