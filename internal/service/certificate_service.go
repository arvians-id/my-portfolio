package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"log"
)

type CertificateServiceContract interface {
	FindAll(ctx context.Context) ([]*model.Certificate, error)
	FindById(ctx context.Context, id int64) (*model.Certificate, error)
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

	var certificatesModel []*model.Certificate
	for _, certificate := range certificates {
		certificatesModel = append(certificatesModel, (*model.Certificate)(certificate))
	}

	return certificatesModel, nil
}

func (service *CertificateService) FindById(ctx context.Context, id int64) (*model.Certificate, error) {
	certificate, err := service.CertificateRepository.FindById(ctx, id)
	if err != nil {
		log.Println("[CertificateService][FindById] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return (*model.Certificate)(certificate), nil
}

func (service *CertificateService) Create(ctx context.Context, request *model.CreateCertificateRequest) (*model.Certificate, error) {
	var certificate entity.Certificate
	certificate.Name = request.Name
	certificate.Organization = request.Organization
	certificate.IssueDate = request.IssueDate
	certificate.ExpirationDate = request.ExpirationDate
	certificate.CredentialID = request.CredentialID
	certificate.ImageURL = request.ImageURL

	certificateCreated, err := service.CertificateRepository.Create(ctx, &certificate)
	if err != nil {
		log.Println("[CertificateService][Create] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return (*model.Certificate)(certificateCreated), nil
}

func (service *CertificateService) Update(ctx context.Context, request *model.UpdateCertificateRequest) (*model.Certificate, error) {
	certificateCheck, err := service.CertificateRepository.FindById(ctx, request.ID)
	if err != nil {
		log.Println("[CertificateService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	certificateCheck.Name = *request.Name
	certificateCheck.Organization = *request.Organization
	certificateCheck.IssueDate = *request.IssueDate
	certificateCheck.ExpirationDate = request.ExpirationDate
	certificateCheck.CredentialID = request.CredentialID
	certificateCheck.ImageURL = request.ImageURL

	err = service.CertificateRepository.Update(ctx, certificateCheck)
	if err != nil {
		log.Println("[CertificateService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return (*model.Certificate)(certificateCheck), nil
}

func (service *CertificateService) Delete(ctx context.Context, id int64) error {
	certificateCheck, err := service.CertificateRepository.FindById(ctx, id)
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
