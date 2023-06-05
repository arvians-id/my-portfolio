package repository

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"gorm.io/gorm"
	"log"
)

type EducationRepositoryContract interface {
	FindAll(ctx context.Context) ([]*entity.Education, error)
	FindByID(ctx context.Context, id int64) (*entity.Education, error)
	Create(ctx context.Context, request *model.CreateEducationRequest) (*entity.Education, error)
	Update(ctx context.Context, request *model.UpdateEducationRequest) error
	Delete(ctx context.Context, id int64) error
}

type EducationRepository struct {
	DB *gorm.DB
}

func NewEducationRepository(db *gorm.DB) EducationRepositoryContract {
	return &EducationRepository{DB: db}
}

func (repository *EducationRepository) FindAll(ctx context.Context) ([]*entity.Education, error) {
	query := "SELECT * FROM educations"
	var educations []*entity.Education
	err := repository.DB.WithContext(ctx).Raw(query).Scan(&educations).Error
	if err != nil {
		log.Println("[EducationRepository][FindAll] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return educations, nil
}

func (repository *EducationRepository) FindByID(ctx context.Context, id int64) (*entity.Education, error) {
	query := "SELECT * FROM educations WHERE id = ?"
	var education entity.Education
	row := repository.DB.WithContext(ctx).Raw(query, id).Row()
	err := row.Scan(
		&education.ID,
		&education.Institution,
		&education.Degree,
		&education.FieldOfStudy,
		&education.Grade,
		&education.Description,
		&education.StartDate,
		&education.EndDate,
	)
	if err != nil {
		log.Println("[EducationRepository][FindByID] problem with scanning db row, err: ", err.Error())
		return nil, err
	}

	return &education, nil
}

func (repository *EducationRepository) Create(ctx context.Context, request *model.CreateEducationRequest) (*entity.Education, error) {
	var education entity.Education
	education.Institution = request.Institution
	education.Degree = request.Degree
	education.FieldOfStudy = request.FieldOfStudy
	education.Grade = request.Grade
	education.Description = request.Description
	education.StartDate = request.StartDate
	education.EndDate = request.EndDate

	err := repository.DB.WithContext(ctx).Create(&education).Error
	if err != nil {
		log.Println("[EducationRepository][Create] problem with scanning db row, err: ", err.Error())
		return nil, err
	}

	return &education, nil
}

func (repository *EducationRepository) Update(ctx context.Context, request *model.UpdateEducationRequest) error {
	var education entity.Education
	education.ID = request.ID
	education.Institution = *request.Institution
	education.Degree = *request.Degree
	education.FieldOfStudy = *request.FieldOfStudy
	education.Grade = *request.Grade
	education.Description = request.Description
	education.StartDate = *request.StartDate
	education.EndDate = request.EndDate

	err := repository.DB.WithContext(ctx).Updates(&education).Error
	if err != nil {
		log.Println("[EducationRepository][Update] problem querying to db, err: ", err.Error())
		return err
	}

	return nil
}

func (repository *EducationRepository) Delete(ctx context.Context, id int64) error {
	err := repository.DB.WithContext(ctx).Delete(&entity.Education{}, id).Error
	if err != nil {
		log.Println("[EducationRepository][Delete] problem querying to db, err: ", err.Error())
		return err
	}

	return nil
}
