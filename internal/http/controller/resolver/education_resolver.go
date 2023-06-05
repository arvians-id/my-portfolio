package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
)

func (q queryResolver) FindAllEducation(ctx context.Context) ([]*model.Education, error) {
	educations, err := q.EducationService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return educations, nil
}

func (q queryResolver) FindByIDEducation(ctx context.Context, id int64) (*model.Education, error) {
	education, err := q.EducationService.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return education, nil
}

func (m mutationResolver) CreateEducation(ctx context.Context, input model.CreateEducationRequest) (*model.Education, error) {
	education, err := m.EducationService.Create(ctx, &input)
	if err != nil {
		return nil, err
	}

	return education, nil
}

func (m mutationResolver) UpdateEducation(ctx context.Context, input model.UpdateEducationRequest) (*model.Education, error) {
	education, err := m.EducationService.Update(ctx, &input)
	if err != nil {
		return nil, err
	}

	return education, nil
}

func (m mutationResolver) DeleteEducation(ctx context.Context, id int64) (bool, error) {
	err := m.EducationService.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
