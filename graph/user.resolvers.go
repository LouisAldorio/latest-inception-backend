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

func (r *userResolver) Avatar(ctx context.Context, obj *model.User) (*string, error) {
	var temp = "https://images.unsplash.com/photo-1615943168243-5b2503679e47?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=1000&q=80"
	return &temp, nil
}

func (r *userResolver) Friends(ctx context.Context, obj *model.User) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) LookingFor(ctx context.Context, obj *model.User) ([]string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Comodity(ctx context.Context, obj *model.User) ([]*model.Comodity, error) {
	return service.ComodityGetByUserID(ctx, obj.ID)
}

func (r *userOpsResolver) Register(ctx context.Context, obj *model.UserOps, input model.NewUser) (*model.LoginResponse, error) {
	return service.Register(ctx, input)
}

func (r *userOpsResolver) Login(ctx context.Context, obj *model.UserOps, input model.LoginRequest) (*model.LoginResponse, error) {
	return service.Login(ctx, input)
}

func (r *userOpsResolver) Update(ctx context.Context, obj *model.UserOps, input model.EditUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userOpsResolver) DeleteUser(ctx context.Context, obj *model.UserOps, username string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// UserOps returns generated.UserOpsResolver implementation.
func (r *Resolver) UserOps() generated.UserOpsResolver { return &userOpsResolver{r} }

type userResolver struct{ *Resolver }
type userOpsResolver struct{ *Resolver }
