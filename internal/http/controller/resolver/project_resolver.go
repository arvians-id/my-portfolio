package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
)

func (q queryResolver) FindAllProject(ctx context.Context) ([]*model.Project, error) {
	projects, err := q.ProjectService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (q queryResolver) FindByIDProject(ctx context.Context, id int64) (*model.Project, error) {
	project, err := q.ProjectService.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (m mutationResolver) CreateProject(ctx context.Context, input model.CreateProjectRequest) (*model.Project, error) {
	project, err := m.ProjectService.Create(ctx, &input)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (m mutationResolver) UpdateProject(ctx context.Context, input model.UpdateProjectRequest) (*model.Project, error) {
	project, err := m.ProjectService.Update(ctx, &input)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (m mutationResolver) DeleteProject(ctx context.Context, id int64) (bool, error) {
	err := m.ProjectService.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p projectResolver) Skills(ctx context.Context, obj *model.Project) ([]*model.Skill, error) {
	//TODO implement me
	panic("implement me")
}
