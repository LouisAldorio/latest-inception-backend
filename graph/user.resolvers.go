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

func (r *userResolver) Image(ctx context.Context, obj *model.User) (*model.Image, error) {
	return service.ImageGetByID(ctx, obj.Avatar)
}

func (r *userResolver) Friends(ctx context.Context, obj *model.User) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) LookingFor(ctx context.Context, obj *model.User) ([]*string, error) {
	return service.LookingForGetByUserID(ctx, obj.ID)
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
	return service.UserUpdate(ctx, input)
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
