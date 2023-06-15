package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/util"
)

func (q queryResolver) FindAllCertificate(ctx context.Context) ([]*model.Certificate, error) {
	certificates, err := q.CertificateService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var results []*model.Certificate
	for _, certificate := range certificates {
		results = append(results, &model.Certificate{
			ID:             certificate.ID,
			Name:           certificate.Name,
			Organization:   certificate.Organization,
			IssueDate:      certificate.IssueDate,
			ExpirationDate: certificate.ExpirationDate,
			CredentialID:   certificate.CredentialID,
			Image:          certificate.Image,
		})
	}

	return results, nil
}

func (q queryResolver) FindByIDCertificate(ctx context.Context, id int64) (*model.Certificate, error) {
	certificate, err := q.CertificateService.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.Certificate{
		ID:             certificate.ID,
		Name:           certificate.Name,
		Organization:   certificate.Organization,
		IssueDate:      certificate.IssueDate,
		ExpirationDate: certificate.ExpirationDate,
		CredentialID:   certificate.CredentialID,
		Image:          certificate.Image,
	}, nil
}

func (m mutationResolver) CreateCertificate(ctx context.Context, input model.CreateCertificateRequest) (*model.Certificate, error) {
	err := util.ValidateStruct(ctx, input)
	if err != nil {
		return nil, err
	}

	var fileName string
	if input.Image != nil {
		path := "images/certificate"
		fileName, err = util.UploadFile(path, *input.Image)
		if err != nil {
			return nil, err
		}
	}

	certificate, err := m.CertificateService.Create(ctx, &entity.Certificate{
		Name:           input.Name,
		Organization:   input.Organization,
		IssueDate:      input.IssueDate,
		ExpirationDate: input.ExpirationDate,
		CredentialID:   input.CredentialID,
		Image:          &fileName,
	})
	if err != nil {
		return nil, err
	}

	return &model.Certificate{
		ID:             certificate.ID,
		Name:           certificate.Name,
		Organization:   certificate.Organization,
		IssueDate:      certificate.IssueDate,
		ExpirationDate: certificate.ExpirationDate,
		CredentialID:   certificate.CredentialID,
		Image:          certificate.Image,
	}, nil
}

func (m mutationResolver) UpdateCertificate(ctx context.Context, input model.UpdateCertificateRequest) (*model.Certificate, error) {
	err := util.ValidateStruct(ctx, input)
	if err != nil {
		return nil, err
	}

	var fileName string
	if input.Image != nil {
		path := "images/certificate"
		fileName, err = util.UploadFile(path, *input.Image)
		if err != nil {
			return nil, err
		}
	}

	certificate, err := m.CertificateService.Update(ctx, &entity.Certificate{
		ID:             input.ID,
		Name:           *input.Name,
		Organization:   *input.Organization,
		IssueDate:      *input.IssueDate,
		ExpirationDate: input.ExpirationDate,
		CredentialID:   input.CredentialID,
		Image:          &fileName,
	})
	if err != nil {
		return nil, err
	}

	return &model.Certificate{
		ID:             certificate.ID,
		Name:           certificate.Name,
		Organization:   certificate.Organization,
		IssueDate:      certificate.IssueDate,
		ExpirationDate: certificate.ExpirationDate,
		CredentialID:   certificate.CredentialID,
		Image:          certificate.Image,
	}, nil
}

func (m mutationResolver) DeleteCertificate(ctx context.Context, id int64) (bool, error) {
	err := m.CertificateService.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
