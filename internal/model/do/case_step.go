// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CaseStep is the golang structure of table case_step for DAO operations like Where/Data.
type CaseStep struct {
	g.Meta       `orm:"table:case_step, do:true"`
	Id           interface{} //
	StepId       interface{} //
	Input        interface{} //
	AssertExpect interface{} //
	Description  interface{} //
	CaseId       interface{} //
	Score        interface{} //
}