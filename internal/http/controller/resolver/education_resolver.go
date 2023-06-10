package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/util"
)

func (q queryResolver) FindAllEducation(ctx context.Context) ([]*model.Education, error) {
	educations, err := q.EducationService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var results []*model.Education
	for _, education := range educations {
		results = append(results, &model.Education{
			ID:           education.ID,
			Institution:  education.Institution,
			Degree:       education.Degree,
			FieldOfStudy: education.FieldOfStudy,
			Grade:        education.Grade,
			Description:  education.Description,
			StartDate:    education.StartDate,
			EndDate:      education.EndDate,
		})
	}

	return results, nil
}

func (q queryResolver) FindByIDEducation(ctx context.Context, id int64) (*model.Education, error) {
	education, err := q.EducationService.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.Education{
		ID:           education.ID,
		Institution:  education.Institution,
		Degree:       education.Degree,
		FieldOfStudy: education.FieldOfStudy,
		Grade:        education.Grade,
		Description:  education.Description,
		StartDate:    education.StartDate,
		EndDate:      education.EndDate,
	}, nil
}

func (m mutationResolver) CreateEducation(ctx context.Context, input model.CreateEducationRequest) (*model.Education, error) {
	err := util.ValidateStruct(ctx, input)
	if err != nil {
		return nil, err
	}

	education, err := m.EducationService.Create(ctx, &entity.Education{
		Institution:  input.Institution,
		Degree:       input.Degree,
		FieldOfStudy: input.FieldOfStudy,
		Grade:        input.Grade,
		Description:  input.Description,
		StartDate:    input.StartDate,
		EndDate:      input.EndDate,
	})
	if err != nil {
		return nil, err
	}

	return &model.Education{
		ID:           education.ID,
		Institution:  education.Institution,
		Degree:       education.Degree,
		FieldOfStudy: education.FieldOfStudy,
		Grade:        education.Grade,
		Description:  education.Description,
		StartDate:    education.StartDate,
		EndDate:      education.EndDate,
	}, nil
}

func (m mutationResolver) UpdateEducation(ctx context.Context, input model.UpdateEducationRequest) (*model.Education, error) {
	err := util.ValidateStruct(ctx, input)
	if err != nil {
		return nil, err
	}

	education, err := m.EducationService.Update(ctx, &entity.Education{
		ID:           input.ID,
		Institution:  *input.Institution,
		Degree:       *input.Degree,
		FieldOfStudy: *input.FieldOfStudy,
		Grade:        *input.Grade,
		Description:  input.Description,
		StartDate:    *input.StartDate,
		EndDate:      input.EndDate,
	})
	if err != nil {
		return nil, err
	}

	return &model.Education{
		ID:           education.ID,
		Institution:  education.Institution,
		Degree:       education.Degree,
		FieldOfStudy: education.FieldOfStudy,
		Grade:        education.Grade,
		Description:  education.Description,
		StartDate:    education.StartDate,
		EndDate:      education.EndDate,
	}, nil
}

func (m mutationResolver) DeleteEducation(ctx context.Context, id int64) (bool, error) {
	err := m.EducationService.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
