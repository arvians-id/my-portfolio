package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"log"
)

type CategorySkillServiceContract interface {
	FindAll(ctx context.Context) ([]*entity.CategorySkill, error)
	FindAllByIDs(ctx context.Context, ids []int64) ([]*entity.CategorySkill, error)
	FindByID(ctx context.Context, id int64) (*entity.CategorySkill, error)
	Create(ctx context.Context, categorySkill *entity.CategorySkill) (*entity.CategorySkill, error)
	Update(ctx context.Context, categorySkill *entity.CategorySkill) (*entity.CategorySkill, error)
	Delete(ctx context.Context, id int64) error
}

type CategorySkillService struct {
	CategorySkillRepository repository.CategorySkillRepositoryContract
}

func NewCategorySkillService(categorySkillRepository repository.CategorySkillRepositoryContract) CategorySkillServiceContract {
	return &CategorySkillService{
		CategorySkillRepository: categorySkillRepository,
	}
}

func (service *CategorySkillService) FindAll(ctx context.Context) ([]*entity.CategorySkill, error) {
	categories, err := service.CategorySkillRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (service *CategorySkillService) FindAllByIDs(ctx context.Context, ids []int64) ([]*entity.CategorySkill, error) {
	categories, err := service.CategorySkillRepository.FindAllByIDs(ctx, ids)
	if err != nil {
		log.Println("[CategorySkillService][FindAll] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return categories, nil
}

func (service *CategorySkillService) FindByID(ctx context.Context, id int64) (*entity.CategorySkill, error) {
	category, err := service.CategorySkillRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[CategorySkillService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return category, nil
}

func (service *CategorySkillService) Create(ctx context.Context, categorySkill *entity.CategorySkill) (*entity.CategorySkill, error) {
	category, err := service.CategorySkillRepository.Create(ctx, categorySkill)
	if err != nil {
		log.Println("[CategorySkillService][Create] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return category, nil
}

func (service *CategorySkillService) Update(ctx context.Context, categorySkill *entity.CategorySkill) (*entity.CategorySkill, error) {
	categoryCheck, err := service.CategorySkillRepository.FindByID(ctx, categorySkill.ID)
	if err != nil {
		log.Println("[CategorySkillService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	err = service.CategorySkillRepository.Update(ctx, categorySkill)
	if err != nil {
		log.Println("[CategorySkillService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return categoryCheck, nil

}

func (service *CategorySkillService) Delete(ctx context.Context, id int64) error {
	categoryCheck, err := service.CategorySkillRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[CategorySkillService][FindByID] problem calling repository, err: ", err.Error())
		return err
	}

	err = service.CategorySkillRepository.Delete(ctx, categoryCheck.ID)
	if err != nil {
		log.Println("[CategorySkillService][Delete] problem calling repository, err: ", err.Error())
		return err
	}

	return nil
}
