package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"log"
)

type CategorySkillServiceContract interface {
	FindAll(ctx context.Context) ([]*model.CategorySkill, error)
	FindAllByIDs(ctx context.Context, ids []int64) ([]*model.CategorySkill, error)
	FindByID(ctx context.Context, id int64) (*model.CategorySkill, error)
	Create(ctx context.Context, request *model.CreateCategorySkillRequest) (*model.CategorySkill, error)
	Update(ctx context.Context, request *model.UpdateCategorySkillRequest) (*model.CategorySkill, error)
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

func (service *CategorySkillService) FindAll(ctx context.Context) ([]*model.CategorySkill, error) {
	categories, err := service.CategorySkillRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var results []*model.CategorySkill
	for _, category := range categories {
		results = append(results, category.ToModel())
	}

	return results, nil
}

func (service *CategorySkillService) FindAllByIDs(ctx context.Context, ids []int64) ([]*model.CategorySkill, error) {
	categories, err := service.CategorySkillRepository.FindAllByIDs(ctx, ids)
	if err != nil {
		log.Println("[CategorySkillService][FindAll] problem calling repository, err: ", err.Error())
		return nil, err
	}

	var results []*model.CategorySkill
	for _, category := range categories {
		results = append(results, category.ToModel())
	}

	return results, nil
}

func (service *CategorySkillService) FindByID(ctx context.Context, id int64) (*model.CategorySkill, error) {
	category, err := service.CategorySkillRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[CategorySkillService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return category.ToModel(), nil
}

func (service *CategorySkillService) Create(ctx context.Context, request *model.CreateCategorySkillRequest) (*model.CategorySkill, error) {
	category, err := service.CategorySkillRepository.Create(ctx, request)
	if err != nil {
		log.Println("[CategorySkillService][Create] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return category.ToModel(), nil
}

func (service *CategorySkillService) Update(ctx context.Context, request *model.UpdateCategorySkillRequest) (*model.CategorySkill, error) {
	categoryCheck, err := service.CategorySkillRepository.FindByID(ctx, request.ID)
	if err != nil {
		log.Println("[CategorySkillService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	err = service.CategorySkillRepository.Update(ctx, request)
	if err != nil {
		log.Println("[CategorySkillService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return categoryCheck.ToModel(), nil

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
