package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"github.com/arvians-id/go-portfolio/util"
)

type ProjectImageServiceContract interface {
	FindAllByProjectIDs(ctx context.Context, projectIDs []int64) ([]*entity.ProjectImage, error)
	Delete(ctx context.Context, id int64) error
}

type ProjectImageService struct {
	ProjectImageRepository repository.ProjectImageRepositoryContract
}

func NewProjectImageService(projectImageRepository repository.ProjectImageRepositoryContract) ProjectImageServiceContract {
	return &ProjectImageService{
		ProjectImageRepository: projectImageRepository,
	}
}

func (service *ProjectImageService) FindAllByProjectIDs(ctx context.Context, projectIDs []int64) ([]*entity.ProjectImage, error) {
	projects, err := service.ProjectImageRepository.FindAllByProjectIDs(ctx, projectIDs)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (service *ProjectImageService) Delete(ctx context.Context, id int64) error {
	projectImageCheck, err := service.ProjectImageRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	err = service.ProjectImageRepository.Delete(ctx, projectImageCheck.ID, projectImageCheck.ProjectID)
	if err != nil {
		return err
	}

	path := "images/project"
	err = util.DeleteFile(path, projectImageCheck.Image)
	if err != nil {
		return err
	}

	return nil
}
