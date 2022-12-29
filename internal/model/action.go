package model

import "github.com/go-rod/rod"

type ActionReturned struct {
	Target   *rod.Element `json:"target" description:"动作目标"`
	Returned string       `json:"returned" description:"动作结果"`
}
