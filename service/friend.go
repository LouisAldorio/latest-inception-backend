package service

import (
	"context"
	"myapp/auth"
	"myapp/config"
	"myapp/graph/model"
)

func FriendGetByUserID(ctx context.Context, userID int) ([]*model.Friend, error) {
	var userIDs []int
	var firstUserIDs []int
	var secondUserIDs []int
	var friend []*model.Friend

	db, sql := config.ConnectDB()
	defer sql.Close()

	if err := db.Table("friendship").Distinct("second_user_id").Where("first_user_id = ?", userID).Find(&firstUserIDs).Error; err != nil {
		return nil, err
	}
	if err := db.Table("friendship").Distinct("first_user_id").Where("second_user_id = ?", userID).Find(&secondUserIDs).Error; err != nil {
		return nil, err
	}

	userIDs = append(firstUserIDs, secondUserIDs...)

	for _, userID := range userIDs {
		friend = append(friend, &model.Friend{UserID: userID})
	}

	return friend, nil
}

func FriendAdd(ctx context.Context, userID int) (*model.User, error) {
	var friendship map[string]interface{}
	var loggedUserID = auth.ForContext(ctx).ID

	db, sql := config.ConnectDB()
	defer sql.Close()

	friendship = map[string]interface{}{
		"first_user_id":  loggedUserID,
		"second_user_id": userID,
	}

	err := db.Table("friendship").Create(&friendship).Error

	user, _ := UserGetByID(ctx, userID)

	return user, err
}
