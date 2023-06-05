package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"log"
)

type ProjectServiceContract interface {
	FindAll(ctx context.Context) ([]*model.Project, error)
	FindAllByCategory(ctx context.Context, category string) ([]*model.Project, error)
	FindByID(ctx context.Context, id int64) (*model.Project, error)
	Create(ctx context.Context, request *model.CreateProjectRequest) (*model.Project, error)
	Update(ctx context.Context, request *model.UpdateProjectRequest) (*model.Project, error)
	Delete(ctx context.Context, id int64) error
}

type ProjectService struct {
	ProjectRepository repository.ProjectRepositoryContract
}

func NewProjectService(projectRepository repository.ProjectRepositoryContract) ProjectServiceContract {
	return &ProjectService{
		ProjectRepository: projectRepository,
	}
}

func (repository *ProjectService) FindAll(ctx context.Context) ([]*model.Project, error) {
	projects, err := repository.ProjectRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var results []*model.Project
	for _, project := range projects {
		results = append(results, project.ToModel())
	}

	return results, nil
}

func (repository *ProjectService) FindAllByCategory(ctx context.Context, category string) ([]*model.Project, error) {
	projects, err := repository.ProjectRepository.FindAllByCategory(ctx, category)
	if err != nil {
		log.Println("[ProjectService][FindAll] problem calling repository, err: ", err.Error())
		return nil, err
	}

	var results []*model.Project
	for _, project := range projects {
		results = append(results, project.ToModel())
	}

	return results, nil

}

func (repository *ProjectService) FindByID(ctx context.Context, id int64) (*model.Project, error) {
	project, err := repository.ProjectRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[ProjectService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return project.ToModel(), nil
}

func (repository *ProjectService) Create(ctx context.Context, request *model.CreateProjectRequest) (*model.Project, error) {
	project, err := repository.ProjectRepository.Create(ctx, request)
	if err != nil {
		log.Println("[ProjectService][Create] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return project.ToModel(), nil
}

func (repository *ProjectService) Update(ctx context.Context, request *model.UpdateProjectRequest) (*model.Project, error) {
	_, err := repository.ProjectRepository.FindByID(ctx, request.ID)
	if err != nil {
		log.Println("[ProjectService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	err = repository.ProjectRepository.Update(ctx, request)
	if err != nil {
		log.Println("[ProjectService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	projectUpdated, err := repository.ProjectRepository.FindByID(ctx, request.ID)
	if err != nil {
		log.Println("[ProjectService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return projectUpdated.ToModel(), nil
}

func (repository *ProjectService) Delete(ctx context.Context, id int64) error {
	projectCheck, err := repository.ProjectRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[ProjectService][Delete] problem calling repository, err: ", err.Error())
		return err
	}

	err = repository.ProjectRepository.Delete(ctx, projectCheck.ID)
	if err != nil {
		log.Println("[ProjectService][Delete] problem calling repository, err: ", err.Error())
		return err
	}

	return nil
}
