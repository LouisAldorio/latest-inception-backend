package service

import (
	"context"
	"myapp/config"
	"myapp/graph/model"
)

func CategoryGetList(ctx context.Context) ([]*model.Category, error) {
	var categories []*model.Category

	db, sql := config.ConnectDB()
	defer sql.Close()

	err := db.Table("category").Find(&categories).Error

	return categories, err
}