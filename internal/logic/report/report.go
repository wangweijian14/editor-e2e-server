package report

import (
	"context"
	"wiki/internal/consts"
	"wiki/internal/dao"
	"wiki/internal/model"
	"wiki/internal/model/entity"
	"wiki/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
)

type sReport struct{}

func init() {
	service.RegisterReport(New())
}
func New() *sReport {
	return &sReport{}
}

// GetList 查询caseReport列表
func (s *sReport) GetList(ctx context.Context, in model.CasesReportListInput) (out *model.CasesReportListOutput, err error) {
	var (
		m = dao.Casereport.Ctx(ctx)
	)
	out = &model.CasesReportListOutput{
		Page:  in.Page,
		Limit: in.Limit,
	}

	// 过滤历史记录ID
	if in.CaseId != 0 {
		m = m.Where(dao.Casereport.Columns().CaseId, in.CaseId)
	}
	// 描述模糊
	if in.Status != 0 {
		m = m.Where(dao.Casereport.Columns().Status, in.Status)
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Limit)
	// 排序方式
	switch in.Sort {
	case consts.PageSortActive:
		listModel = listModel.OrderAsc(dao.Casereport.Columns().RunTime)

	default:
		listModel = listModel.OrderDesc(dao.Casereport.Columns().RunTime)
	}

	// 执行查询
	var list []*entity.Casereport
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

// GetById 根据ID查询CaseReport
func (s *sReport) GetById(ctx context.Context, in uint64) (out *model.CasesReportOutput, err error) {
	var (
		m = dao.Casereport.Ctx(ctx)
	)
	out = &model.CasesReportOutput{}
	caseReport := &entity.Casereport{}
	err = m.Where(dao.Casereport.Columns().Id, in).Scan(caseReport)
	out.CaseReport = caseReport
	// 没有数据
	if err != nil {
		return nil, err
	}

	return out, nil
}

// Create 创建 caseReport
func (s *sReport) Create(ctx context.Context, in *entity.Casereport) (err error) {
	var (
		m = dao.Casereport.Ctx(ctx)
	)
	_, err = m.Data(in).Insert()
	if err != nil {
		return err
	}
	return nil
}

// Update caseReport Update
func (s *sReport) Update(ctx context.Context, in *entity.Casereport) (err error) {
	err = dao.Casereport.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.Cases.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.Casereport.Columns().Id).
			Where(dao.Casereport.Columns().Id, in.Id).
			Update()
		return err
	})
	return err
}
