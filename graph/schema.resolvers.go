package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"myapp/auth"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *mutationResolver) User(ctx context.Context) (*model.UserOps, error) {
	return &model.UserOps{}, nil
}

func (r *mutationResolver) Commodity(ctx context.Context) (*model.CommodityOps, error) {
	return &model.CommodityOps{}, nil
}

func (r *mutationResolver) Schedule(ctx context.Context) (*model.ScheduleOps, error) {
	return &model.ScheduleOps{}, nil
}

func (r *mutationResolver) Friends(ctx context.Context) (*model.FriendOps, error) {
	return &model.FriendOps{}, nil
}

func (r *queryResolver) UserByUsername(ctx context.Context, username string) (*model.User, error) {
	return service.UserGetByUsername(ctx, username)
}

func (r *queryResolver) Comodities(ctx context.Context, limit *int, page *int) (*model.ComodityPagination, error) {
	return &model.ComodityPagination{
		Limit: limit,
		Page:  page,
	}, nil
}

func (r *queryResolver) ComoditiesWithCategories(ctx context.Context, limit *int, page *int) ([]*model.ComodityWithCategory, error) {
	return service.ComodityWithCategoryGet(ctx, limit, page)
}

func (r *queryResolver) ComoditiesByCategory(ctx context.Context, categoryID int) ([]*model.Comodity, error) {
	return service.ComodityGetByCategoryID(ctx, categoryID, nil, nil)
}

func (r *queryResolver) UsersByRole(ctx context.Context, role string) ([]*model.User, error) {
	return service.UserGetByRole(ctx, role)
}

func (r *queryResolver) ScheduleByUser(ctx context.Context) ([]*model.Schedule, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FriendList(ctx context.Context) ([]*model.Friend, error) {
	var userID = auth.ForContext(ctx).ID
	return service.FriendGetByUserID(ctx, userID)
}

func (r *queryResolver) CategoryList(ctx context.Context) ([]*model.Category, error) {
	return service.CategoryGetList(ctx)
}

func (r *queryResolver) HelloWorld(ctx context.Context) (string, error) {
	return "Last Attempt Succes!",nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
