package dataloader

import (
	"context"
	"myapp/graph/model"
	"myapp/service"
	"net/http"
	"time"
)

type key struct {
	name string
}

var ctxKey = key{"inceptionCtx"}

type loaders struct {
	ImagesGetByComodityIDLoader   *ImagesLoaderByComodityID
	ComodityGetByCategoryIDLoader *ComodityLoaderByCategoryID
	UserGetByIDLoader             *UserLoaderByID
}

func LoaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ldrs := loaders{}
		wait := 1 * time.Millisecond

		ldrs.ImagesGetByComodityIDLoader = &ImagesLoaderByComodityID{
			wait:     wait,
			maxBatch: 100,
			fetch:    imageByComodtiyIDLoader,
		}

		ldrs.ComodityGetByCategoryIDLoader = &ComodityLoaderByCategoryID{
			wait:     wait,
			maxBatch: 100,
			fetch:    comodityByCategoryIDLoader,
		}

		ldrs.UserGetByIDLoader = &UserLoaderByID{
			wait:     wait,
			maxBatch: 100,
			fetch:    userByIDLoader,
		}

		dlCtx := context.WithValue(r.Context(), ctxKey, ldrs)
		next.ServeHTTP(w, r.WithContext(dlCtx))
	})
}

func CtxLoaders(ctx context.Context) loaders {
	return ctx.Value(ctxKey).(loaders)
}

func userByIDLoader(ids []int) ([]*model.User, []error) {
	var ctx = context.Background()

	res, _ := service.UserGetByIDs(ctx, ids)
	userByID := map[int]*model.User{}
	errors := make([]error, len(ids))

	for _, v := range res {
		userByID[v.ID] = v
	}

	users := make([]*model.User, len(ids))

	for i, id := range ids {
		users[i] = userByID[id]
		i++
	}

	return users, errors
}

func comodityByCategoryIDLoader(categoryIDs []int) ([][]*model.Comodity, []error) {
	var ctx = context.Background()

	res, _ := service.ComodityGetByCategoryIDs(ctx, categoryIDs)
	comodityByCategoryID := map[int][]*model.Comodity{}
	errors := make([]error, len(categoryIDs))

	for _, v := range res {
		comodityByCategoryID[v.CategoryID] = append(comodityByCategoryID[v.CategoryID], v)
	}

	comodities := make([][]*model.Comodity, len(categoryIDs))

	for i, categoryID := range categoryIDs {
		comodities[i] = comodityByCategoryID[categoryID]
		i++
	}

	return comodities, errors
}

func imageByComodtiyIDLoader(comodityIDs []int) ([][]*string, []error) {
	var ctx = context.Background()

	res, _ := service.ImageGetByComodityIDs(ctx, comodityIDs)
	imageByComodityID := map[int][]*model.ImageWithComodityID{}
	errors := make([]error, len(comodityIDs))

	for _, v := range res {
		imageByComodityID[v.ComodityID] = append(imageByComodityID[v.ComodityID], v)
	}

	images := make([][]*string, len(comodityIDs))

	for i, comodityID := range comodityIDs {
		var temp []*string
		for _, v := range imageByComodityID[comodityID] {
			temp = append(temp, &v.Link)
		}

		images[i] = temp
		i++
	}

	return images, errors
}
