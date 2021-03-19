package service

import (
	"context"
	"myapp/auth"
	"myapp/config"
	"myapp/graph/model"
	"time"
)

func ComodityGetCount(ctx context.Context, limit, page *int) (int, error) {
	var count int64

	db, sql := config.ConnectDB()
	defer sql.Close()

	err := db.Table("comodity").Count(&count).Error

	if limit != nil && page != nil {
		offset := *limit * (*page - 1)
		mod := (int(count) - offset) % *limit

		if mod == 0 {
			return *limit, err
		}

		return mod, err
	}

	return int(count), err
}

func ComodityGetList(ctx context.Context, limit, page *int) ([]*model.Comodity, error) {
	var comodities []*model.Comodity

	db, sql := config.ConnectDB()
	defer sql.Close()

	tx := db.Table("comodity")

	if limit != nil && page != nil {
		offset := *limit * (*page - 1)

		tx.Limit(*limit).Offset(offset)
	}

	err := tx.Find(&comodities).Error

	return comodities, err
}

func ComodityCreate(ctx context.Context, input model.NewComodity) (*model.Comodity, error) {
	var user = auth.ForContext(ctx)

	var comodity = model.Comodity{
		Name:        input.Name,
		UnitPrice:   input.UnitPrice,
		UnitType:    input.UnitType,
		MinPurchase: input.MinPurchase,
		Description: input.Description,
		CategoryID:  input.CategoryID,
		UserID:      user.ID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	db, sql := config.ConnectDB()
	defer sql.Close()

	err := db.Table("comodity").Omit("image").Create(&comodity).Error

	return &comodity, err
}

func ComodityGetByUserID(ctx context.Context, userID int) ([]*model.Comodity, error) {
	var comodities []*model.Comodity

	db, sql := config.ConnectDB()
	defer sql.Close()

	err := db.Table("comodity").Where("user_id = ?", userID).Find(&comodities).Error

	return comodities, err
}

func ComodityWithCategoryGet(ctx context.Context, limit, page *int) ([]*model.ComodityWithCategory, error) {
	var comodities []*model.ComodityWithCategory

	categories, err := CategoryGetList(ctx)
	if err != nil {
		return nil, err
	}

	for _, category := range categories {
		comodity := model.ComodityWithCategory{
			Limit:      limit,
			Page:       page,
			CategoryID: category.ID,
			Category:   category,
		}

		comodities = append(comodities, &comodity)
	}

	return comodities, err
}

func ComodityGetByCategoryID(ctx context.Context, categoryID int, limit, page *int) ([]*model.Comodity, error) {
	var comodities []*model.Comodity

	db, sql := config.ConnectDB()
	defer sql.Close()

	tx := db.Table("comodity").Where("category_id = ?", categoryID)

	if limit != nil && page != nil {
		offset := *limit * (*page - 1)

		tx.Limit(*limit).Offset(offset)
	}

	err := tx.Find(&comodities).Error

	return comodities, err
}

func ComodityGetCountByCategoryID(ctx context.Context, categoryID int, limit, page *int) (int, error) {
	var count int64

	db, sql := config.ConnectDB()
	defer sql.Close()

	err := db.Table("comodity").Where("category_id = ?", categoryID).Count(&count).Error

	if limit != nil && page != nil {
		offset := *limit * (*page - 1)
		mod := (int(count) - offset) % *limit

		if mod == 0 {
			return *limit, err
		}

		return mod, err
	}

	return int(count), err
}

// func ComodityGetByCategoryID(ctx context.Context, categoryID int) ([]*model.Comodity, error) {
// 	var comodities []*model.Comodity

// 	db, sql := config.ConnectDB()
// 	defer sql.Close()

// 	err := db.Table("comodity").Where("category_id = ?", categoryID).Find(&comodities).Error

// 	return comodities, err
// }
