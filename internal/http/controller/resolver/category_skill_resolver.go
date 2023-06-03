package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
)

func (q queryResolver) CategorySkillFindAll(ctx context.Context) ([]*model.CategorySkill, error) {
	//TODO implement me
	panic("implement me")
}

func (q queryResolver) CategorySkillFindByID(ctx context.Context, id int64) (*model.CategorySkill, error) {
	//TODO implement me
	panic("implement me")
}

func (c categorySkillResolver) Skills(ctx context.Context, obj *model.CategorySkill) ([]*model.Skill, error) {
	//TODO implement me
	panic("implement me")
}
