package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
	"time"
)

func (r *userResolver) Avatar(ctx context.Context, obj *model.User) (*string, error) {
	var temp = "https://via.placeholder.com/500"
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *userResolver) CreatedAt(ctx context.Context, obj *model.User) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *userResolver) UpdatedAt(ctx context.Context, obj *model.User) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *userResolver) Password(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *userResolver) Products(ctx context.Context, obj *model.User) ([]*model.Comodity, error) {
	panic(fmt.Errorf("not implemented"))
}
