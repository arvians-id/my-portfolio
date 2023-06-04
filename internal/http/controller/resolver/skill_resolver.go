package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
)

func (q queryResolver) FindAllSkill(ctx context.Context) ([]*model.Skill, error) {
	//TODO implement me
	panic("implement me")
}

func (q queryResolver) FindByIDSkill(ctx context.Context, id int64) (*model.Skill, error) {
	//TODO implement me
	panic("implement me")
}

func (m mutationResolver) CreateSkill(ctx context.Context, input model.CreateSkillRequest) (*model.Skill, error) {
	//TODO implement me
	panic("implement me")
}

func (m mutationResolver) UpdateSkill(ctx context.Context, input model.UpdateSkillRequest) (*model.Skill, error) {
	//TODO implement me
	panic("implement me")
}

func (m mutationResolver) DeleteSkill(ctx context.Context, id int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (q skillResolver) CategorySkill(ctx context.Context, obj *model.Skill) (*model.CategorySkill, error) {
	//TODO implement me
	panic("implement me")
}
