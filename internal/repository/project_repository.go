package repository

import (
	"context"
	"errors"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"gorm.io/gorm"
	"log"
)

type ProjectRepositoryContract interface {
	FindAll(ctx context.Context) ([]*entity.Project, error)
	FindAllByIDs(ctx context.Context, id []int64) ([]*entity.Project, error)
	FindAllByCategory(ctx context.Context, category string) ([]*entity.Project, error)
	FindByID(ctx context.Context, id int64) (*entity.Project, error)
	Create(ctx context.Context, project *entity.Project) (*entity.Project, error)
	Update(ctx context.Context, project *entity.Project) error
	Delete(ctx context.Context, id int64) error
}

type ProjectRepository struct {
	DB *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepositoryContract {
	return &ProjectRepository{
		DB: db,
	}
}
func (repository *ProjectRepository) FindAll(ctx context.Context) ([]*entity.Project, error) {
	query := "SELECT * FROM projects ORDER BY created_at DESC"
	var projects []*entity.Project
	err := repository.DB.WithContext(ctx).Raw(query).Scan(&projects).Error
	if err != nil {
		log.Println("[ProjectRepository][FindAll] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return projects, nil
}

func (repository *ProjectRepository) FindAllByIDs(ctx context.Context, id []int64) ([]*entity.Project, error) {
	query := "SELECT * FROM projects WHERE id IN (?) ORDER BY created_at DESC"
	var projects []*entity.Project
	err := repository.DB.WithContext(ctx).Raw(query, id).Scan(&projects).Error
	if err != nil {
		log.Println("[ProjectRepository][FindAllByCategory] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return projects, nil
}

func (repository *ProjectRepository) FindAllByCategory(ctx context.Context, category string) ([]*entity.Project, error) {
	query := "SELECT * FROM projects WHERE category = ? ORDER BY created_at DESC"
	var projects []*entity.Project
	err := repository.DB.WithContext(ctx).Raw(query, category).Scan(&projects).Error
	if err != nil {
		log.Println("[ProjectRepository][FindAllByCategory] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return projects, nil
}

func (repository *ProjectRepository) FindByID(ctx context.Context, id int64) (*entity.Project, error) {
	query := "SELECT * FROM projects WHERE id = ?"
	var project entity.Project
	row := repository.DB.WithContext(ctx).Raw(query, id).Row()
	err := row.Scan(
		&project.ID,
		&project.Category,
		&project.Title,
		&project.Description,
		&project.URL,
		&project.IsFeatured,
		&project.Date,
		&project.WorkingType,
		&project.CreatedAt,
		&project.UpdatedAt,
	)
	if err != nil {
		log.Println("[ProjectRepository][FindByID] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return &project, nil
}

func (repository *ProjectRepository) Create(ctx context.Context, project *entity.Project) (*entity.Project, error) {
	err := repository.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var skills []*entity.Skill
		for _, id := range project.Skills {
			var skill entity.Skill
			query := "SELECT id FROM skills WHERE id = ? LIMIT 1"
			row := tx.WithContext(ctx).Raw(query, id.ID).Row()
			err := row.Scan(&skill.ID)
			if err != nil {
				log.Println("[ProjectRepository][Scan] problem querying to db, err: ", err.Error())
				return errors.New("sql: there are skills that are not registered")
			}

			skills = append(skills, &skill)
		}

		project.Skills = skills
		err := tx.WithContext(ctx).Create(&project).Error
		if err != nil {
			log.Println("[ProjectRepository][Create] problem querying to db, err: ", err.Error())
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (repository *ProjectRepository) Update(ctx context.Context, project *entity.Project) error {
	err := repository.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var skills []*entity.Skill
		for _, id := range project.Skills {
			var skill entity.Skill
			query := "SELECT id FROM skills WHERE id = ? LIMIT 1"
			row := tx.WithContext(ctx).Raw(query, id.ID).Row()
			err := row.Scan(&skill.ID)
			if err != nil {
				log.Println("[ProjectRepository][Scan] problem querying to db, err: ", err.Error())
				return errors.New("sql: there are skills that are not registered")
			}

			skills = append(skills, &skill)
		}

		project.Skills = skills
		err := tx.WithContext(ctx).Updates(&project).Error
		if err != nil {
			log.Println("[ProjectRepository][Updates] problem querying to db, err: ", err.Error())
			return err
		}

		err = tx.WithContext(ctx).Model(&project).Association("Skills").Replace(skills)
		if err != nil {
			log.Println("[ProjectRepository][Replace] problem querying to db, err: ", err.Error())
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (repository *ProjectRepository) Delete(ctx context.Context, id int64) error {
	err := repository.DB.WithContext(ctx).Delete(&entity.Project{}, id).Error
	if err != nil {
		log.Println("[ProjectRepository][Delete] problem querying to db, err: ", err.Error())
		return err
	}

	return nil
}
