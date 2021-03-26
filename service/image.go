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

func ImageCreateBatch(ctx context.Context, input []*model.NewImage) ([]*model.Image, error) {
	var images []*model.Image

	db, sql := config.ConnectDB()
	defer sql.Close()

	for _, k := range input {
		image := model.Image{
			Path: k.Path,
			Link: k.Link,
		}

		images = append(images, &image)
	}

	err := db.Table("image").Create(&images).Error

	return images, err
}

func ComodityImageCreate(ctx context.Context, input []*model.NewImage, comodityID int) ([]*model.Image, error) {
	var comodityImages []*model.ComodityImage

	db, sql := config.ConnectDB()
	defer sql.Close()

	images, err := ImageCreateBatch(ctx, input)
	if err != nil {
		return nil, err
	}

	for _, k := range images {
		comodityImage := model.ComodityImage{
			ComodityID: comodityID,
			ImageID:    k.ID,
		}

		comodityImages = append(comodityImages, &comodityImage)
	}

	err = db.Table("comodity_image").Create(&comodityImages).Error
	if err != nil {
		return nil, err
	}

	return images, nil
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

func ImageGetByComodityID(ctx context.Context, comodityID int) ([]*string, error) {
	var imageID []int
	var image []*string

	db, sql := config.ConnectDB()
	defer sql.Close()

	err := db.Table("comodity_image").Select("image_id").Where("comodity_id = ?", comodityID).Find(&imageID).Error
	if err != nil {
		return nil, err
	}

	err = db.Table("image").Select("link").Where("id IN (?)", imageID).Find(&image).Error
	if err != nil {
		return nil, err
	}

	return image, nil
}
