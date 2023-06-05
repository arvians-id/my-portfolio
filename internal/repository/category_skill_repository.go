package repository

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"gorm.io/gorm"
	"log"
)

type CategorySkillRepositoryContract interface {
	FindAll(ctx context.Context) ([]*entity.CategorySkill, error)
	FindAllByIDs(ctx context.Context, ids []int64) ([]*entity.CategorySkill, error)
	FindByID(ctx context.Context, id int64) (*entity.CategorySkill, error)
	Create(ctx context.Context, request *model.CreateCategorySkillRequest) (*entity.CategorySkill, error)
	Update(ctx context.Context, request *model.UpdateCategorySkillRequest) error
	Delete(ctx context.Context, id int64) error
}

type CategorySkillRepository struct {
	DB *gorm.DB
}

func NewCategorySkillRepository(db *gorm.DB) CategorySkillRepositoryContract {
	return &CategorySkillRepository{
		DB: db,
	}
}
func (repository *CategorySkillRepository) FindAll(ctx context.Context) ([]*entity.CategorySkill, error) {
	query := "SELECT * FROM category_skills ORDER BY created_at DESC"
	var categorySkills []*entity.CategorySkill
	err := repository.DB.WithContext(ctx).Raw(query).Scan(&categorySkills).Error
	if err != nil {
		log.Println("[CategorySkillRepository][FindAll] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return categorySkills, nil
}

func (repository *CategorySkillRepository) FindAllByIDs(ctx context.Context, ids []int64) ([]*entity.CategorySkill, error) {
	query := "SELECT * FROM category_skills WHERE id IN (?) ORDER BY created_at DESC"
	var categorySkills []*entity.CategorySkill
	err := repository.DB.WithContext(ctx).Raw(query, ids).Scan(&categorySkills).Error
	if err != nil {
		log.Println("[CategorySkillRepository][FindByIDs] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return categorySkills, nil
}

func (repository *CategorySkillRepository) FindByID(ctx context.Context, id int64) (*entity.CategorySkill, error) {
	query := "SELECT * FROM category_skills WHERE id = ?"
	var categorySkill entity.CategorySkill
	row := repository.DB.WithContext(ctx).Raw(query, id).Row()
	err := row.Scan(&categorySkill.ID, &categorySkill.Name, &categorySkill.CreatedAt, &categorySkill.UpdatedAt)
	if err != nil {
		log.Println("[CategorySkillRepository][FindByID] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return &categorySkill, nil
}

func (repository *CategorySkillRepository) Create(ctx context.Context, request *model.CreateCategorySkillRequest) (*entity.CategorySkill, error) {
	var categorySkill entity.CategorySkill
	categorySkill.Name = request.Name

	err := repository.DB.WithContext(ctx).Create(&categorySkill).Error
	if err != nil {
		log.Println("[CategorySkillRepository][Create] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return &categorySkill, nil
}

func (repository *CategorySkillRepository) Update(ctx context.Context, request *model.UpdateCategorySkillRequest) error {
	var categorySkill entity.CategorySkill
	categorySkill.ID = request.ID
	categorySkill.Name = *request.Name

	err := repository.DB.WithContext(ctx).Updates(&categorySkill).Error
	if err != nil {
		log.Println("[CategorySkillRepository][Update] problem querying to db, err: ", err.Error())
		return err
	}

	return nil
}

func (repository *CategorySkillRepository) Delete(ctx context.Context, id int64) error {
	err := repository.DB.WithContext(ctx).Delete(&entity.CategorySkill{}, id).Error
	if err != nil {
		log.Println("[CategorySkillRepository][Delete] problem querying to db, err: ", err.Error())
		return err
	}

	return nil
}
