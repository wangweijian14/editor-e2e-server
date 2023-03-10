package controller

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
	v1 "wiki/api/v1"
	"wiki/internal/consts"
	"wiki/internal/model"
	"wiki/internal/model/entity"
	"wiki/internal/service"

	"github.com/go-rod/rod/lib/input"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *cCases) Execute(ctx context.Context, req *v1.CasesExecuteReq) (res *v1.CasesExecuteRes, err error) {
	res = &v1.CasesExecuteRes{IsSuccess: true, StartTime: gtime.Datetime()}
	cases, err := service.Cases().GetById(ctx, gconv.Uint64(req.Id))
	if err != nil {
		return nil, err
	}
	res.CaseDescription = cases.Cases.Description
	page2Open, err := service.Page().GetById(ctx, gconv.Uint64(cases.Cases.OpenId))
	if err != nil {
		return nil, err
	}
	// 用例执行完，不管失败与成功均关闭页面
	// defer service.Action().ClosePage()
	// 计算cese执行时间
	defer ExcutreCasetimeCost(gtime.Now().Time, res, cases.Cases, ctx)
	// 进入MustPage
	g.Log().Info(ctx, page2Open, "aaa")
	if page2Open != nil {
		service.Action().MustPage(page2Open.Page.Url)
	}
	// defer service.Action().ClosePage()
	for _, caseStep := range cases.CaseStepOutput {
		resI, err := parseStep(ctx, gconv.Uint64(caseStep.CaseStep.StepId), caseStep.CaseStep.Input)
		v1StepRes := &v1.StepResults{
			IsSuccess:           true,
			StepId:              caseStep.Step.Step.Id,
			CaseStepDescription: caseStep.Step.Step.Description,
			ElementName:         caseStep.Step.Element.Element.Name,
			ElementPath:         caseStep.Step.Element.Element.Path,
			UnixMilli:           gtime.TimestampMilli(),
		}
		if err != nil {
			// service.Action().MustScreenshot("resource/public/resource/image/error-page.jpg")
			res.ErrPageSnapshot = "https://localhots:8000/image/error-page.jpg"
			res.IsSuccess = false
			v1StepRes.Error = err
			res.StepResults = append(res.StepResults, v1StepRes)
			res.Error = err.Error()
			updateCases(false, cases.Cases, ctx)
			return res, err
		}

		//处理断言
		if caseStep.CaseStep.AssertExpect != "" {
			canzhao := strings.Split(caseStep.CaseStep.AssertExpect, "<->")
			if len(canzhao) > 0 {
				assertRes := &v1.AssertRes{}
				for _, kk := range canzhao {
					canzhaoInner := strings.Split(kk, ":")
					if len(canzhaoInner) == 3 {
						description := ""
						if resI.Target != nil {
							description = resI.Target.Object.Description
						}

						vfunc := service.ValidatorI().GetValidator(canzhaoInner[1])
						var r *model.ValidatorResult
						assertRes.Method = strings.ToLower(canzhaoInner[0])

						switch strings.ToLower(canzhaoInner[0]) {
						case "text":
							r = vfunc(resI.Returned, canzhaoInner[2])
						case "target":
							r = vfunc(description, canzhaoInner[2])
						}

						assertRes.Message = r.Message
						assertRes.Error = r.Error
						assertRes.IsSuccess = r.IsPass

						if r.IsPass == false {
							v1StepRes.IsSuccess = false
							res.IsSuccess = false
						}
					}
				}
				v1StepRes.AssertRes = append(v1StepRes.AssertRes, assertRes)
			}
		}

		g.Log().Infof(ctx, "resI: %v", resI)
		res.StepResults = append(res.StepResults, v1StepRes)
	}
	updateCases(res.IsSuccess, cases.Cases, ctx)
	return res, nil
}

func updateCases(ok bool, cases *entity.Cases, ctx context.Context) {
	traceId := gctx.CtxId(ctx)
	// 统计case 执行次数
	cases.ExecCount = cases.ExecCount + 1
	if ok {
		cases.PassCount = cases.PassCount + 1
	} else {
		cases.FailedCount = cases.FailedCount + 1
	}

	cases.HisId = traceId
	service.Cases().Update(ctx, cases)
}

// insert  report
func insertReport(res *v1.CasesExecuteRes, cases *entity.Cases, ctx context.Context) {
	traceId := gctx.CtxId(ctx)
	report := &entity.Casereport{
		Id:       traceId,
		Message:  gjson.New(res).MustToJsonString(),
		CaseId:   cases.Id,
		CaseDesc: cases.Description,
		RunTime:  gtime.Now(),
		RunUser:  "wangweijian",
	}
	if res.IsSuccess {
		report.Status = consts.StatusSuccess
	} else {
		report.Status = consts.StatusFailed
	}
	service.Report().Create(ctx, report)
	res.HisId = traceId
}

func (c *cCases) ExecuteStep(ctx context.Context, req *v1.StepExecuteReq) (res *v1.StepExecuteRes, err error) {
	resI, err := parseStep(ctx, req.Id, req.Input)
	if err != nil {
		return nil, err
	}

	//处理断言
	assertRes := make(map[string]interface{})
	if req.AssertExpect != "" {
		canzhao := strings.Split(req.AssertExpect, "<->")
		if len(canzhao) > 0 {
			for _, kk := range canzhao {
				canzhaoInner := strings.Split(kk, ":")
				if len(canzhaoInner) == 3 {
					description := ""
					if resI.Target != nil {
						description = resI.Target.Object.Description
					}

					vfunc := service.ValidatorI().GetValidator(canzhaoInner[1])
					switch strings.ToLower(canzhaoInner[0]) {
					case "text":
						r := vfunc(resI.Returned, canzhaoInner[2])
						assertRes[canzhaoInner[0]] = r
					case "target":
						r := vfunc(description, canzhaoInner[2])
						assertRes[canzhaoInner[0]] = r
					}
				}
			}
		}
	}

	return &v1.StepExecuteRes{R: resI, AssertRes: assertRes}, nil
}

