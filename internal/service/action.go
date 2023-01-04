// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"time"
	"wiki/internal/model"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

type (
	IAction interface {
		MustPage(uri string) *rod.Page
		MustFindTarget(ctx context.Context, target *model.ElementOutput) *rod.Element
		MustFindTargets(ctx context.Context, target *model.ElementOutput) []*rod.Element
		MustFindTargetAndMustClick(ctx context.Context, target *model.ElementOutput) *rod.Element
		MustFindTargetAndMustInputText(ctx context.Context, target *model.ElementOutput, input string) *rod.Element
		MustFindTargetAndMustInputTime(ctx context.Context, target *model.ElementOutput, time time.Time) *rod.Element
		MustFindTargetAttribute(ctx context.Context, target *model.ElementOutput) string
		MustFindTargetText(ctx context.Context, target *model.ElementOutput) string
		MustFindTargetAndIteratorChildAttribute(ctx context.Context, target *model.ElementOutput) []string
		MustFindTargetAndIteratorChildText(ctx context.Context, target *model.ElementOutput) []string
		MustFindTargetAndIteratorChildHtml(ctx context.Context, target *model.ElementOutput) []string
		MustUseKeyboardPressControlLeftAndType(ctx context.Context, keys ...input.Key) *rod.Element
		MustUseKeyboardType(ctx context.Context, keys ...input.Key) *rod.Element
		MustFindTargetAndMustHover(ctx context.Context, target *model.ElementOutput) *rod.Element
		MustScreenshot(name string)
		ClosePage()
		GetServeMonitorPath() string
	}
)

var (
	localAction IAction
)

func Action() IAction {
	if localAction == nil {
		panic("implement not found for interface IAction, forgot register?")
	}
	return localAction
}

func RegisterAction(i IAction) {
	localAction = i
}
