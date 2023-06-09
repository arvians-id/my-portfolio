package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/util"
)

func (q queryResolver) FindAllCategorySkill(ctx context.Context) ([]*model.CategorySkill, error) {
	categorySkills, err := q.CategorySkillService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var results []*model.CategorySkill
	for _, categorySkill := range categorySkills {
		results = append(results, &model.CategorySkill{
			ID:        categorySkill.ID,
			Name:      categorySkill.Name,
			CreatedAt: categorySkill.CreatedAt.String(),
			UpdatedAt: categorySkill.UpdatedAt.String(),
		})
	}

	return results, nil
}

func (q queryResolver) FindByIDCategorySkill(ctx context.Context, id int64) (*model.CategorySkill, error) {
	categorySkill, err := q.CategorySkillService.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.CategorySkill{
		ID:        categorySkill.ID,
		Name:      categorySkill.Name,
		CreatedAt: categorySkill.CreatedAt.String(),
		UpdatedAt: categorySkill.UpdatedAt.String(),
	}, nil
}

func (m mutationResolver) CreateCategorySkill(ctx context.Context, input model.CreateCategorySkillRequest) (*model.CategorySkill, error) {
	err := util.ValidateStruct(ctx, input)
	if err != nil {
		return nil, err
	}

	categorySkill, err := m.CategorySkillService.Create(ctx, &entity.CategorySkill{
		Name: input.Name,
	})
	if err != nil {
		return nil, err
	}

	return &model.CategorySkill{
		ID:        categorySkill.ID,
		Name:      categorySkill.Name,
		CreatedAt: categorySkill.CreatedAt.String(),
		UpdatedAt: categorySkill.UpdatedAt.String(),
	}, nil
}

func (m mutationResolver) UpdateCategorySkill(ctx context.Context, input model.UpdateCategorySkillRequest) (*model.CategorySkill, error) {
	err := util.ValidateStruct(ctx, input)
	if err != nil {
		return nil, err
	}

	categorySkill, err := m.CategorySkillService.Update(ctx, &entity.CategorySkill{
		ID:   input.ID,
		Name: *input.Name,
	})
	if err != nil {
		return nil, err
	}

	return &model.CategorySkill{
		ID:        categorySkill.ID,
		Name:      categorySkill.Name,
		CreatedAt: categorySkill.CreatedAt.String(),
		UpdatedAt: categorySkill.UpdatedAt.String(),
	}, nil
}

func (m mutationResolver) DeleteCategorySkill(ctx context.Context, id int64) (bool, error) {
	err := m.CategorySkillService.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c categorySkillResolver) Skills(ctx context.Context, obj *model.CategorySkill) ([]*model.Skill, error) {
	skills, err := GetLoaders(ctx).ListSkillsByCategoryIDs.Load(obj.ID)
	if err != nil {
		return nil, err
	}

	return skills, nil
}
