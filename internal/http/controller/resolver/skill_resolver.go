package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
)

func (q queryResolver) SkillFindAll(ctx context.Context) ([]*model.Skill, error) {
	//TODO implement me
	panic("implement me")
}

func (q queryResolver) SkillFindByID(ctx context.Context, id int64) (*model.Skill, error) {
	//TODO implement me
	panic("implement me")
}

func (q skillResolver) CategorySkill(ctx context.Context, obj *model.Skill) (*model.CategorySkill, error) {
	//TODO implement me
	panic("implement me")
}