func (c *cCases) OpenPage(ctx context.Context, req *v1.StepOpenPageReq) (res *v1.StepOpenPageRes, err error) {
	return &v1.StepOpenPageRes{
		R: service.Action().MustPage(req.Url),
	}, nil
}

func (c *cCases) OpenPageById(ctx context.Context, req *v1.StepOpenPageByIdReq) (res *v1.StepOpenPageRes, err error) {
	page, err := service.Page().GetById(ctx, gconv.Uint64(req.PageId))
	if err != nil {
		return nil, err
	}
	return &v1.StepOpenPageRes{
		R: service.Action().MustPage(page.Page.Url),
	}, nil
}

func parseStep(ctx context.Context, stepId uint64, inputText string) (res *model.ActionReturned, err error) {
	step, err := service.Step().GetById(ctx, stepId)
	if err != nil {
		return nil, err
	}
	res = &model.ActionReturned{}
	res.StepOutput = step
	fmt.Printf("执行step %v : action-%v ,元素 : %v...\n", step.Step.Description,
		step.Step.ActionId, step.Element.Element.Description)
	switch step.Step.ActionId {
	case consts.FindTarget:
		res.Target = service.Action().MustFindTarget(ctx, step.Element)

	case consts.FindAndClick:
		res.Target = service.Action().MustFindTargetAndMustClick(ctx, step.Element)

	case consts.FindAndInputSth:
		res.Target = service.Action().MustFindTargetAndMustInputText(ctx, step.Element, inputText)

	case consts.FindTargetText:
		res.Returned = service.Action().MustFindTargetText(ctx, step.Element)

	case consts.FindTargetAttribute:
		res.Returned = service.Action().MustFindTargetAttribute(ctx, step.Element)

	case consts.FindTargetAndIteratorChildAttribute:
		res.Returned = strings.Join(service.Action().MustFindTargetAndIteratorChildAttribute(ctx, step.Element), "<->")

	case consts.FindTargetAndIteratorChildText:
		res.Returned = strings.Join(service.Action().MustFindTargetAndIteratorChildText(ctx, step.Element), "<->")

	case consts.FindAndHover:
		res.Target = service.Action().MustFindTargetAndMustHover(ctx, step.Element)

	case consts.KeybordCommandA:
		service.Action().MustUseKeyboardPressControlLeftAndType(ctx, input.KeyA)

	case consts.KeybordCommandZ:
		service.Action().MustUseKeyboardPressControlLeftAndType(ctx, 'z')

	case consts.KeybordCommandX:
		service.Action().MustUseKeyboardPressControlLeftAndType(ctx, 'x')

	case consts.KeybordCommandC:
		service.Action().MustUseKeyboardPressControlLeftAndType(ctx, 'c')

	case consts.KeybordCommandV:
		service.Action().MustUseKeyboardPressControlLeftAndType(ctx, input.KeyV)

	case consts.KeybordEnter:
		service.Action().MustUseKeyboardType(ctx, keyStepLogic(gconv.Int(inputText), input.Enter)...)

	case consts.KeybordEnterEsc:
		service.Action().MustUseKeyboardType(ctx, keyStepLogic(gconv.Int(inputText), input.Escape)...)

	case consts.KeybordArrowLeft:
		service.Action().MustUseKeyboardType(ctx, keyStepLogic(gconv.Int(inputText), input.ArrowLeft)...)

	case consts.KeybordArrowUp:
		service.Action().MustUseKeyboardType(ctx, keyStepLogic(gconv.Int(inputText), input.ArrowUp)...)

	case consts.KeybordArrowDown:
		service.Action().MustUseKeyboardType(ctx, keyStepLogic(gconv.Int(inputText), input.ArrowDown)...)

	case consts.KeybordArrowRight:
		service.Action().MustUseKeyboardType(ctx, keyStepLogic(gconv.Int(inputText), input.ArrowRight)...)

	default:
		err = errors.New("action 没有被实现，联系测试实现这个action")
	}

	return res, err
}

func keyStepLogic(step int, key input.Key) []input.Key {
	if step == 0 {
		step = 1
	}
	keys := make([]input.Key, 0)
	for i := 0; i < step; i++ {
		keys = append(keys, key)
	}
	return keys
}

/**
笔记：测试结果： Go 中 defer 和 return 执行的先后顺序
多个defer的执行顺序为“后进先出”；
defer、return、返回值三者的执行逻辑应该是：return最先执行，return负责将结果写入返回值中；接着defer开始执行一些收尾工作；最后函数携带当前返回值退出。

如果函数的返回值是无名的（不带命名返回值），则go语言会在执行return的时候会执行一个类似创建一个临时变量作为保存return值的动作，而有名返回值的函数，由于返回值在函数定义的时候已经将该变量进行定义，在执行return的时候会先执行返回值保存操作，而后续的defer函数会改变这个返回值(虽然defer是在return之后执行的，但是由于使用的函数定义的变量，所以执行defer操作后对该变量的修改会影响到return的值
*/
//耗时统计函数 单位毫秒
func ExcutreCasetimeCost(start time.Time, res *v1.CasesExecuteRes, cases *entity.Cases, ctx context.Context) {
	tc := time.Since(start)
	fmt.Printf("time cost = %v\n", tc)
	res.Duration = tc.Milliseconds()
	insertReport(res, cases, ctx)
}
