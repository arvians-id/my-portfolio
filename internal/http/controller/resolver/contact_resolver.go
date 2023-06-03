package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
)

func (q queryResolver) ContactFindAll(ctx context.Context) ([]*model.Contact, error) {
	//TODO implement me
	panic("implement me")
}

func (q queryResolver) ContactFindByID(ctx context.Context, id int64) (*model.Contact, error) {
	//TODO implement me
	panic("implement me")
}
