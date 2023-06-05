package repository

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"gorm.io/gorm"
	"log"
)

type CertificateRepositoryContract interface {
	FindAll(ctx context.Context) ([]*entity.Certificate, error)
	FindByID(ctx context.Context, id int64) (*entity.Certificate, error)
	Create(ctx context.Context, request *model.CreateCertificateRequest) (*entity.Certificate, error)
	Update(ctx context.Context, request *model.UpdateCertificateRequest) error
	Delete(ctx context.Context, id int64) error
}

type CertificateRepository struct {
	DB *gorm.DB
}

func NewCertificateRepository(db *gorm.DB) CertificateRepositoryContract {
	return &CertificateRepository{DB: db}
}

func (repository *CertificateRepository) FindAll(ctx context.Context) ([]*entity.Certificate, error) {
	query := "SELECT * FROM certificates"
	var certificates []*entity.Certificate
	err := repository.DB.WithContext(ctx).Raw(query).Scan(&certificates).Error
	if err != nil {
		log.Println("[CertificateRepository][FindAll] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return certificates, nil
}

func (repository *CertificateRepository) FindByID(ctx context.Context, id int64) (*entity.Certificate, error) {
	query := "SELECT * FROM certificates WHERE id = ?"
	var certificate entity.Certificate
	row := repository.DB.WithContext(ctx).Raw(query, id).Row()
	err := row.Scan(
		&certificate.ID,
		&certificate.Name,
		&certificate.Organization,
		&certificate.IssueDate,
		&certificate.ExpirationDate,
		&certificate.CredentialID,
		&certificate.ImageURL,
	)
	if err != nil {
		log.Println("[CertificateRepository][FindByID] problem with scanning db row, err: ", err.Error())
		return nil, err
	}

	return &certificate, nil
}

func (repository *CertificateRepository) Create(ctx context.Context, request *model.CreateCertificateRequest) (*entity.Certificate, error) {
	var certificate entity.Certificate
	certificate.Name = request.Name
	certificate.Organization = request.Organization
	certificate.IssueDate = request.IssueDate
	certificate.ExpirationDate = request.ExpirationDate
	certificate.CredentialID = request.CredentialID
	certificate.ImageURL = request.ImageURL

	err := repository.DB.WithContext(ctx).Create(&certificate).Error
	if err != nil {
		log.Println("[CertificateRepository][Create] problem with scanning db row, err: ", err.Error())
		return nil, err
	}

	return &certificate, nil
}

func (repository *CertificateRepository) Update(ctx context.Context, request *model.UpdateCertificateRequest) error {
	var certificate entity.Certificate
	certificate.ID = request.ID
	certificate.Name = *request.Name
	certificate.Organization = *request.Organization
	certificate.IssueDate = *request.IssueDate
	certificate.ExpirationDate = request.ExpirationDate
	certificate.CredentialID = request.CredentialID
	certificate.ImageURL = request.ImageURL

	err := repository.DB.WithContext(ctx).Updates(&certificate).Error
	if err != nil {
		log.Println("[CertificateRepository][Update] problem querying to db, err: ", err.Error())
		return err
	}

	return nil
}

func (repository *CertificateRepository) Delete(ctx context.Context, id int64) error {
	err := repository.DB.WithContext(ctx).Delete(&entity.Certificate{}, id).Error
	if err != nil {
		log.Println("[CertificateRepository][Delete] problem querying to db, err: ", err.Error())
		return err
	}

	return nil
}
