package action

import (
	"context"
	"time"
	"wiki/internal/model"
	"wiki/internal/service"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/gogf/gf/v2/frame/g"
)

type sAction struct {
	Browser *rod.Browser
	Page    *rod.Page
}

func init() {
	service.RegisterAction(New())
}
func New() *sAction {
	return &sAction{
		Browser: rod.New().MustConnect(),
	}
}

// MustPage 进入Page ，输入url
func (s *sAction) MustPage(uri string) *rod.Page {
	s.Page = s.Browser.MustPage(uri).MustWaitLoad()
	return s.Page
}

// MustFindTarget 寻找目标,单个
func (s *sAction) MustFindTarget(ctx context.Context, target *model.ElementOutput) *rod.Element {
	return paresRetElement(ctx, s.Page, target)
}

// MustFindTargets 寻找目标，多个
func (s *sAction) MustFindTargets(ctx context.Context, target *model.ElementOutput) []*rod.Element {
	return paresRetElements(ctx, s.Page, target)
}

// MustFindTargetAndMustClick 找到目标点击
func (s *sAction) MustFindTargetAndMustClick(ctx context.Context, target *model.ElementOutput) *rod.Element {
	return paresRetElement(ctx, s.Page, target).MustClick()
}

func (s *sAction) MustFindTargetAndMustInputText(ctx context.Context, target *model.ElementOutput, input string) *rod.Element {
	//return paresRetElement(ctx,s.Page,target).MustSelectAllText().MustInput(input)
	return paresRetElement(ctx, s.Page, target).MustClick().MustInput(input)
}

func (s *sAction) MustFindTargetAndMustInputTime(ctx context.Context, target *model.ElementOutput, time time.Time) *rod.Element {

	return paresRetElement(ctx, s.Page, target).MustInputTime(time)
}

func (s *sAction) MustFindTargetAttribute(ctx context.Context, target *model.ElementOutput) string {
	attribute := paresRetElement(ctx, s.Page, target).MustAttribute(target.Element.Attribute)
	if attribute != nil {
		return *attribute
	}
	return ""
}

func (s *sAction) MustFindTargetText(ctx context.Context, target *model.ElementOutput) string {
	return paresRetElement(ctx, s.Page, target).MustText()
}

func (s *sAction) MustFindTargetAndIteratorChildAttribute(ctx context.Context, target *model.ElementOutput) []string {
	elements := paresRetElements(ctx, s.Page, target)
	res := make([]string, len(elements))
	for _, el := range elements {
		att := el.MustAttribute(target.Element.Attribute)
		if att != nil {
			res = append(res, *att)
		}
	}

	return res
}

func (s *sAction) MustFindTargetAndIteratorChildText(ctx context.Context, target *model.ElementOutput) []string {
	elements := paresRetElements(ctx, s.Page, target)
	res := make([]string, len(elements))
	for _, el := range elements {
		res = append(res, el.MustText())
	}
	return res
}

func (s *sAction) MustFindTargetAndIteratorChildHtml(ctx context.Context, target *model.ElementOutput) []string {
	elements := paresRetElements(ctx, s.Page, target)
	res := make([]string, len(elements))
	for _, el := range elements {
		res = append(res, el.MustHTML())
	}
	return res
}

// 模拟键盘操作 按着左侧Shift的同时，输入Keyboard
// 参考规则
// https://github.com/microsoft/playwright/blob/main/packages/playwright-core/src/server/usKeyboardLayout.ts
func (s *sAction) MustUseKeyboardPressControlLeftAndType(ctx context.Context, keys ...input.Key) *rod.Element {
	s.Page.KeyActions().Press(input.MetaLeft).Type(keys...).MustDo()
	return nil
}

func (s *sAction) MustUseKeyboardType(ctx context.Context, keys ...input.Key) *rod.Element {
	s.Page.KeyActions().Type(keys...).MustDo()
	return nil
}

func (s *sAction) MustFindTargetAndMustHover(ctx context.Context, target *model.ElementOutput) *rod.Element {
	return paresRetElement(ctx, s.Page, target).MustHover()
}

func (s *sAction) MustScreenshot(name string) {
	s.Page.MustScreenshot(name)
}

func (s *sAction) ClosePage() {
	s.Page.MustClose()
}

func paresRetElement(ctx context.Context, page *rod.Page, target *model.ElementOutput) *rod.Element {
	g.Log("action").Infof(ctx, "page: %v desc: %v path: %v", target.Page.Name, target.Element.Description, target.Element.Path)
	if target.FatherElement != nil {
		return page.MustElement(target.FatherElement.Path).MustElement(target.Element.Path)
	}
	return page.MustElement(target.Element.Path)
}

func paresRetElements(ctx context.Context, page *rod.Page, target *model.ElementOutput) rod.Elements {
	g.Log("action").Info(ctx)
	if target.FatherElement != nil {
		return page.MustElement(target.FatherElement.Path).MustElements(target.Element.Path)
	}
	return page.MustElements(target.Element.Path)
}
