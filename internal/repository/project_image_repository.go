package repository

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"gorm.io/gorm"
	"log"
)

type ProjectImageRepositoryContract interface {
	FindAllByProjectIDs(ctx context.Context, projectIDs []int64) ([]*entity.ProjectImage, error)
	FindAllByProjectID(ctx context.Context, projectID int64) ([]*entity.ProjectImage, error)
	FindByID(ctx context.Context, id int64) (*entity.ProjectImage, error)
	Delete(ctx context.Context, id int64, projectID int64) error
}

type ProjectImageRepository struct {
	DB *gorm.DB
}

func NewProjectImageRepository(db *gorm.DB) ProjectImageRepositoryContract {
	return &ProjectImageRepository{
		DB: db,
	}
}

func (repository *ProjectImageRepository) FindAllByProjectIDs(ctx context.Context, projectIDs []int64) ([]*entity.ProjectImage, error) {
	query := "SELECT * FROM project_images WHERE project_id IN (?)"
	var projectImages []*entity.ProjectImage
	err := repository.DB.WithContext(ctx).Raw(query, projectIDs).Scan(&projectImages).Error
	if err != nil {
		log.Println("[ProjectImageRepository][FindImagesByProjectID] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return projectImages, nil
}

func (repository *ProjectImageRepository) FindAllByProjectID(ctx context.Context, projectID int64) ([]*entity.ProjectImage, error) {
	query := "SELECT * FROM project_images WHERE project_id = ?"
	var projectImages []*entity.ProjectImage
	err := repository.DB.WithContext(ctx).Raw(query, projectID).Scan(&projectImages).Error
	if err != nil {
		log.Println("[ProjectImageRepository][FindImagesByProjectID] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return projectImages, nil
}

func (repository *ProjectImageRepository) FindByID(ctx context.Context, id int64) (*entity.ProjectImage, error) {
	query := "SELECT * FROM project_images WHERE id = ?"
	var projectImage entity.ProjectImage
	row := repository.DB.WithContext(ctx).Raw(query, id).Row()
	err := row.Scan(&projectImage.ID, &projectImage.ProjectID, &projectImage.Image)
	if err != nil {
		log.Println("[ProjectImageRepository][FindByID] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return &projectImage, nil
}

func (repository *ProjectImageRepository) Delete(ctx context.Context, id int64, projectID int64) error {
	err := repository.DB.WithContext(ctx).Where("project_id = ?", projectID).Delete(&entity.ProjectImage{}, id).Error
	if err != nil {
		log.Println("[ProjectImageRepository][DeleteImage] problem querying to db, err: ", err.Error())
		return err
	}

	return nil
}
