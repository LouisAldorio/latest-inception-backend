package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"myapp/graph/generated"
	"myapp/graph/model"
)

func (r *scheduleResolver) InvolvedUserID(ctx context.Context, obj *model.Schedule) ([]*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *scheduleResolver) InvolvedUsers(ctx context.Context, obj *model.Schedule) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *scheduleOpsResolver) Create(ctx context.Context, obj *model.ScheduleOps, input model.NewSchedule) (*model.Schedule, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *scheduleOpsResolver) Update(ctx context.Context, obj *model.ScheduleOps, input model.EditSchedule) (*model.Schedule, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *scheduleOpsResolver) Delete(ctx context.Context, obj *model.ScheduleOps, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Schedule returns generated.ScheduleResolver implementation.
func (r *Resolver) Schedule() generated.ScheduleResolver { return &scheduleResolver{r} }

// ScheduleOps returns generated.ScheduleOpsResolver implementation.
func (r *Resolver) ScheduleOps() generated.ScheduleOpsResolver { return &scheduleOpsResolver{r} }

type scheduleResolver struct{ *Resolver }
type scheduleOpsResolver struct{ *Resolver }
