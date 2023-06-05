package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"log"
)

type CertificateServiceContract interface {
	FindAll(ctx context.Context) ([]*model.Certificate, error)
	FindByID(ctx context.Context, id int64) (*model.Certificate, error)
	Create(ctx context.Context, request *model.CreateCertificateRequest) (*model.Certificate, error)
	Update(ctx context.Context, request *model.UpdateCertificateRequest) (*model.Certificate, error)
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

func (service *CertificateService) FindAll(ctx context.Context) ([]*model.Certificate, error) {
	certificates, err := service.CertificateRepository.FindAll(ctx)
	if err != nil {
		log.Println("[CertificateService][FindAll] problem calling repository, err: ", err.Error())
		return nil, err
	}

	var results []*model.Certificate
	for _, certificate := range certificates {
		results = append(results, certificate.ToModel())
	}

	return results, nil
}

func (service *CertificateService) FindByID(ctx context.Context, id int64) (*model.Certificate, error) {
	certificate, err := service.CertificateRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[CertificateService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return certificate.ToModel(), nil
}

func (service *CertificateService) Create(ctx context.Context, request *model.CreateCertificateRequest) (*model.Certificate, error) {
	certificateCreated, err := service.CertificateRepository.Create(ctx, request)
	if err != nil {
		log.Println("[CertificateService][Create] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return certificateCreated.ToModel(), nil
}

func (service *CertificateService) Update(ctx context.Context, request *model.UpdateCertificateRequest) (*model.Certificate, error) {
	_, err := service.CertificateRepository.FindByID(ctx, request.ID)
	if err != nil {
		log.Println("[CertificateService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	err = service.CertificateRepository.Update(ctx, request)
	if err != nil {
		log.Println("[CertificateService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	certificateUpdated, err := service.CertificateRepository.FindByID(ctx, request.ID)
	if err != nil {
		log.Println("[CertificateService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return certificateUpdated.ToModel(), nil
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

	return nil
}
