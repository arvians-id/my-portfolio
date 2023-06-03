package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
)

func (q queryResolver) CertificateFindAll(ctx context.Context) ([]*model.Certificate, error) {
	//TODO implement me
	panic("implement me")
}

func (q queryResolver) CertificateFindByID(ctx context.Context, id int64) (*model.Certificate, error) {
	//TODO implement me
	panic("implement me")
}
