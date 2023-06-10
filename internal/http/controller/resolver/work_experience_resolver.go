package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/util"
)

func (q queryResolver) FindAllWorkExperience(ctx context.Context) ([]*model.WorkExperience, error) {
	workExperiences, err := q.WorkExperienceService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var results []*model.WorkExperience
	for _, workExperience := range workExperiences {
		results = append(results, &model.WorkExperience{
			ID:          workExperience.ID,
			Role:        workExperience.Role,
			Company:     workExperience.Company,
			Description: workExperience.Description,
			StartDate:   workExperience.StartDate,
			EndDate:     workExperience.EndDate,
			JobType:     workExperience.JobType,
			CreatedAt:   workExperience.CreatedAt.String(),
			UpdatedAt:   workExperience.UpdatedAt.String(),
		})
	}

	return results, nil
}

func (q queryResolver) FindByIDWorkExperience(ctx context.Context, id int64) (*model.WorkExperience, error) {
	workExperience, err := q.WorkExperienceService.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.WorkExperience{
		ID:          workExperience.ID,
		Role:        workExperience.Role,
		Company:     workExperience.Company,
		Description: workExperience.Description,
		StartDate:   workExperience.StartDate,
		EndDate:     workExperience.EndDate,
		JobType:     workExperience.JobType,
		CreatedAt:   workExperience.CreatedAt.String(),
		UpdatedAt:   workExperience.UpdatedAt.String(),
	}, nil
}

func (m mutationResolver) CreateWorkExperience(ctx context.Context, input model.CreateWorkExperienceRequest) (*model.WorkExperience, error) {
	err := util.ValidateStruct(ctx, input)
	if err != nil {
		return nil, err
	}

	var skills []*entity.Skill
	for _, id := range input.Skills {
		skills = append(skills, &entity.Skill{
			ID: id,
		})
	}

	workExperience, err := m.WorkExperienceService.Create(ctx, &entity.WorkExperience{
		Role:        input.Role,
		Company:     input.Company,
		Description: input.Description,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		JobType:     input.JobType,
		Skills:      skills,
	})
	if err != nil {
		return nil, err
	}

	return &model.WorkExperience{
		ID:          workExperience.ID,
		Role:        workExperience.Role,
		Company:     workExperience.Company,
		Description: workExperience.Description,
		StartDate:   workExperience.StartDate,
		EndDate:     workExperience.EndDate,
		JobType:     workExperience.JobType,
		CreatedAt:   workExperience.CreatedAt.String(),
		UpdatedAt:   workExperience.UpdatedAt.String(),
	}, nil
}

func (m mutationResolver) UpdateWorkExperience(ctx context.Context, input model.UpdateWorkExperienceRequest) (*model.WorkExperience, error) {
	err := util.ValidateStruct(ctx, input)
	if err != nil {
		return nil, err
	}

	var skills []*entity.Skill
	for _, id := range input.Skills {
		skills = append(skills, &entity.Skill{
			ID: id,
		})
	}

	workExperience, err := m.WorkExperienceService.Update(ctx, &entity.WorkExperience{
		ID:          input.ID,
		Role:        *input.Role,
		Company:     *input.Company,
		Description: input.Description,
		StartDate:   *input.StartDate,
		EndDate:     input.EndDate,
		JobType:     *input.JobType,
		Skills:      skills,
	})
	if err != nil {
		return nil, err
	}

	return &model.WorkExperience{
		ID:          workExperience.ID,
		Role:        workExperience.Role,
		Company:     workExperience.Company,
		Description: workExperience.Description,
		StartDate:   workExperience.StartDate,
		EndDate:     workExperience.EndDate,
		JobType:     workExperience.JobType,
		CreatedAt:   workExperience.CreatedAt.String(),
		UpdatedAt:   workExperience.UpdatedAt.String(),
	}, nil
}

func (m mutationResolver) DeleteWorkExperience(ctx context.Context, id int64) (bool, error) {
	err := m.WorkExperienceService.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (w workExperienceResolver) Skills(ctx context.Context, obj *model.WorkExperience) ([]*model.Skill, error) {
	skills, err := GetLoaders(ctx).ListSkillsByWorkExperienceIDs.Load(obj.ID)
	if err != nil {
		return nil, err
	}

	return skills, nil
}
