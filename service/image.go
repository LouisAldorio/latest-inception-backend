package service

import (
	"context"
	"myapp/config"
	"myapp/graph/model"
)

func ImageCreate(ctx context.Context, input model.NewImage) (*model.Image, error) {
	var image = model.Image{
		Path: input.Path,
		Link: input.Link,
	}

	db, sql := config.ConnectDB()
	defer sql.Close()

	err := db.Table("image").Create(&image).Error

	return &image, err
}

func ImageGetByID(ctx context.Context, id *int) (*model.Image, error) {
	var image model.Image

	if id == nil {
		return &model.Image{
			ID:   0,
			Path: "",
			Link: "https://images.unsplash.com/photo-1615943168243-5b2503679e47?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=1000&q=80",
		}, nil
	}

	db, sql := config.ConnectDB()
	defer sql.Close()

	err := db.Table("image").First(&image, id).Error

	return &image, err
}
