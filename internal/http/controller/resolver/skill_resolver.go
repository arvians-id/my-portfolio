package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
)

func (q queryResolver) FindAllSkill(ctx context.Context) ([]*model.Skill, error) {
	skills, err := q.SkillService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return skills, nil
}

func (q queryResolver) FindByIDSkill(ctx context.Context, id int64) (*model.Skill, error) {
	skill, err := q.SkillService.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return skill, nil
}

func (m mutationResolver) CreateSkill(ctx context.Context, input model.CreateSkillRequest) (*model.Skill, error) {
	skill, err := m.SkillService.Create(ctx, &input)
	if err != nil {
		return nil, err
	}

	return skill, nil
}

func (m mutationResolver) UpdateSkill(ctx context.Context, input model.UpdateSkillRequest) (*model.Skill, error) {
	skill, err := m.SkillService.Update(ctx, &input)
	if err != nil {
		return nil, err
	}

	return skill, nil
}

func (m mutationResolver) DeleteSkill(ctx context.Context, id int64) (bool, error) {
	err := m.SkillService.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (q skillResolver) CategorySkill(ctx context.Context, obj *model.Skill) (*model.CategorySkill, error) {
	//TODO implement me
	panic("implement me")
}
