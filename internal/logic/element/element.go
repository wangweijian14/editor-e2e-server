package element

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

type sElement struct{}

func init() {
	service.RegisterElement(New())
}
func New() *sElement {
	return &sElement{}
}

// GetList 查询Elements列表
func (s *sElement) GetList(ctx context.Context, in model.ElementListInput) (out *model.ElementListOutput, err error) {
	var (
		m = dao.Element.Ctx(ctx)
	)
	out = &model.ElementListOutput{
		Page:  in.Page,
		Limit: in.Limit,
	}

	// 名称模糊
	if in.ElementName != "" {
		m = m.WhereLike(dao.Element.Columns().Name, in.ElementName)
	}
	// 描述模糊
	if in.Description != "" {
		m = m.WhereLike(dao.Element.Columns().Description, in.Description)
	}

	if in.Father != 0 {
		m = m.Where(dao.Element.Columns().Father, in.Father)
	}

	// pageId 过滤
	if in.PageId != 0 {
		m = m.Where(dao.Element.Columns().PageId, in.PageId)
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Limit)
	// 排序方式
	switch in.Sort {
	case consts.PageSortActive:
		listModel = listModel.OrderAsc(dao.Element.Columns().Id)

	default:
		listModel = listModel.OrderDesc(dao.Element.Columns().Id)
	}

	// 执行查询
	var list []*entity.Element
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

// GetById 根据ID查询Element
func (s *sElement) GetById(ctx context.Context, in uint64) (out *model.ElementOutput, err error) {
	var (
		m = dao.Element.Ctx(ctx)
	)
	out = &model.ElementOutput{}
	element := &entity.Element{}
	err = m.Where(dao.Element.Columns().Id, in).Scan(element)
	out.Element = element
	// 没有数据
	if err != nil {
		return nil, err
	}

	if element.PageId != 0 {
		page, err := service.Page().GetById(ctx, gconv.Uint64(element.PageId))
		if err != nil {
			return out, err
		}
		out.Page = page.Page
	}
	// 查询父元素
	if element.Father != 0 {
		fatherE, err := service.Element().GetById(ctx, gconv.Uint64(element.Father))
		if err != nil {
			return out, err
		}
		out.FatherElement = fatherE.Element
	}

	return out, nil
}

// Create 创建 Element
func (s *sElement) Create(ctx context.Context, in model.ElementCreateInput) (elementId int64, err error) {
	var (
		m = dao.Element.Ctx(ctx)
	)
	elementId, err = m.Data(in).InsertAndGetId()
	if err != nil {
		return elementId, err
	}
	return elementId, nil
}

// Update Element Update
func (s *sElement) Update(ctx context.Context, in *entity.Element) (err error) {
	err = dao.Element.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.Element.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.Element.Columns().Id).
			Where(dao.Element.Columns().Id, in.Id).
			Update()
		return err
	})
	return err
}

// Update Element Update
func (s *sElement) Delete(ctx context.Context, id int) (err error) {
	var (
		m = dao.Element.Ctx(ctx)
	)
	_, err = m.Delete(dao.Element.Columns().Id, id)
	return err
}
