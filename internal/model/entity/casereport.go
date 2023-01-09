// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Casereport is the golang structure for table casereport.
type Casereport struct {
	Id       string      `json:"id"       ` //
	Message  string      `json:"message"  ` //
	CaseId   int         `json:"caseId"   ` //
	CaseDesc string      `json:"caseDesc" ` //
	RunTime  *gtime.Time `json:"runTime"  ` //
	Status   int         `json:"status"   ` //
	RunUser  string      `json:"runUser"  ` //
	RunIp    string      `json:"runIp"    ` //
}
