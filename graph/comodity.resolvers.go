package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *commodityOpsResolver) Create(ctx context.Context, obj *model.CommodityOps, input model.NewComodity) (*model.Comodity, error) {
	return service.ComodityCreate(ctx, input)
}

func (r *commodityOpsResolver) Update(ctx context.Context, obj *model.CommodityOps, input model.NewComodity) (*model.Comodity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *comodityResolver) Image(ctx context.Context, obj *model.Comodity) ([]*string, error) {
	var img1 = "https://images.unsplash.com/photo-1615932114921-0ffc1852c5d5?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=1000&q=80"
	var img2 = "https://images.unsplash.com/photo-1615910388452-fdad95ebb038?ixlib=rb-1.2.1&ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&auto=format&fit=crop&w=1000&q=80"
	var img3 = "https://images.unsplash.com/photo-1615886490002-ae8172eea2de?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=1000&q=80"
	var img4 = "https://images.unsplash.com/photo-1615729947596-a598e5de0ab3?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=1000&q=80"
	var img5 = "https://images.unsplash.com/photo-1615879965520-cde89ea80726?ixlib=rb-1.2.1&ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&auto=format&fit=crop&w=1000&q=80"
	var img6 = "https://images.unsplash.com/photo-1615839031296-a23180ec04ba?ixlib=rb-1.2.1&ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&auto=format&fit=crop&w=1000&q=80"
	var temp = []*string{&img1, &img2, &img3, &img4, &img5, &img6}
	return temp, nil
}

func (r *comodityResolver) User(ctx context.Context, obj *model.Comodity) (*model.User, error) {
	return service.UserGetByID(ctx, obj.UserID)
}

func (r *comodityPaginationResolver) TotalItem(ctx context.Context, obj *model.ComodityPagination) (int, error) {
	return service.ComodityGetCount(ctx, obj.Limit, obj.Page)
}

func (r *comodityPaginationResolver) Nodes(ctx context.Context, obj *model.ComodityPagination) ([]*model.Comodity, error) {
	return service.ComodityGetList(ctx, obj.Limit, obj.Page)
}

func (r *comodityWithCategoryResolver) TotalItem(ctx context.Context, obj *model.ComodityWithCategory) (int, error) {
	return service.ComodityGetCountByCategoryID(ctx, obj.CategoryID, obj.Limit, obj.Page)
}

func (r *comodityWithCategoryResolver) Nodes(ctx context.Context, obj *model.ComodityWithCategory) ([]*model.Comodity, error) {
	return service.ComodityGetByCategoryID(ctx, obj.CategoryID, obj.Limit, obj.Page)
}

// CommodityOps returns generated.CommodityOpsResolver implementation.
func (r *Resolver) CommodityOps() generated.CommodityOpsResolver { return &commodityOpsResolver{r} }

// Comodity returns generated.ComodityResolver implementation.
func (r *Resolver) Comodity() generated.ComodityResolver { return &comodityResolver{r} }

// ComodityPagination returns generated.ComodityPaginationResolver implementation.
func (r *Resolver) ComodityPagination() generated.ComodityPaginationResolver {
	return &comodityPaginationResolver{r}
}

// ComodityWithCategory returns generated.ComodityWithCategoryResolver implementation.
func (r *Resolver) ComodityWithCategory() generated.ComodityWithCategoryResolver {
	return &comodityWithCategoryResolver{r}
}

type commodityOpsResolver struct{ *Resolver }
type comodityResolver struct{ *Resolver }
type comodityPaginationResolver struct{ *Resolver }
type comodityWithCategoryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *comodityResolver) CategoryID(ctx context.Context, obj *model.Comodity) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
