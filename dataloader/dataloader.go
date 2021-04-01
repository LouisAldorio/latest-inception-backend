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
	ImagesGetByComodityIDLoader *ImagesLoaderByComodityID
}

func LoaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ldrs := loaders{}
		wait := 1 * time.Millisecond

		ldrs.ImagesGetByComodityIDLoader = &ImagesLoaderByComodityID{
			wait:     wait,
			maxBatch: 100,
			fetch: func(comodityIDs []int) ([][]*string, []error) {
				res, _ := service.ImageGetByComodityIDs(r.Context(), comodityIDs)
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
			},
		}

		dlCtx := context.WithValue(r.Context(), ctxKey, ldrs)
		next.ServeHTTP(w, r.WithContext(dlCtx))
	})
}

func CtxLoaders(ctx context.Context) loaders {
	return ctx.Value(ctxKey).(loaders)
}
