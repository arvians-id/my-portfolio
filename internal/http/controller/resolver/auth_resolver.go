package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/util"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func (m mutationResolver) Login(ctx context.Context, input model.AuthLoginRequest) (*model.AuthLoginResponse, error) {
	err := util.ValidateStruct(ctx, input)
	if err != nil {
		return nil, err
	}

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

func (m mutationResolver) Logout(ctx context.Context) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m mutationResolver) Register(ctx context.Context, input model.AuthRegisterRequest) (*model.User, error) {
	err := util.ValidateStruct(ctx, input)
	if err != nil {
		return nil, err
	}

	user, err := m.UserService.Create(ctx, &entity.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Pronouns: input.Pronouns,
		Country:  input.Country,
		JobTitle: input.JobTitle,
	})
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Bio:       user.Bio,
		Pronouns:  user.Pronouns,
		Country:   user.Country,
		JobTitle:  user.JobTitle,
		Image:     user.Image,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}
