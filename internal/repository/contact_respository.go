package repository

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"gorm.io/gorm"
	"log"
)

type ContactRepositoryContract interface {
	FindAll(ctx context.Context) ([]*entity.Contact, error)
	FindByID(ctx context.Context, id int64) (*entity.Contact, error)
	Create(ctx context.Context, request *model.CreateContactRequest) (*entity.Contact, error)
	Update(ctx context.Context, request *model.UpdateContactRequest) error
	Delete(ctx context.Context, id int64) error
}

type ContactRepository struct {
	DB *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepositoryContract {
	return &ContactRepository{DB: db}
}

func (repository *ContactRepository) FindAll(ctx context.Context) ([]*entity.Contact, error) {
	query := "SELECT * FROM contacts"
	var contacts []*entity.Contact
	err := repository.DB.WithContext(ctx).Raw(query).Scan(&contacts).Error
	if err != nil {
		log.Println("[ContactRepository][FindAll] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return contacts, nil
}

func (repository *ContactRepository) FindByID(ctx context.Context, id int64) (*entity.Contact, error) {
	query := "SELECT * FROM contacts WHERE id = ?"
	var contact entity.Contact
	row := repository.DB.WithContext(ctx).Raw(query, id).Row()
	err := row.Scan(
		&contact.ID,
		&contact.Platform,
		&contact.URL,
		&contact.Icon,
	)
	if err != nil {
		log.Println("[ContactRepository][FindByID] problem with scanning db row, err: ", err.Error())
		return nil, err
	}

	return &contact, nil
}

func (repository *ContactRepository) Create(ctx context.Context, request *model.CreateContactRequest) (*entity.Contact, error) {
	var contact entity.Contact
	contact.Platform = request.Platform
	contact.URL = request.URL
	contact.Icon = request.Icon

	err := repository.DB.WithContext(ctx).Create(&contact).Error
	if err != nil {
		log.Println("[ContactRepository][Create] problem with scanning db row, err: ", err.Error())
		return nil, err
	}

	return &contact, nil
}

func (repository *ContactRepository) Update(ctx context.Context, request *model.UpdateContactRequest) error {
	var contact entity.Contact
	contact.ID = request.ID
	contact.Platform = *request.Platform
	contact.URL = *request.URL
	contact.Icon = request.Icon

	err := repository.DB.WithContext(ctx).Updates(&contact).Error
	if err != nil {
		log.Println("[ContactRepository][Update] problem querying to db, err: ", err.Error())
		return err
	}

	return nil
}

func (repository *ContactRepository) Delete(ctx context.Context, id int64) error {
	err := repository.DB.WithContext(ctx).Delete(&entity.Contact{}, id).Error
	if err != nil {
		log.Println("[ContactRepository][Delete] problem querying to db, err: ", err.Error())
		return err
	}

	return nil
}
