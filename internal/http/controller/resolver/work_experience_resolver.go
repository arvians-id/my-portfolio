package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
)

func (q queryResolver) FindAllWorkExperience(ctx context.Context) ([]*model.WorkExperience, error) {
	//TODO implement me
	panic("implement me")
}

func (q queryResolver) FindByIDWorkExperience(ctx context.Context, id int64) (*model.WorkExperience, error) {
	//TODO implement me
	panic("implement me")
}

func (m mutationResolver) CreateWorkExperience(ctx context.Context, input model.CreateWorkExperienceRequest) (*model.WorkExperience, error) {
	//TODO implement me
	panic("implement me")
}

func (m mutationResolver) UpdateWorkExperience(ctx context.Context, input model.UpdateWorkExperienceRequest) (*model.WorkExperience, error) {
	//TODO implement me
	panic("implement me")
}

func (m mutationResolver) DeleteWorkExperience(ctx context.Context, id int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (w workExperienceResolver) Skills(ctx context.Context, obj *model.WorkExperience) ([]*model.Skill, error) {
	//TODO implement me
	panic("implement me")
}
