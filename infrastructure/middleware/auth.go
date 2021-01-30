package middleware

import (
	"app/config"
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JwtSecret), nil
		})
		if err != nil {
			ctx = context.WithValue(
				ctx,
				"authError",
				&gqlerror.Error{
					Message: "トークンが有効ではありません",
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
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx = context.WithValue(
				ctx,
				"authError",
				&gqlerror.Error{
					Message: "トークンが有効ではありません",
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
		ctx = context.WithValue(ctx, "userID", claims["userID"])
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
