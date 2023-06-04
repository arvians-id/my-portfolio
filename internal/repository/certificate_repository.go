package repository

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"gorm.io/gorm"
	"log"
)

type CertificateRepositoryContract interface {
	FindAll(ctx context.Context) ([]*entity.Certificate, error)
	FindById(ctx context.Context, id int64) (*entity.Certificate, error)
	Create(ctx context.Context, certificate *entity.Certificate) (*entity.Certificate, error)
	Update(ctx context.Context, certificate *entity.Certificate) error
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

func (repository *CertificateRepository) FindById(ctx context.Context, id int64) (*entity.Certificate, error) {
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

func (repository *CertificateRepository) Create(ctx context.Context, certificate *entity.Certificate) (*entity.Certificate, error) {
	err := repository.DB.WithContext(ctx).Create(&certificate).Error
	if err != nil {
		log.Println("[CertificateRepository][Create] problem with scanning db row, err: ", err.Error())
		return nil, err
	}

	return certificate, nil
}

func (repository *CertificateRepository) Update(ctx context.Context, certificate *entity.Certificate) error {
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
