package service

import (
	"context"
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

	var results []*model.User
	for _, user := range users {
		results = append(results, user.ToModel())
	}

	return results, nil
}

func (service *UserService) FindByID(ctx context.Context, id int64) (*model.User, error) {
	user, err := service.UserRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[UserService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return user.ToModel(), nil
}

func (service *UserService) Create(ctx context.Context, request *model.CreateUserRequest) (*model.User, error) {
	passwordHashed, err := util.HashPassword(request.Password)
	if err != nil {
		log.Println("[UserService][HashPassword] problem hashing password, err: ", err.Error())
		return nil, err
	}

	request.Password = passwordHashed
	userCreated, err := service.UserRepository.Create(ctx, request)
	if err != nil {
		log.Println("[UserService][Create] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return userCreated.ToModel(), nil
}

func (service *UserService) Update(ctx context.Context, request *model.UpdateUserRequest) (*model.User, error) {
	_, err := service.UserRepository.FindByID(ctx, request.ID)
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

		request.Password = &hashedPassword
	}

	err = service.UserRepository.Update(ctx, request)
	if err != nil {
		log.Println("[UserService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	userUpdated, err := service.UserRepository.FindByID(ctx, request.ID)
	if err != nil {
		log.Println("[UserService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return userUpdated.ToModel(), nil
}

func (service *UserService) Delete(ctx context.Context, id int64) error {
	workExperienceCheck, err := service.UserRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[UserService][FindByID] problem calling repository, err: ", err.Error())
		return err
	}

	err = service.UserRepository.Delete(ctx, workExperienceCheck.ID)
	if err != nil {
		log.Println("[UserService][Delete] problem calling repository, err: ", err.Error())
		return err
	}

	return nil
}
