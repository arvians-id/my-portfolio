package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
)

func (q queryResolver) FindAllUser(ctx context.Context) ([]*model.User, error) {
	users, err := q.UserService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (q queryResolver) FindByIDUser(ctx context.Context, id int64) (*model.User, error) {
	user, err := q.UserService.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m mutationResolver) CreateUser(ctx context.Context, input model.CreateUserRequest) (*model.User, error) {
	user, err := m.UserService.Create(ctx, &input)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserRequest) (*model.User, error) {
	user, err := m.UserService.Update(ctx, &input)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m mutationResolver) DeleteUser(ctx context.Context, id int64) (bool, error) {
	err := m.UserService.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
