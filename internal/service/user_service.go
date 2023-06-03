package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"github.com/arvians-id/go-portfolio/util"
)

type UserServiceContract interface {
	FindAll(ctx context.Context) ([]*entity.User, error)
	FindByID(ctx context.Context, id int64) (*entity.User, error)
	Create(ctx context.Context, request *model.CreateUserRequest) (*entity.User, error)
	Update(ctx context.Context, request *model.UpdateUserRequest) (*entity.User, error)
	Delete(ctx context.Context, id int64) error
}

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return UserService{
		UserRepository: userRepository,
	}
}

func (service *UserService) FindAll(ctx context.Context) ([]*entity.User, error) {
	return service.UserRepository.FindAll(ctx)
}

func (service *UserService) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	return service.UserRepository.FindByID(ctx, id)
}

func (service *UserService) Create(ctx context.Context, request *model.CreateUserRequest) (*entity.User, error) {
	passwordHashed, err := util.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	var user entity.User
	user.Name = request.Name
	user.Email = request.Email
	user.Password = passwordHashed
	user.Pronouns = request.Pronouns
	user.Country = request.Country
	user.JobTitle = request.JobTitle

	return service.UserRepository.Create(ctx, &user)
}

func (service *UserService) Update(ctx context.Context, request *model.UpdateUserRequest) (*entity.User, error) {
	userCheck, err := service.FindByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	if request.Password != nil {
		hashedPassword, err := util.HashPassword(*request.Password)
		if err != nil {
			return nil, err
		}

		userCheck.Password = hashedPassword
	}

	userCheck.ID = request.ID
	userCheck.Name = *request.Name
	userCheck.Email = *request.Email
	userCheck.Bio = request.Bio
	userCheck.Pronouns = *request.Pronouns
	userCheck.Country = *request.Country
	userCheck.JobTitle = *request.JobTitle
	userCheck.Image = request.Image

	return service.UserRepository.Update(ctx, userCheck)
}

func (service *UserService) Delete(ctx context.Context, id int64) error {
	_, err := service.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return service.UserRepository.Delete(ctx, id)
}
