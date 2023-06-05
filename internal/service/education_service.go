package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"log"
)

type EducationServiceContract interface {
	FindAll(ctx context.Context) ([]*model.Education, error)
	FindByID(ctx context.Context, id int64) (*model.Education, error)
	Create(ctx context.Context, request *model.CreateEducationRequest) (*model.Education, error)
	Update(ctx context.Context, request *model.UpdateEducationRequest) (*model.Education, error)
	Delete(ctx context.Context, id int64) error
}

type EducationService struct {
	EducationRepository repository.EducationRepositoryContract
}

func NewEducationService(EducationRepository repository.EducationRepositoryContract) EducationServiceContract {
	return &EducationService{
		EducationRepository: EducationRepository,
	}
}

func (service *EducationService) FindAll(ctx context.Context) ([]*model.Education, error) {
	educations, err := service.EducationRepository.FindAll(ctx)
	if err != nil {
		log.Println("[EducationService][FindAll] problem calling repository, err: ", err.Error())
		return nil, err
	}

	var results []*model.Education
	for _, education := range educations {
		results = append(results, education.ToModel())
	}

	return results, nil
}

func (service *EducationService) FindByID(ctx context.Context, id int64) (*model.Education, error) {
	education, err := service.EducationRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[EducationService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return education.ToModel(), nil
}

func (service *EducationService) Create(ctx context.Context, request *model.CreateEducationRequest) (*model.Education, error) {
	educationCreated, err := service.EducationRepository.Create(ctx, request)
	if err != nil {
		log.Println("[EducationService][Create] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return educationCreated.ToModel(), nil
}

func (service *EducationService) Update(ctx context.Context, request *model.UpdateEducationRequest) (*model.Education, error) {
	_, err := service.EducationRepository.FindByID(ctx, request.ID)
	if err != nil {
		log.Println("[EducationService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	err = service.EducationRepository.Update(ctx, request)
	if err != nil {
		log.Println("[EducationService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	educationUpdated, err := service.EducationRepository.FindByID(ctx, request.ID)
	if err != nil {
		log.Println("[EducationService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return educationUpdated.ToModel(), nil
}

func (service *EducationService) Delete(ctx context.Context, id int64) error {
	educationCheck, err := service.EducationRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[EducationService][FindByID] problem calling repository, err: ", err.Error())
		return err
	}

	err = service.EducationRepository.Delete(ctx, educationCheck.ID)
	if err != nil {
		log.Println("[EducationService][Delete] problem calling repository, err: ", err.Error())
		return err
	}

	return nil
}
