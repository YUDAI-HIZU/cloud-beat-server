package middleware

import (
	"app/infrastructure/auth"
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		client := auth.NewAuthClient(ctx)
		authorization := c.Request.Header.Get("Authorization")
		idToken := strings.Replace(authorization, "Bearer ", "", 1)
		token, err := client.VerifyIDToken(ctx, idToken)
		if err != nil {
			ctx = context.WithValue(
				ctx,
				"authError",
				&gqlerror.Error{
					Message: err.Error(),
					Extensions: map[string]interface{}{
						"status": http.StatusUnauthorized,
						"code":   "BAD_CREDENTIALS",
					},
				},
			)
			c.Request = c.Request.WithContext(ctx)
			c.Next()
			return
		}
		id := int(token.Claims["ID"].(float64))
		ctx = context.WithValue(ctx, "ID", id)
		c.Set("ID", id)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
