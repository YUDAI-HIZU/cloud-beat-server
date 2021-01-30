package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Authentication(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	err, ok := ctx.Value("authError").(*gqlerror.Error)
	if ok {
		return nil, err
	}
	return next(ctx)
}
