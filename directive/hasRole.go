package directive

import (
	"context"
	"fmt"
	"myapp/auth"

	"github.com/99designs/gqlgen/graphql"
)

func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role string) (interface{}, error) {
	// username := auth.ForContext(ctx)
	// if username != role {
	// 	return nil, fmt.Errorf("Access denied")
	// }
	fmt.Println(role)
	if !auth.ForContext(ctx).HasRole(role) {
		return nil, fmt.Errorf("Access denied")
	}

	return next(ctx)
}
