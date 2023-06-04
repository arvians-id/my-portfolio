package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"github.com/arvians-id/go-portfolio/util"
	"log"
)

type UserServiceContract interface {
	FindAll(ctx context.Context) ([]*model.User, error)
	FindByID(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, request *model.CreateUserRequest) (*model.User, error)
	Update(ctx context.Context, request *model.UpdateUserRequest) (*model.User, error)
	Delete(ctx context.Context, id int64) error
}

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserServiceContract {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (service *UserService) FindAll(ctx context.Context) ([]*model.User, error) {
	users, err := service.UserRepository.FindAll(ctx)
	if err != nil {
		log.Println("[UserService][FindAll] problem calling repository, err: ", err.Error())
		return nil, err
	}

	var usersModel []*model.User
	for _, user := range users {
		usersModel = append(usersModel, &model.User{
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
		})
	}

	return usersModel, nil
}

func (service *UserService) FindByID(ctx context.Context, id int64) (*model.User, error) {
	user, err := service.UserRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[UserService][FindByID] problem calling repository, err: ", err.Error())
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

func (service *UserService) Create(ctx context.Context, request *model.CreateUserRequest) (*model.User, error) {
	passwordHashed, err := util.HashPassword(request.Password)
	if err != nil {
		log.Println("[UserService][HashPassword] problem hashing password, err: ", err.Error())
		return nil, err
	}

	var user entity.User
	user.Name = request.Name
	user.Email = request.Email
	user.Password = passwordHashed
	user.Pronouns = request.Pronouns
	user.Country = request.Country
	user.JobTitle = request.JobTitle

	userCreated, err := service.UserRepository.Create(ctx, &user)
	if err != nil {
		log.Println("[UserService][Create] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return &model.User{
		ID:        userCreated.ID,
		Name:      userCreated.Name,
		Email:     userCreated.Email,
		Bio:       userCreated.Bio,
		Pronouns:  userCreated.Pronouns,
		Country:   userCreated.Country,
		JobTitle:  userCreated.JobTitle,
		Image:     userCreated.Image,
		CreatedAt: userCreated.CreatedAt.String(),
		UpdatedAt: userCreated.UpdatedAt.String(),
	}, nil
}

func (service *UserService) Update(ctx context.Context, request *model.UpdateUserRequest) (*model.User, error) {
	userCheck, err := service.UserRepository.FindByID(ctx, request.ID)
	if err != nil {
		log.Println("[UserService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	if request.Password != nil {
		hashedPassword, err := util.HashPassword(*request.Password)
		if err != nil {
			log.Println("[UserService][HashPassword] problem hashing password, err: ", err.Error())
			return nil, err
		}

		userCheck.Password = hashedPassword
	}

	userCheck.Name = *request.Name
	userCheck.Bio = request.Bio
	userCheck.Pronouns = *request.Pronouns
	userCheck.Country = *request.Country
	userCheck.JobTitle = *request.JobTitle
	userCheck.Image = request.Image

	err = service.UserRepository.Update(ctx, &entity.User{
		ID:       userCheck.ID,
		Name:     userCheck.Name,
		Bio:      userCheck.Bio,
		Pronouns: userCheck.Pronouns,
		Country:  userCheck.Country,
		JobTitle: userCheck.JobTitle,
		Image:    userCheck.Image,
	})
	if err != nil {
		log.Println("[UserService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return &model.User{
		ID:        userCheck.ID,
		Name:      userCheck.Name,
		Email:     userCheck.Email,
		Bio:       userCheck.Bio,
		Pronouns:  userCheck.Pronouns,
		Country:   userCheck.Country,
		JobTitle:  userCheck.JobTitle,
		Image:     userCheck.Image,
		CreatedAt: userCheck.CreatedAt.String(),
		UpdatedAt: userCheck.UpdatedAt.String(),
	}, nil
}

func (service *UserService) Delete(ctx context.Context, id int64) error {
	_, err := service.UserRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[UserService][FindByID] problem calling repository, err: ", err.Error())
		return err
	}

	err = service.UserRepository.Delete(ctx, id)
	if err != nil {
		log.Println("[UserService][Delete] problem calling repository, err: ", err.Error())
		return err
	}

	return nil
}
