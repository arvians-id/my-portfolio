package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
)

func (q queryResolver) FindAllProject(ctx context.Context) ([]*model.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (q queryResolver) FindByIDProject(ctx context.Context, id int64) (*model.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (m mutationResolver) CreateProject(ctx context.Context, input model.CreateProjectRequest) (*model.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (m mutationResolver) UpdateProject(ctx context.Context, input model.UpdateProjectRequest) (*model.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (m mutationResolver) DeleteProject(ctx context.Context, id int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (p projectResolver) Skills(ctx context.Context, obj *model.Project) ([]*model.Skill, error) {
	//TODO implement me
	panic("implement me")
}
