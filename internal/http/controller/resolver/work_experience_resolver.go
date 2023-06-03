package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
)

func (q queryResolver) WorkExperienceFindAll(ctx context.Context) ([]*model.WorkExperience, error) {
	//TODO implement me
	panic("implement me")
}

func (q queryResolver) WorkExperienceFindByID(ctx context.Context, id int64) (*model.WorkExperience, error) {
	//TODO implement me
	panic("implement me")
}

func (w workExperienceResolver) Skills(ctx context.Context, obj *model.WorkExperience) ([]*model.Skill, error) {
	//TODO implement me
	panic("implement me")
}
