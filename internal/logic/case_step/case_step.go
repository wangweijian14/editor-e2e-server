package case_step

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

type sCaseStep struct{}

func init() {
	service.RegisterCaseStep(New())
}
func New() *sCaseStep {
	return &sCaseStep{}
}

// GetList 查询CaseSteps列表
func (s *sCaseStep) GetList(ctx context.Context, in model.CaseStepListInput) (out *model.CaseStepListOutput, err error) {
	var (
		m = dao.CaseStep.Ctx(ctx)
	)
	out = &model.CaseStepListOutput{
		Page:  in.Page,
		Limit: in.Limit,
	}

	// 名称模糊
	if in.StepId != 0 {
		m = m.Where(dao.CaseStep.Columns().StepId, in.StepId)
	}
	// 描述模糊
	if in.Input != "" {
		m = m.WhereLike(dao.CaseStep.Columns().Input, in.Input)
	}

	if in.AssertExpect != "" {
		m = m.Where(dao.CaseStep.Columns().AssertExpect, in.AssertExpect)
	}

	if in.CaseId != 0 {
		m = m.Where(
			dao.CaseStep.Columns().CaseId, in.CaseId,
		)
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Limit)
	// 排序方式
	switch in.Sort {
	case consts.PageSortActive:
		listModel = listModel.OrderAsc(dao.CaseStep.Columns().Id)

	case consts.PageSortScore:
		listModel = listModel.OrderAsc(dao.CaseStep.Columns().Score)

	default:
		listModel = listModel.OrderDesc(dao.CaseStep.Columns().Id)
	}

	// 执行查询
	var list []*entity.CaseStep
	if err := listModel.Scan(&list); err != nil {
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

// GetById 根据ID查询CaseStep
func (s *sCaseStep) GetById(ctx context.Context, in uint64) (out *model.CaseStepOutput, err error) {
	var (
		m = dao.CaseStep.Ctx(ctx)
	)
	out = &model.CaseStepOutput{}
	caseStep := &entity.CaseStep{}
	err = m.Where(dao.CaseStep.Columns().Id, in).Scan(caseStep)
	out.CaseStep = caseStep
	// 没有数据
	if err != nil {
		return nil, err
	}

	if caseStep.StepId != 0 {
		step, err := service.Step().GetById(ctx, gconv.Uint64(caseStep.StepId))
		if err != nil {
			return out, err
		}
		out.Step = step
	}

	return out, nil
}

// Create 创建 CaseStep
func (s *sCaseStep) Create(ctx context.Context, in model.CaseStepCreateInput) (CaseStepId int64, err error) {
	var (
		m = dao.CaseStep.Ctx(ctx)
	)
	CaseStepId, err = m.Data(in).InsertAndGetId()
	if err != nil {
		return CaseStepId, err
	}
	return CaseStepId, nil
}

// Update CaseStep Update
func (s *sCaseStep) Update(ctx context.Context, in *entity.CaseStep) (err error) {
	err = dao.CaseStep.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.CaseStep.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.CaseStep.Columns().Id).
			Where(dao.CaseStep.Columns().Id, in.Id).
			Update()
		return err
	})
	return err
}

func (s *sCaseStep) Delete(ctx context.Context, id int) (err error) {
	var (
		m = dao.CaseStep.Ctx(ctx)
	)
	_, err = m.Delete(dao.CaseStep.Columns().Id, id)
	return err
}
