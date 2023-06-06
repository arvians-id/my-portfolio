package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"log"
)

type EducationServiceContract interface {
	FindAll(ctx context.Context) ([]*entity.Education, error)
	FindByID(ctx context.Context, id int64) (*entity.Education, error)
	Create(ctx context.Context, education *entity.Education) (*entity.Education, error)
	Update(ctx context.Context, education *entity.Education) (*entity.Education, error)
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

func (service *EducationService) FindAll(ctx context.Context) ([]*entity.Education, error) {
	educations, err := service.EducationRepository.FindAll(ctx)
	if err != nil {
		log.Println("[EducationService][FindAll] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return educations, nil
}

func (service *EducationService) FindByID(ctx context.Context, id int64) (*entity.Education, error) {
	education, err := service.EducationRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[EducationService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return education, nil
}

func (service *EducationService) Create(ctx context.Context, education *entity.Education) (*entity.Education, error) {
	educationCreated, err := service.EducationRepository.Create(ctx, education)
	if err != nil {
		log.Println("[EducationService][Create] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return educationCreated, nil
}

func (service *EducationService) Update(ctx context.Context, education *entity.Education) (*entity.Education, error) {
	_, err := service.EducationRepository.FindByID(ctx, education.ID)
	if err != nil {
		log.Println("[EducationService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	err = service.EducationRepository.Update(ctx, education)
	if err != nil {
		log.Println("[EducationService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	educationUpdated, err := service.EducationRepository.FindByID(ctx, education.ID)
	if err != nil {
		log.Println("[EducationService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return educationUpdated, nil
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
