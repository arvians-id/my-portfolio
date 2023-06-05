package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
)

func (q queryResolver) FindAllCertificate(ctx context.Context) ([]*model.Certificate, error) {
	certificates, err := q.CertificateService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return certificates, nil
}

func (q queryResolver) FindByIDCertificate(ctx context.Context, id int64) (*model.Certificate, error) {
	certificate, err := q.CertificateService.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return certificate, nil
}

func (m mutationResolver) CreateCertificate(ctx context.Context, input model.CreateCertificateRequest) (*model.Certificate, error) {
	certificate, err := m.CertificateService.Create(ctx, &input)
	if err != nil {
		return nil, err
	}

	return certificate, nil
}

func (m mutationResolver) UpdateCertificate(ctx context.Context, input model.UpdateCertificateRequest) (*model.Certificate, error) {
	certificate, err := m.CertificateService.Update(ctx, &input)
	if err != nil {
		return nil, err
	}

	return certificate, nil
}

func (m mutationResolver) DeleteCertificate(ctx context.Context, id int64) (bool, error) {
	err := m.CertificateService.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
