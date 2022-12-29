package cases

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

type sCases struct{}

func init() {
	service.RegisterCases(New())
}
func New() *sCases {
	return &sCases{}
}

// GetList 查询cases列表
func (s *sCases) GetList(ctx context.Context, in model.CasesListInput) (out *model.CasesListOutput, err error) {
	var (
		m = dao.Cases.Ctx(ctx)
	)
	out = &model.CasesListOutput{
		Page:  in.Page,
		Limit: in.Limit,
	}

	// 过滤历史记录ID
	if in.HisId != 0 {
		m = m.Where(dao.Cases.Columns().HisId, in.HisId)
	}
	// 描述模糊
	if in.Description != "" {
		m = m.WhereLike(dao.Cases.Columns().Description, in.Description)
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Limit)
	// 排序方式
	switch in.Sort {
	case consts.PageSortActive:
		listModel = listModel.OrderAsc(dao.Cases.Columns().Id)

	default:
		listModel = listModel.OrderDesc(dao.Cases.Columns().Id)
	}

	// 执行查询
	var list []*entity.Cases
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

// GetById 根据ID查询Cases
func (s *sCases) GetById(ctx context.Context, in uint64) (out *model.CasesOutput, err error) {
	var (
		m = dao.Cases.Ctx(ctx)
	)
	out = &model.CasesOutput{}
	cases := &entity.Cases{}
	err = m.Where(dao.Cases.Columns().Id, in).Scan(cases)
	out.Cases = cases
	// 没有数据
	if err != nil {
		return nil, err
	}

	// 查询此条用例中的全部测试步骤
	stepsOut, err := service.CaseStep().GetList(ctx, model.CaseStepListInput{
		CaseId: cases.Id,
		Sort:   consts.PageSortScore,
	})

	if err != nil {
		return nil, err
	}

	// out.CaseStepOutput = steps.List

	// 匹配全部步骤信息
	for _, step := range stepsOut.List {
		stepId := gconv.Uint64(step.Id)
		if stepId != 0 {
			caseStep, err := service.CaseStep().GetById(ctx, stepId)
			if err != nil {
				return out, err
			}
			if caseStep != nil {
				out.CaseStepOutput = append(out.CaseStepOutput, caseStep)
			}
		}

	}

	return out, nil
}

// Create 创建 Cases
func (s *sCases) Create(ctx context.Context, in model.CasesCreateInput) (CasesId int64, err error) {
	var (
		m = dao.Cases.Ctx(ctx)
	)
	CasesId, err = m.Data(in).InsertAndGetId()
	if err != nil {
		return CasesId, err
	}
	return CasesId, nil
}

// Update Cases Update
func (s *sCases) Update(ctx context.Context, in *entity.Cases) (err error) {
	err = dao.Cases.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.Cases.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.Cases.Columns().Id).
			Where(dao.Cases.Columns().Id, in.Id).
			Update()
		return err
	})
	return err
}
