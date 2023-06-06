package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"github.com/arvians-id/go-portfolio/util"
	"log"
)

type UserServiceContract interface {
	FindAll(ctx context.Context) ([]*entity.User, error)
	FindByID(ctx context.Context, id int64) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) (*entity.User, error)
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

func (service *UserService) FindAll(ctx context.Context) ([]*entity.User, error) {
	users, err := service.UserRepository.FindAll(ctx)
	if err != nil {
		log.Println("[UserService][FindAll] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return users, nil
}

func (service *UserService) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	user, err := service.UserRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[UserService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return user, nil
}

func (service *UserService) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	passwordHashed, err := util.HashPassword(user.Password)
	if err != nil {
		log.Println("[UserService][HashPassword] problem hashing password, err: ", err.Error())
		return nil, err
	}

	user.Password = passwordHashed
	userCreated, err := service.UserRepository.Create(ctx, user)
	if err != nil {
		log.Println("[UserService][Create] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return userCreated, nil
}

func (service *UserService) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	_, err := service.UserRepository.FindByID(ctx, user.ID)
	if err != nil {
		log.Println("[UserService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	if user.Password != "" {
		hashedPassword, err := util.HashPassword(user.Password)
		if err != nil {
			log.Println("[UserService][HashPassword] problem hashing password, err: ", err.Error())
			return nil, err
		}

		user.Password = hashedPassword
	}

	err = service.UserRepository.Update(ctx, user)
	if err != nil {
		log.Println("[UserService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	userUpdated, err := service.UserRepository.FindByID(ctx, user.ID)
	if err != nil {
		log.Println("[UserService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return userUpdated, nil
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
