package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"myapp/dataloader"
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
	var img1 = "https://via.placeholder.com/500"
	var temp = []*string{&img1}

	images, err := dataloader.CtxLoaders(ctx).ImagesGetByComodityIDLoader.Load(obj.ID)

	if len(images) < 1 {
		return temp, err
	}

	return images, err
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
