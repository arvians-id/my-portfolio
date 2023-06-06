package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func (m mutationResolver) Login(ctx context.Context, input model.AuthLoginRequest) (*model.AuthLoginResponse, error) {
	validateLogin, err := m.UserService.ValidateLogin(ctx, input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	claims := jwt.Claims(jwt.MapClaims{
		"id":    validateLogin.ID,
		"email": validateLogin.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return &model.AuthLoginResponse{
		Token: signedToken,
	}, nil
}
