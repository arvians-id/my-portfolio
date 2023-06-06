package middleware

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

var SecretKey = []byte("secret")

func NewJWTMiddlewareGraphQL(ctx context.Context, obj interface{}, next graphql.Resolver, isLoggedIn bool) (interface{}, error) {
	fiberCtx := FiberContext(ctx)

	if !isLoggedIn {
		return next(ctx)
	}

	authorizationHeader := fiberCtx.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		return nil, errors.New("invalid authorization header")
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx = context.WithValue(ctx, "id", claims["id"])
		ctx = context.WithValue(ctx, "email", claims["email"])
	} else {
		return nil, errors.New("invalid JWT token")
	}

	return next(ctx)
}
