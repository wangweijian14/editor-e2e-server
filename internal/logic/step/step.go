package step

import (
	"context"
	"wiki/internal/consts"
	"wiki/internal/dao"
	"wiki/internal/model"
	"wiki/internal/model/entity"
	"wiki/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
)

type sStep struct{}

func init() {
	service.RegisterStep(New())
}
func New() *sStep {
	return &sStep{}
}

// GetList 查询step列表
func (s *sStep) GetList(ctx context.Context, in model.StepListInput) (out *model.StepListOutput, err error) {
	var (
		m = dao.Step.Ctx(ctx)
	)
	out = &model.StepListOutput{
		Page:  in.Page,
		Limit: in.Limit,
	}

	// 名称模糊
	if in.Description != "" {
		m = m.WhereLike(dao.Step.Columns().Description, in.Description)
	}

	// pageId 过滤
	if in.TargetPage != 0 {
		m = m.Where(dao.Step.Columns().TargetPage, in.TargetPage)
	}

	// elementId 过滤
	if in.TargetElementId != 0 {
		m = m.Where(dao.Step.Columns().TargetElementId, in.TargetElementId)
	}

	// Action 过滤
	if in.ActionId != 0 {
		m = m.Where(dao.Step.Columns().ActionId, in.ActionId)
	}

	// 多 id 查询
	if in.Ids != "" {
		m = m.WhereIn(dao.Step.Columns().Id, in.Ids)
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Limit)
	// 排序方式
	switch in.Sort {
	case consts.PageSortActive:
		listModel = listModel.OrderAsc(dao.Step.Columns().Id)

	default:
		listModel = listModel.OrderDesc(dao.Step.Columns().Id)
	}

	// 执行查询
	var list []*entity.Step
	if err = listModel.Scan(&list); err != nil {
		return out, err
	}
	out.Total, err = m.Count()
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.List = list
	if err != nil {
		return out, err
	}
	return
}

// GetById 根据ID查询 Step
func (s *sStep) GetById(ctx context.Context, in uint64) (out *model.StepOutput, err error) {
	var (
		m = dao.Step.Ctx(ctx)
	)
	out = &model.StepOutput{}
	step := &entity.Step{}

	err = m.Where(dao.Step.Columns().Id, in).Scan(step)
	if err != nil {
		return nil, err
	}
	// 没有数据
	if nil == step || step.Id == 0 {
		return nil, err
	}
	out.Step = step
	// 查询 所属 page 信息
	if step.TargetPage != 0 {
		page, err := service.Page().GetById(ctx, gconv.Uint64(step.TargetPage))
		if err != nil {
			return out, err
		}
		out.Page = page.Page
	}
	// 查询 所属 用到的 Page 信息
	if step.TargetPage != 0 {
		page, err := service.Page().GetById(ctx, gconv.Uint64(step.TargetPage))
		if err != nil {
			return out, err
		}
		out.Page = page.Page

	}
	// 查询 用到的Element
	if step.TargetElementId != 0 {
		element, err := service.Element().GetById(ctx, gconv.Uint64(step.TargetElementId))
		if err != nil {
			return out, err
		}
		out.Element = element
	}

	return out, nil
}

// Create 创建 Step
func (s *sStep) Create(ctx context.Context, in model.StepCreateInput) (stepId int64, err error) {
	var (
		m = dao.Step.Ctx(ctx)
	)
	stepId, err = m.Data(in).InsertAndGetId()
	if err != nil {
		return stepId, err
	}
	return stepId, nil
}

// Update Step Update
func (s *sStep) Update(ctx context.Context, in *entity.Step) (err error) {
	err = dao.Step.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.Step.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.Step.Columns().Id).
			Where(dao.Step.Columns().Id, in.Id).
			Update()
		return err
	})
	return err
}

// remove Step
func (s *sStep) Delete(ctx context.Context, id int) (err error) {
	var (
		m = dao.Step.Ctx(ctx)
	)
	_, err = m.Delete(dao.Step.Columns().Id, id)
	return err
}
