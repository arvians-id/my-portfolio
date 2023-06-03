package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
)

func (q queryResolver) ProjectFindAll(ctx context.Context) ([]*model.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (q queryResolver) ProjectFindByID(ctx context.Context, id int64) (*model.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p projectResolver) Skills(ctx context.Context, obj *model.Project) ([]*model.Skill, error) {
	//TODO implement me
	panic("implement me")
}
