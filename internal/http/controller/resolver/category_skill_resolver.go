package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
)

func (q queryResolver) FindAllCategorySkill(ctx context.Context) ([]*model.CategorySkill, error) {
	//TODO implement me
	panic("implement me")
}

func (q queryResolver) FindByIDCategorySkill(ctx context.Context, id int64) (*model.CategorySkill, error) {
	//TODO implement me
	panic("implement me")
}

func (m mutationResolver) CreateCategorySkill(ctx context.Context, input model.CreateCategorySkillRequest) (*model.CategorySkill, error) {
	//TODO implement me
	panic("implement me")
}

func (m mutationResolver) UpdateCategorySkill(ctx context.Context, input model.UpdateCategorySkillRequest) (*model.CategorySkill, error) {
	//TODO implement me
	panic("implement me")
}

func (m mutationResolver) DeleteCategorySkill(ctx context.Context, id int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c categorySkillResolver) Skills(ctx context.Context, obj *model.CategorySkill) ([]*model.Skill, error) {
	//TODO implement me
	panic("implement me")
}
