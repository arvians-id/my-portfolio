package repository

import (
	"context"
	"errors"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"gorm.io/gorm"
	"log"
)

type WorkExperienceRepositoryContract interface {
	FindAll(ctx context.Context) ([]*entity.WorkExperience, error)
	FindByID(ctx context.Context, id int64) (*entity.WorkExperience, error)
	Create(ctx context.Context, request *model.CreateWorkExperienceRequest) (*entity.WorkExperience, error)
	Update(ctx context.Context, request *model.UpdateWorkExperienceRequest) error
	Delete(ctx context.Context, id int64) error
}

type WorkExperienceRepository struct {
	DB *gorm.DB
}

func NewWorkExperienceRepository(db *gorm.DB) WorkExperienceRepositoryContract {
	return &WorkExperienceRepository{
		DB: db,
	}
}

func (repository *WorkExperienceRepository) FindAll(ctx context.Context) ([]*entity.WorkExperience, error) {
	query := "SELECT * FROM work_experiences ORDER BY created_at DESC"
	var workExperiences []*entity.WorkExperience
	err := repository.DB.WithContext(ctx).Raw(query).Scan(&workExperiences).Error
	if err != nil {
		log.Println("[WorkExperienceRepository][FindAll] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return workExperiences, nil
}

func (repository *WorkExperienceRepository) FindByID(ctx context.Context, id int64) (*entity.WorkExperience, error) {
	query := "SELECT * FROM work_experiences WHERE id = ?"
	var workExperience entity.WorkExperience
	row := repository.DB.WithContext(ctx).Raw(query, id).Row()
	err := row.Scan(
		&workExperience.ID,
		&workExperience.Role,
		&workExperience.Company,
		&workExperience.Description,
		&workExperience.StartDate,
		&workExperience.EndDate,
		&workExperience.JobType,
		&workExperience.CreatedAt,
		&workExperience.UpdatedAt,
	)
	if err != nil {
		log.Println("[WorkExperienceRepository][FindByID] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return &workExperience, nil
}

func (repository *WorkExperienceRepository) Create(ctx context.Context, request *model.CreateWorkExperienceRequest) (*entity.WorkExperience, error) {
	var workExperience entity.WorkExperience
	err := repository.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var skills []*entity.Skill
		for _, id := range request.Skills {
			query := "SELECT id FROM skills WHERE id = ? LIMIT 1"
			var skill entity.Skill
			row := tx.WithContext(ctx).Raw(query, id).Row()
			err := row.Scan(&skill.ID)
			if err != nil {
				log.Println("[WorkExperienceRepository][Scan] problem querying to db, err: ", err.Error())
				return errors.New("sql: there are skills that are not registered")
			}

			skills = append(skills, &skill)
		}

		workExperience.Role = request.Role
		workExperience.Company = request.Company
		workExperience.Description = request.Description
		workExperience.StartDate = request.StartDate
		workExperience.EndDate = request.EndDate
		workExperience.JobType = request.JobType
		workExperience.Skills = skills

		err := tx.WithContext(ctx).Create(&workExperience).Error
		if err != nil {
			log.Println("[WorkExperienceRepository][Create] problem querying to db, err: ", err.Error())
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &workExperience, nil
}

func (repository *WorkExperienceRepository) Update(ctx context.Context, request *model.UpdateWorkExperienceRequest) error {
	var workExperience entity.WorkExperience
	err := repository.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var skills []*entity.Skill
		for _, id := range request.Skills {
			query := "SELECT id FROM skills WHERE id = ? LIMIT 1"
			var skill entity.Skill
			row := tx.WithContext(ctx).Raw(query, id).Row()
			err := row.Scan(&skill.ID)
			if err != nil {
				log.Println("[WorkExperienceRepository][Scan] problem querying to db, err: ", err.Error())
				return errors.New("sql: there are skills that are not registered")
			}

			skills = append(skills, &skill)
		}

		workExperience.ID = request.ID
		workExperience.Role = *request.Role
		workExperience.Company = *request.Company
		workExperience.Description = request.Description
		workExperience.StartDate = *request.StartDate
		workExperience.EndDate = request.EndDate
		workExperience.JobType = *request.JobType
		workExperience.Skills = skills

		err := tx.WithContext(ctx).Updates(&workExperience).Error
		if err != nil {
			log.Println("[WorkExperienceRepository][Updates] problem querying to db, err: ", err.Error())
			return err
		}

		err = tx.WithContext(ctx).Model(&workExperience).Association("Skills").Replace(skills)
		if err != nil {
			log.Println("[WorkExperienceRepository][Replace] problem querying to db, err: ", err.Error())
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (repository *WorkExperienceRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM work_experiences WHERE id = ?"
	err := repository.DB.WithContext(ctx).Raw(query, id).Error
	if err != nil {
		log.Println("[WorkExperienceRepository][Delete] problem querying to db, err: ", err.Error())
		return err
	}

	return nil
}
