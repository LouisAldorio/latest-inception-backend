package service

import (
	"context"
	"log"
	"myapp/config"
	"myapp/graph/model"
	"myapp/tool"

	"github.com/vektah/gqlparser/gqlerror"
)

func Login(ctx context.Context, input model.LoginRequest) (*model.LoginResponse, error) {
	var response model.LoginResponse
	var hashedPassword string

	db, sql := config.ConnectDB()
	defer sql.Close()

	res := db.Table("user").Select("password").Where("username = ?", input.Username).Find(&hashedPassword)
	if res.RowsAffected < 1 {
		return nil, gqlerror.Errorf("Username or password incorrect")
	}

	if isValid := tool.CheckPassword(hashedPassword, input.Password); !isValid {
		return nil, gqlerror.Errorf("Username or password incorrect")
	}

	user, err := UserGetByUsername(ctx, input.Username)
	if err != nil {
		return nil, err
	}

	accessToken, err := tool.CreateToken(user)
	if err != nil {
		return nil, err
	}

	response = model.LoginResponse{
		AccessToken: accessToken,
		User:        user,
	}

	return &response, nil
}

func Register(ctx context.Context, input model.NewUser) (*model.LoginResponse, error) {
	if isValid, err := tool.ValidateInput(ctx, input); !isValid {
		return nil, err
	}

	var response model.LoginResponse

	if ok := isUsernameAvailable(ctx, input.Username); !ok {
		return nil, gqlerror.Errorf("Username unavailable")
	}

	user, err := UserCreate(ctx, input)
	if err != nil {
		return nil, err
	}

	accessToken, err := tool.CreateToken(user)
	if err != nil {
		return nil, err
	}

	response = model.LoginResponse{
		AccessToken: accessToken,
		User:        user,
	}

	return &response, nil
}

func isUsernameAvailable(ctx context.Context, username string) bool {
	var count int64

	db, sql := config.ConnectDB()
	defer sql.Close()

	if err := db.Table("user").Where("username = ?", username).Count(&count).Error; err != nil {
		log.Println(err)
		return false
	}

	if int(count) > 0 {
		return false
	}

	return true
}
