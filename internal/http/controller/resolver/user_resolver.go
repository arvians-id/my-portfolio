package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/util"
)

func (q queryResolver) FindAllUser(ctx context.Context) ([]*model.User, error) {
	users, err := q.UserService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var results []*model.User
	for _, user := range users {
		results = append(results, &model.User{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Bio:       user.Bio,
			Pronouns:  user.Pronouns,
			Country:   user.Country,
			JobTitle:  user.JobTitle,
			Image:     user.Image,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		})
	}

	return results, nil
}

func (q queryResolver) FindByIDUser(ctx context.Context, id int64) (*model.User, error) {
	user, err := q.UserService.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Bio:       user.Bio,
		Pronouns:  user.Pronouns,
		Country:   user.Country,
		JobTitle:  user.JobTitle,
		Image:     user.Image,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (m mutationResolver) CreateUser(ctx context.Context, input model.CreateUserRequest) (*model.User, error) {
	err := util.ValidateStruct(ctx, input)
	if err != nil {
		return nil, err
	}

	user, err := m.UserService.Create(ctx, &entity.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Pronouns: input.Pronouns,
		Country:  input.Country,
		JobTitle: input.JobTitle,
	})
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Bio:       user.Bio,
		Pronouns:  user.Pronouns,
		Country:   user.Country,
		JobTitle:  user.JobTitle,
		Image:     user.Image,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (m mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserRequest) (*model.User, error) {
	err := util.ValidateStruct(ctx, input)
	if err != nil {
		return nil, err
	}

	var fileName string
	path := "images/user"
	if input.Image != nil {
		fileName, err = util.UploadFile(path, *input.Image)
		if err != nil {
			return nil, err
		}
	}

	user, err := m.UserService.Update(ctx, &entity.User{
		ID:       input.ID,
		Name:     *input.Name,
		Password: *input.Password,
		Bio:      input.Bio,
		Pronouns: *input.Pronouns,
		Country:  *input.Country,
		JobTitle: *input.JobTitle,
		Image:    &fileName,
	})
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Bio:       user.Bio,
		Pronouns:  user.Pronouns,
		Country:   user.Country,
		JobTitle:  user.JobTitle,
		Image:     user.Image,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (m mutationResolver) DeleteUser(ctx context.Context, id int64) (bool, error) {
	err := m.UserService.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
