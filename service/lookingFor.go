package service

import (
	"context"
	"myapp/config"
	"myapp/graph/model"
)

func LookingForCreate(ctx context.Context, userID int, input []string) ([]*model.LookingFor, error) {
	var lookingFor []*model.LookingFor

	db, sql := config.ConnectDB()
	defer sql.Close()

	for _, k := range input {
		var temp = model.LookingFor{
			UserID: userID,
			Item:   k,
		}
		lookingFor = append(lookingFor, &temp)
	}

	err := db.Table("looking_for").Create(&lookingFor).Error

	return lookingFor, err
}

func LookingForGetByUserID(ctx context.Context, userID int) ([]*string, error) {
	var result []*string

	db, sql := config.ConnectDB()
	defer sql.Close()

	err := db.Table("looking_for").Select("item").Where("user_id = ?", userID).Find(&result).Error

	return result, err
}
