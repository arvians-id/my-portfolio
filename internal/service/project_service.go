package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"github.com/arvians-id/go-portfolio/util"
	"log"
	"strconv"
)

type ProjectServiceContract interface {
	Query(ctx context.Context, query string) ([]*entity.Project, error)
	FindAll(ctx context.Context) ([]*entity.Project, error)
	FindAllByCategory(ctx context.Context, category string) ([]*entity.Project, error)
	FindByID(ctx context.Context, id int64) (*entity.Project, error)
	Create(ctx context.Context, project *entity.Project) (*entity.Project, error)
	Update(ctx context.Context, project *entity.Project) (*entity.Project, error)
	Delete(ctx context.Context, id int64) error
}

type ProjectService struct {
	ProjectRepository repository.ProjectRepositoryContract
	BleveSearch       BleveSearchServiceContract
}

func NewProjectService(projectRepository repository.ProjectRepositoryContract, bleveSearch BleveSearchServiceContract) ProjectServiceContract {
	return &ProjectService{
		ProjectRepository: projectRepository,
		BleveSearch:       bleveSearch,
	}
}

func (repository *ProjectService) Query(ctx context.Context, query string) ([]*entity.Project, error) {
	ids, err := repository.BleveSearch.Search(query)
	if err != nil {
		log.Println("[ProjectService][Query] problem calling bleve search service, err: ", err.Error())
		return nil, err
	}

	var idsInt64 []int64
	if len(ids) > 0 {
		for _, id := range ids {
			i, _ := strconv.Atoi(id)
			idsInt64 = append(idsInt64, int64(i))
		}
	}

	projects, err := repository.ProjectRepository.FindAllByIDs(ctx, idsInt64)
	if err != nil {
		log.Println("[ProjectService][Query] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return projects, nil
}

func (repository *ProjectService) FindAll(ctx context.Context) ([]*entity.Project, error) {
	projects, err := repository.ProjectRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (repository *ProjectService) FindAllByCategory(ctx context.Context, category string) ([]*entity.Project, error) {
	projects, err := repository.ProjectRepository.FindAllByCategory(ctx, category)
	if err != nil {
		log.Println("[ProjectService][FindAll] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return projects, nil
}

func (repository *ProjectService) FindByID(ctx context.Context, id int64) (*entity.Project, error) {
	project, err := repository.ProjectRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[ProjectService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return project, nil
}

func (repository *ProjectService) Create(ctx context.Context, project *entity.Project) (*entity.Project, error) {
	project, err := repository.ProjectRepository.Create(ctx, project)
	if err != nil {
		log.Println("[ProjectService][Create] problem calling repository, err: ", err.Error())
		return nil, err
	}

	var searchItem entity.SearchItem
	searchItem.ID = strconv.FormatInt(project.ID, 10)
	searchItem.Title = project.Title
	err = repository.BleveSearch.InsertOrUpdate(&searchItem)
	if err != nil {
		log.Println("[ProjectService][InsertBleve] problem calling bleve search service, err: ", err.Error())
		return nil, err
	}

	return project, nil
}

func (repository *ProjectService) Update(ctx context.Context, project *entity.Project) (*entity.Project, error) {
	_, err := repository.ProjectRepository.FindByID(ctx, project.ID)
	if err != nil {
		log.Println("[ProjectService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	err = repository.ProjectRepository.Update(ctx, project)
	if err != nil {
		log.Println("[ProjectService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	projectUpdated, err := repository.ProjectRepository.FindByID(ctx, project.ID)
	if err != nil {
		log.Println("[ProjectService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	var searchItem entity.SearchItem
	searchItem.ID = strconv.FormatInt(project.ID, 10)
	searchItem.Title = project.Title
	err = repository.BleveSearch.InsertOrUpdate(&searchItem)
	if err != nil {
		log.Println("[ProjectService][UpdateBleve] problem calling bleve search service, err: ", err.Error())
		return nil, err
	}

	return projectUpdated, nil
}

func (repository *ProjectService) Delete(ctx context.Context, id int64) error {
	projectCheck, err := repository.ProjectRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[ProjectService][Delete] problem calling repository, err: ", err.Error())
		return err
	}

	images, err := repository.ProjectRepository.FindAllImagesByProjectID(ctx, projectCheck.ID)
	if err != nil {
		log.Println("[ProjectService][FindAllImagesByProjectID] problem calling repository, err: ", err.Error())
		return err
	}

	err = repository.ProjectRepository.Delete(ctx, projectCheck.ID)
	if err != nil {
		log.Println("[ProjectService][Delete] problem calling repository, err: ", err.Error())
		return err
	}

	if images != nil {
		for _, image := range images {
			path := "images/project"
			err = util.DeleteFile(path, image.Image)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
