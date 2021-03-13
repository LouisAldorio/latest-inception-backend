package service

import (
	"context"
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

func UserGetByRole(ctx context.Context, role string) ([]*model.User, error) {
	var users []*model.User

	db, sql := config.ConnectDB()
	defer sql.Close()

	err := db.Table("user").Where("role = ?", role).Find(&users).Error

	return users, err
}
