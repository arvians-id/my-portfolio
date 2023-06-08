package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"sync"
)

func (q queryResolver) FindAllProject(ctx context.Context, name *string) ([]*model.Project, error) {
	var projects []*entity.Project
	var err error

	projects, err = q.ProjectService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	if name != nil && *name != "" {
		projects, err = q.ProjectService.Query(ctx, *name)
		if err != nil {
			return nil, err
		}
	}

	var results []*model.Project
	for _, project := range projects {
		results = append(results, &model.Project{
			ID:          project.ID,
			Category:    project.Category,
			Title:       project.Title,
			Description: project.Description,
			URL:         project.URL,
			IsFeatured:  project.IsFeatured,
			Date:        project.Date,
			WorkingType: project.WorkingType,
			CreatedAt:   project.CreatedAt.String(),
			UpdatedAt:   project.UpdatedAt.String(),
		})
	}

	return results, nil
}

func (q queryResolver) FindAllProjectByCategory(ctx context.Context, category string) ([]*model.Project, error) {
	var projects []*entity.Project

	projects, err := q.ProjectService.FindAllByCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	var results []*model.Project
	for _, project := range projects {
		results = append(results, &model.Project{
			ID:          project.ID,
			Category:    project.Category,
			Title:       project.Title,
			Description: project.Description,
			URL:         project.URL,
			IsFeatured:  project.IsFeatured,
			Date:        project.Date,
			WorkingType: project.WorkingType,
			CreatedAt:   project.CreatedAt.String(),
			UpdatedAt:   project.UpdatedAt.String(),
		})
	}

	return results, nil
}

func (q queryResolver) FindByIDProject(ctx context.Context, id int64) (*model.Project, error) {
	project, err := q.ProjectService.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.Project{
		ID:          project.ID,
		Category:    project.Category,
		Title:       project.Title,
		Description: project.Description,
		URL:         project.URL,
		IsFeatured:  project.IsFeatured,
		Date:        project.Date,
		WorkingType: project.WorkingType,
		CreatedAt:   project.CreatedAt.String(),
		UpdatedAt:   project.UpdatedAt.String(),
	}, nil
}

func (m mutationResolver) CreateProject(ctx context.Context, input model.CreateProjectRequest) (*model.Project, error) {
	var skills []*entity.Skill
	var images []*entity.ProjectImages
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, id := range input.Skills {
			mu.Lock()
			skills = append(skills, &entity.Skill{
				ID: id,
			})
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for _, image := range input.Images {
			mu.Lock()
			images = append(images, &entity.ProjectImages{
				Image: image.Image,
			})
			mu.Unlock()
		}
	}()
	wg.Wait()

	project, err := m.ProjectService.Create(ctx, &entity.Project{
		Category:    input.Category,
		Title:       input.Title,
		Description: input.Description,
		URL:         input.URL,
		IsFeatured:  input.IsFeatured,
		Date:        input.Date,
		WorkingType: input.WorkingType,
		Skills:      skills,
		Images:      images,
	})
	if err != nil {
		return nil, err
	}

	return &model.Project{
		ID:          project.ID,
		Category:    project.Category,
		Title:       project.Title,
		Description: project.Description,
		URL:         project.URL,
		IsFeatured:  project.IsFeatured,
		Date:        project.Date,
		WorkingType: project.WorkingType,
		CreatedAt:   project.CreatedAt.String(),
		UpdatedAt:   project.UpdatedAt.String(),
	}, nil
}

func (m mutationResolver) UpdateProject(ctx context.Context, input model.UpdateProjectRequest) (*model.Project, error) {
	var skills []*entity.Skill
	var images []*entity.ProjectImages
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, id := range input.Skills {
			mu.Lock()
			skills = append(skills, &entity.Skill{
				ID: id,
			})
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for _, image := range input.Images {
			mu.Lock()
			images = append(images, &entity.ProjectImages{
				Image: image.Image,
			})
			mu.Unlock()
		}
	}()
	wg.Wait()

	project, err := m.ProjectService.Update(ctx, &entity.Project{
		ID:          input.ID,
		Category:    input.Category,
		Title:       input.Title,
		Description: input.Description,
		URL:         input.URL,
		IsFeatured:  input.IsFeatured,
		Date:        input.Date,
		WorkingType: input.WorkingType,
		Skills:      skills,
		Images:      images,
	})
	if err != nil {
		return nil, err
	}

	return &model.Project{
		ID:          project.ID,
		Category:    project.Category,
		Title:       project.Title,
		Description: project.Description,
		URL:         project.URL,
		IsFeatured:  project.IsFeatured,
		Date:        project.Date,
		WorkingType: project.WorkingType,
		CreatedAt:   project.CreatedAt.String(),
		UpdatedAt:   project.UpdatedAt.String(),
	}, nil
}

func (m mutationResolver) DeleteProject(ctx context.Context, id int64) (bool, error) {
	err := m.ProjectService.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p projectResolver) Skills(ctx context.Context, obj *model.Project) ([]*model.Skill, error) {
	skills, err := GetLoaders(ctx).ListSkillsByProjectIDs.Load(obj.ID)
	if err != nil {
		return nil, err
	}

	return skills, nil
}

func (p projectResolver) Images(ctx context.Context, obj *model.Project) ([]*model.ProjectImages, error) {
	//TODO implement me
	panic("implement me")
}
