package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
)

func (q queryResolver) UserFindAll(ctx context.Context) ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (q queryResolver) UserFindByID(ctx context.Context, id int64) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m mutationResolver) CreateUser(ctx context.Context, input model.CreateUserRequest) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}
