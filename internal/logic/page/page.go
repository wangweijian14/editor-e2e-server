package page

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"wiki/internal/consts"
	"wiki/internal/dao"
	"wiki/internal/model"
	"wiki/internal/model/entity"
	"wiki/internal/service"
)

type sPage struct{}

func init() {
	service.RegisterPage(New())
}
func New() *sPage {
	return &sPage{}
}

// GetList 查询内容列表
func (s *sPage) GetList(ctx context.Context, in model.PageGetListInput) (out *model.PageGetListOutput, err error) {
	var (
		m = dao.Page.Ctx(ctx)
	)
	out = &model.PageGetListOutput{
		Page: in.Page,
		Limit: in.Limit,
	}

	// 名称模糊
	if in.PageName != "" {
		m = m.WhereLike(dao.Page.Columns().Name, in.PageName)
	}
	// 描述模糊
	if in.Description != ""{
		m = m.WhereLike(dao.Page.Columns().Description,in.Description)
	}
	// 分配查询
	listModel := m.Page(in.Page, in.Limit)
	// 排序方式
	switch in.Sort {
	case consts.PageSortActive:
		listModel = listModel.OrderAsc(dao.Page.Columns().Id)

	default:
		listModel = listModel.OrderDesc(dao.Page.Columns().Id)
	}

	// 执行查询
	var list []*entity.Page
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	out.Total, err = m.Count()
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.List.Page = list
	if err != nil {
		return out, err
	}
	return
}

// GetById 根据ID查询Page
func (s *sPage) GetById(ctx context.Context, in uint64) (out *model.PageOutputItem, err error) {
	var (
		m = dao.Page.Ctx(ctx)
	)
	out = &model.PageOutputItem{}
	p := &entity.Page{}
	err = m.Where(dao.Page.Columns().Id,in).Scan(p)
	out.Page = p
	// 没有数据
	if out.Page == nil{
		return out , nil
	}
	if err != nil {
		return out,err
	}

	// 关联 elements
	elements , err :=service.Element().GetList(ctx,model.ElementListInput{
		PageId: p.Id,
	})
	if err != nil {
		return out,err
	}
	out.Elements = elements.List

	return out,nil
}

// Create 创建Page
func (s *sPage) Create(ctx context.Context, in model.PageCreateInput) (out model.PageCreateOutput, err error) {
	var (
		m = dao.Page.Ctx(ctx)
	)
	id,err := m.Data(in).InsertAndGetId()
	if err != nil {
		return out,err
	}
	return model.PageCreateOutput{PageId: id},nil
}

// Update Page Update
func (s *sPage) Update(ctx context.Context, in model.PageUpdateInput) (err error) {
	err = dao.Page.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.Page.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.Page.Columns().Id).
			Where(dao.Page.Columns().Id, in.Id).
			Update()
		return err
	})
	return err
}