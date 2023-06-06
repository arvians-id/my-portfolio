package repository

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"gorm.io/gorm"
	"log"
)

type SkillRepositoryContract interface {
	FindAll(ctx context.Context) ([]*entity.Skill, error)
	FindAllByProjectIDs(ctx context.Context, projectIDs []int64) ([]*entity.SkillBelongsTo, error)
	FindAllByWorkExperienceIDs(ctx context.Context, workExperienceIDs []int64) ([]*entity.SkillBelongsTo, error)
	FindAllByCategorySkillIDs(ctx context.Context, categorySkillIDs []int64) ([]*entity.Skill, error)
	FindByID(ctx context.Context, id int64) (*entity.Skill, error)
	Create(ctx context.Context, skill *entity.Skill) (*entity.Skill, error)
	Update(ctx context.Context, skill *entity.Skill) error
	Delete(ctx context.Context, id int64) error
}

type SkillRepository struct {
	DB *gorm.DB
}

func NewSkillRepository(db *gorm.DB) SkillRepositoryContract {
	return &SkillRepository{DB: db}
}

func (repository *SkillRepository) FindAll(ctx context.Context) ([]*entity.Skill, error) {
	query := "SELECT * FROM skills"
	var skills []*entity.Skill
	err := repository.DB.WithContext(ctx).Raw(query).Scan(&skills).Error
	if err != nil {
		log.Println("[SkillRepository][FindAll] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return skills, nil
}

func (repository *SkillRepository) FindAllByProjectIDs(ctx context.Context, projectIDs []int64) ([]*entity.SkillBelongsTo, error) {
	query := "SELECT s.*, ps.project_id  FROM skills s LEFT JOIN project_skill ps ON ps.skill_id = s.id WHERE ps.project_id IN (?)"
	var skills []*entity.SkillBelongsTo
	err := repository.DB.WithContext(ctx).Raw(query, projectIDs).Scan(&skills).Error
	if err != nil {
		log.Println("[SkillRepository][FindAllByProjectIDs] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return skills, nil
}

func (repository *SkillRepository) FindAllByWorkExperienceIDs(ctx context.Context, workExperienceIDs []int64) ([]*entity.SkillBelongsTo, error) {
	query := "SELECT s.*, ws.work_experience_id  FROM skills s LEFT JOIN work_experience_skill ws ON ws.skill_id = s.id WHERE ws.work_experience_id IN (?)"
	var skills []*entity.SkillBelongsTo
	err := repository.DB.WithContext(ctx).Raw(query, workExperienceIDs).Scan(&skills).Error
	if err != nil {
		log.Println("[SkillRepository][FindAllByWorkExperienceIDs] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return skills, nil
}

func (repository *SkillRepository) FindAllByCategorySkillIDs(ctx context.Context, categorySkillIDs []int64) ([]*entity.Skill, error) {
	query := "SELECT * FROM skills WHERE category_skill_id IN (?)"
	var skills []*entity.Skill
	err := repository.DB.WithContext(ctx).Raw(query, categorySkillIDs).Scan(&skills).Error
	if err != nil {
		log.Println("[SkillRepository][FindAllByCategorySkillIDs] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return skills, nil
}

func (repository *SkillRepository) FindByID(ctx context.Context, id int64) (*entity.Skill, error) {
	query := "SELECT * FROM skills WHERE id = ?"
	var skill entity.Skill
	row := repository.DB.WithContext(ctx).Raw(query, id).Row()
	err := row.Scan(&skill.ID, &skill.CategorySkillID, &skill.Name, &skill.Icon)
	if err != nil {
		log.Println("[SkillRepository][FindByID] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return &skill, nil
}

func (repository *SkillRepository) Create(ctx context.Context, skill *entity.Skill) (*entity.Skill, error) {
	err := repository.DB.WithContext(ctx).Create(&skill).Error
	if err != nil {
		log.Println("[SkillRepository][Create] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return skill, nil
}

func (repository *SkillRepository) Update(ctx context.Context, skill *entity.Skill) error {
	err := repository.DB.WithContext(ctx).Updates(&skill).Error
	if err != nil {
		log.Println("[SkillRepository][Updates] problem querying to db, err: ", err.Error())
		return err
	}

	return nil
}

func (repository *SkillRepository) Delete(ctx context.Context, id int64) error {
	err := repository.DB.WithContext(ctx).Delete(&entity.Skill{}, id).Error
	if err != nil {
		log.Println("[SkillRepository][Delete] problem querying to db, err: ", err.Error())
		return err
	}

	return nil
}
