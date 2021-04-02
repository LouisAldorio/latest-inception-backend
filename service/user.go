package service

import (
	"context"
	"myapp/auth"
	"myapp/config"
	"myapp/graph/model"
	"myapp/tool"
	"time"
)

func UserGetByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User

	db, sql := config.ConnectDB()
	defer sql.Close()

	err := db.Table("user").Where("username = ?", username).First(&user).Error

	return &user, err
}

func UserCreate(ctx context.Context, input model.NewUser) (*model.User, error) {
	var user = model.User{
		Username:  input.Username,
		Email:     input.Email,
		Role:      input.Role,
		Whatsapp:  input.Whatsapp,
		Password:  tool.HashPassword(input.Password),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db, sql := config.ConnectDB()
	defer sql.Close()

	err := db.Table("user").Create(&user).Error

	return &user, err
}

func UserGetByID(ctx context.Context, id int) (*model.User, error) {
	var user model.User

	db, sql := config.ConnectDB()
	defer sql.Close()

	err := db.Table("user").Where("id = ?", id).First(&user).Error

	return &user, err
}

func UserGetByIDs(ctx context.Context, ids []int) ([]*model.User, error) {
	var users []*model.User

	db, sql := config.ConnectDB()
	defer sql.Close()

	err := db.Table("user").Where("id IN (?)", ids).Find(&users).Error

	return users, err
}

func UserGetByRole(ctx context.Context, role string) ([]*model.User, error) {
	var users []*model.User

	db, sql := config.ConnectDB()
	defer sql.Close()

	err := db.Table("user").Where("role = ?", role).Find(&users).Error

	return users, err
}

func UserUpdate(ctx context.Context, input model.EditUser) (*model.User, error) {
	var userID = auth.ForContext(ctx).ID

	db, sql := config.ConnectDB()
	defer sql.Close()

	image, err := ImageCreate(ctx, model.NewImage{
		Path: "",
		Link: input.Avatar,
	})
	if err != nil {
		return nil, err
	}

	user := map[string]interface{}{
		"email":    input.Email,
		"whatsapp": input.Whatsapp,
		"avatar":   image.ID,
	}

	if err := db.Table("user").Where("id = ?", userID).Updates(&user).Error; err != nil {
		return nil, err
	}

	LookingForCreate(ctx, userID, input.LookingFor)

	return UserGetByID(ctx, userID)
}
