package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"myapp/graph/generated"
	"myapp/graph/model"
)

func (r *mutationResolver) User(ctx context.Context) (*model.UserOps, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Commodity(ctx context.Context) (*model.CommodityOps, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Schedule(ctx context.Context) (*model.ScheduleOps, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Friends(ctx context.Context) (*model.FriendOps, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserByUsername(ctx context.Context, username string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Comodities(ctx context.Context, limit *int, page *int) (*model.ComodityPagination, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UsersByRole(ctx context.Context, role string) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ScheduleByUser(ctx context.Context) ([]*model.Schedule, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FriendList(ctx context.Context) ([]*model.Friend, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
