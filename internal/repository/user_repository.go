package repository

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"gorm.io/gorm"
	"log"
)

type UserRepositoryContract interface {
	FindAll(ctx context.Context) ([]*entity.User, error)
	FindByID(ctx context.Context, id int64) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int64) error
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		DB: db,
	}
}

func (repository *UserRepository) FindAll(ctx context.Context) ([]*entity.User, error) {
	query := "SELECT * FROM users ORDER BY created_at DESC"
	var users []*entity.User
	err := repository.DB.WithContext(ctx).Raw(query).Scan(&users).Error
	if err != nil {
		log.Println("[UserRepository][FindAll] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return users, nil
}

func (repository *UserRepository) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	var user entity.User
	row := repository.DB.WithContext(ctx).Raw(query, id).Row()
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Bio,
		&user.Pronouns,
		&user.Country,
		&user.JobTitle,
		&user.Image,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		log.Println("[UserRepository][FindByID] problem with scanning db row, err: ", err.Error())
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := "SELECT * FROM users WHERE email = ?"
	var user entity.User
	row := repository.DB.WithContext(ctx).Raw(query, email).Row()
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Bio,
		&user.Pronouns,
		&user.Country,
		&user.JobTitle,
		&user.Image,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		log.Println("[UserRepository][FindByEmail] problem with scanning db row, err: ", err.Error())
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := repository.DB.WithContext(ctx).Select("name", "email", "password", "pronouns", "country", "job_title").Create(&user).Error
	if err != nil {
		log.Println("[UserRepository][Create] problem with scanning db row, err: ", err.Error())
		return nil, err
	}

	return user, nil
}

func (repository *UserRepository) Update(ctx context.Context, user *entity.User) error {
	err := repository.DB.WithContext(ctx).Updates(&user).Error
	if err != nil {
		log.Println("[UserRepository][Update] problem querying to db, err: ", err.Error())
		return err
	}

	return nil
}

func (repository *UserRepository) Delete(ctx context.Context, id int64) error {
	err := repository.DB.WithContext(ctx).Delete(&entity.User{}, id).Error
	if err != nil {
		log.Println("[UserRepository][Delete] problem querying to db, err: ", err.Error())
		return err
	}

	return nil
}
