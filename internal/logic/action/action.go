package action

import (
	"context"
	"strings"
	"time"
	"wiki/internal/model"
	"wiki/internal/service"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAction struct {
	Browser      *rod.Browser
	Page         *rod.Page
	Managed      string
	ServeMonitor string
}

func init() {
	service.RegisterAction(New())
}
func New() *sAction {

	// 分布式浏览器支持
	// This example is to launch a browser remotely, not connect to a running browser remotely,
	// to connect to a running browser check the "../connect-browser" example.
	// Rod provides a docker image for beginners, run the below to start a launcher.Manager:
	//
	//     docker run -p 7317:7317 ghcr.io/go-rod/rod
	//
	// For available CLI flags run: docker run ghcr.io/go-rod/rod rod-manager -h
	// For more information, check the doc of launcher.Manager
	var ctx = gctx.New()
	cf, err := gcfg.Instance().Get(ctx, "browser")
	if err != nil {
		panic("读取browser 配置失败...")
	}
	browserCf := &model.BrowserConfig{}

	err = gconv.Struct(cf, browserCf)
	if err != nil {
		panic("读取browser 配置失败...")
	}

	if browserCf.LocalBrowser {

		// 本地浏览器connect，计划使用Linux，若本地，使用此方式
		// TODO： 实现可配置
		browser := rod.New().MustConnect()
		return &sAction{
			Browser: browser,
		}
	}

	l := launcher.MustNewManaged(browserCf.LauncherManager)

	// You can also set any flag remotely before you launch the remote browser.
	// Available flags: https://peter.sh/experiments/chromium-command-line-switches
	l.Set("disable-gpu").Delete("disable-gpu")

	// Launch with headful mode
	if !browserCf.LaunchHeadfulMode {
		l.Headless(false).XVFB("--server-num=5", "--server-args=-screen 0 1600x900x16")
	}

	browser := rod.New().Client(l.MustClient()).MustConnect()

	// You may want to start a server to watch the screenshots of the remote browser.
	serveMonitor := browser.ServeMonitor(browserCf.ServeMonitor) //TODO 可配置
	g.Log().Info(ctx, "serveMonitor : ", serveMonitor)
	launcher.Open(serveMonitor)

	return &sAction{
		Browser:      browser,
		Managed:      "",
		ServeMonitor: serveMonitor,
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

func (s *sAction) GetServeMonitorPath() string {
	return s.ServeMonitor
}

func paresRetElement(ctx context.Context, page *rod.Page, target *model.ElementOutput) *rod.Element {
	g.Log("action").Infof(ctx, "page: %v desc: %v path: %v", target.Page.Name, target.Element.Description, target.Element.Path)
	if target.FatherElement != nil {
		fatherSplit := strings.Split(target.FatherElement.Path, ",")
		elSplite := strings.Split(target.Element.Path, ",")

		var f *rod.Element
		if len(fatherSplit) > 1 {
			f = page.MustElementR(fatherSplit[0], fatherSplit[1])
		} else {
			f = page.MustElement(target.FatherElement.Path)
		}

		if len(elSplite) > 1 {
			return f.MustElementR(elSplite[0], elSplite[1])
		}

		return f.MustElement(target.Element.Path)
	}

	split := strings.Split(target.Element.Path, ",")
	if len(split) > 1 {
		return page.MustElementR(split[0], split[1])
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
