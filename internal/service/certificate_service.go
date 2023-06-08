package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"github.com/arvians-id/go-portfolio/util"
	"log"
)

type CertificateServiceContract interface {
	FindAll(ctx context.Context) ([]*entity.Certificate, error)
	FindByID(ctx context.Context, id int64) (*entity.Certificate, error)
	Create(ctx context.Context, certificate *entity.Certificate) (*entity.Certificate, error)
	Update(ctx context.Context, certificate *entity.Certificate) (*entity.Certificate, error)
	Delete(ctx context.Context, id int64) error
}

type CertificateService struct {
	CertificateRepository repository.CertificateRepositoryContract
}

func NewCertificateService(CertificateRepository repository.CertificateRepositoryContract) CertificateServiceContract {
	return &CertificateService{
		CertificateRepository: CertificateRepository,
	}
}

func (service *CertificateService) FindAll(ctx context.Context) ([]*entity.Certificate, error) {
	certificates, err := service.CertificateRepository.FindAll(ctx)
	if err != nil {
		log.Println("[CertificateService][FindAll] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return certificates, nil
}

func (service *CertificateService) FindByID(ctx context.Context, id int64) (*entity.Certificate, error) {
	certificate, err := service.CertificateRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[CertificateService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return certificate, nil
}

func (service *CertificateService) Create(ctx context.Context, certificate *entity.Certificate) (*entity.Certificate, error) {
	certificateCreated, err := service.CertificateRepository.Create(ctx, certificate)
	if err != nil {
		log.Println("[CertificateService][Create] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return certificateCreated, nil
}

func (service *CertificateService) Update(ctx context.Context, certificate *entity.Certificate) (*entity.Certificate, error) {
	certificateCheck, err := service.CertificateRepository.FindByID(ctx, certificate.ID)
	if err != nil {
		log.Println("[CertificateService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	err = service.CertificateRepository.Update(ctx, certificate)
	if err != nil {
		log.Println("[CertificateService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	if certificateCheck.Image != certificate.Image {
		path := "images/certificate"
		err = util.DeleteFile(path, *certificateCheck.Image)
		if err != nil {
			return nil, err
		}
	}

	certificateUpdated, err := service.CertificateRepository.FindByID(ctx, certificate.ID)
	if err != nil {
		log.Println("[CertificateService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return certificateUpdated, nil
}

func (service *CertificateService) Delete(ctx context.Context, id int64) error {
	certificateCheck, err := service.CertificateRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[CertificateService][Delete] problem calling repository, err: ", err.Error())
		return err
	}

	err = service.CertificateRepository.Delete(ctx, certificateCheck.ID)
	if err != nil {
		log.Println("[CertificateService][Delete] problem calling repository, err: ", err.Error())
		return err
	}

	if certificateCheck.Image != nil {
		path := "images/certificate"
		err = util.DeleteFile(path, *certificateCheck.Image)
		if err != nil {
			return err
		}
	}

	return nil
}
