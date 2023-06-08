package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
)

func (q queryResolver) FindAllSkill(ctx context.Context) ([]*model.Skill, error) {
	skills, err := q.SkillService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var results []*model.Skill
	for _, skill := range skills {
		results = append(results, &model.Skill{
			ID:              skill.ID,
			CategorySkillID: skill.CategorySkillID,
			Name:            skill.Name,
			Icon:            skill.Icon,
		})
	}

	return results, nil
}

func (q queryResolver) FindByIDSkill(ctx context.Context, id int64) (*model.Skill, error) {
	skill, err := q.SkillService.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.Skill{
		ID:              skill.ID,
		CategorySkillID: skill.CategorySkillID,
		Name:            skill.Name,
		Icon:            skill.Icon,
	}, nil
}

func (m mutationResolver) CreateSkill(ctx context.Context, input model.CreateSkillRequest) (*model.Skill, error) {
	skill, err := m.SkillService.Create(ctx, &entity.Skill{
		CategorySkillID: input.CategorySkillID,
		Name:            input.Name,
		Icon:            input.Icon,
	})
	if err != nil {
		return nil, err
	}

	return &model.Skill{
		ID:              skill.ID,
		CategorySkillID: skill.CategorySkillID,
		Name:            skill.Name,
		Icon:            skill.Icon,
	}, nil
}

func (m mutationResolver) UpdateSkill(ctx context.Context, input model.UpdateSkillRequest) (*model.Skill, error) {
	skill, err := m.SkillService.Update(ctx, &entity.Skill{
		ID:              input.ID,
		CategorySkillID: input.CategorySkillID,
		Name:            input.Name,
		Icon:            input.Icon,
	})
	if err != nil {
		return nil, err
	}

	return &model.Skill{
		ID:              skill.ID,
		CategorySkillID: skill.CategorySkillID,
		Name:            skill.Name,
		Icon:            skill.Icon,
	}, nil
}

func (m mutationResolver) DeleteSkill(ctx context.Context, id int64) (bool, error) {
	err := m.SkillService.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (q skillResolver) CategorySkill(ctx context.Context, obj *model.Skill) (*model.CategorySkill, error) {
	categorySkill, err := GetLoaders(ctx).ListCategoryBySkillIDs.Load(obj.ID)
	if err != nil {
		return nil, err
	}

	return categorySkill, nil
}
