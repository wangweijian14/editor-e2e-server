package controller

import (
	"context"
	"wiki/internal/consts"
	"wiki/internal/model"
	"wiki/internal/service"

	_ "github.com/gogf/gf/v2/frame/g"

	v1 "wiki/api/v1"
)

var (
	Page = cPage{}
)

type cPage struct{}

func (c *cPage) Create(ctx context.Context, req *v1.PageCreateReq) (res *v1.PageCreateRes, err error) {
	out, err := service.Page().Create(ctx, model.PageCreateInput{
		Description: req.Description,
		Name:        req.Name, Url: req.Url},
	)
	if err != nil {
		return nil, err
	}
	return &v1.PageCreateRes{PageId: out.PageId}, nil
}

func (c *cPage) Update(ctx context.Context, req *v1.PageUpdateReq) (res *v1.PageUpdateRes, err error) {
	err = service.Page().Update(ctx, model.PageUpdateInput{
		Description: req.Description,
		Id:          req.Id,
		Name:        req.Name,
		Url:         req.Url},
	)
	r := &v1.PageUpdateRes{Ok: true}
	if err != nil {
		r.Ok = false
	}
	return r, err
}

func (c *cPage) FindById(ctx context.Context, req *v1.PageFindByIdReq) (res *v1.PageGetOneRes, err error) {
	out, err := service.Page().GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.PageGetOneRes{Page: out.Page}, nil
}

func (c *cPage) FindPageByIdWithElement(ctx context.Context, req *v1.PageFindPageByIdWithElementReq) (res *v1.PageGetOneWithElementsRes, err error) {
	out, err := service.Page().GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.PageGetOneWithElementsRes{
		Page:     out.Page,
		Elements: out.Elements,
	}, nil
}

func (c *cPage) FindList(ctx context.Context, req *v1.PageGetListReq) (res *v1.PageGetListRes, err error) {
	out, err := service.Page().GetList(ctx, model.PageGetListInput{
		Page:        req.Page,
		Description: req.Description,
		PageName:    req.PageName,
		Limit:       req.Limit,
		Sort:        req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &v1.PageGetListRes{
		Row:   out.List.Page,
		Total: out.Total,
		Limit: out.Limit,
		Page:  out.Page,
	}, nil
}

func (c *cPage) FindTree(ctx context.Context, req *v1.PageTreeReq) (res *v1.PageGetTreeRes, err error) {
	out, err := service.Page().GetList(ctx, model.PageGetListInput{
		Page:        req.Page,
		Description: req.Description,
		PageName:    req.PageName,
		Limit:       req.Limit,
		Sort:        consts.PageSortActive,
	})
	if err != nil {
		return nil, err
	}
	childless := make([]PageTree, 0)
	ps := make([]PageTree, 0)
	base := PageTree{
		Id:       0,
		Title:    "wiz-editor-page",
		Last:     false,
		ParentId: 9527,
	}
	for _, page := range out.List.Page {
		children := PageTree{
			Id:       page.Id,
			Title:    page.Description,
			Last:     true,
			ParentId: 0,
		}
		childless = append(childless, children)
	}
	base.Children = childless
	ps = append(ps, base)

	return &v1.PageGetTreeRes{
		Row:   ps,
		Total: out.Total,
		Limit: out.Limit,
		Page:  out.Page,
	}, nil
}

type PageTree struct {
	Id       int        `json:"id"`
	Title    string     `json:"title"`
	Last     bool       `json:"last"`
	ParentId int        `json:"parentId"`
	Children []PageTree `json:"children"`
}
