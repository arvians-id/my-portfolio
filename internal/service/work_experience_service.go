package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"log"
)

type WorkExperienceServiceContract interface {
	FindAll(ctx context.Context) ([]*entity.WorkExperience, error)
	FindByID(ctx context.Context, id int64) (*entity.WorkExperience, error)
	Create(ctx context.Context, workExperience *entity.WorkExperience) (*entity.WorkExperience, error)
	Update(ctx context.Context, workExperience *entity.WorkExperience) (*entity.WorkExperience, error)
	Delete(ctx context.Context, id int64) error
}

type WorkExperienceService struct {
	WorkExperienceRepository repository.WorkExperienceRepositoryContract
}

func NewWorkExperienceService(workExperienceRepository repository.WorkExperienceRepositoryContract) WorkExperienceServiceContract {
	return &WorkExperienceService{
		WorkExperienceRepository: workExperienceRepository,
	}
}
func (service *WorkExperienceService) FindAll(ctx context.Context) ([]*entity.WorkExperience, error) {
	workExperiences, err := service.WorkExperienceRepository.FindAll(ctx)
	if err != nil {
		log.Println("[WorkExperienceService][FindAll] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return workExperiences, nil
}

func (service *WorkExperienceService) FindByID(ctx context.Context, id int64) (*entity.WorkExperience, error) {
	workExperience, err := service.WorkExperienceRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[WorkExperienceService][FindById] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return workExperience, nil
}

func (service *WorkExperienceService) Create(ctx context.Context, workExperience *entity.WorkExperience) (*entity.WorkExperience, error) {
	workExperience, err := service.WorkExperienceRepository.Create(ctx, workExperience)
	if err != nil {
		log.Println("[WorkExperienceService][Create] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return workExperience, nil
}

func (service *WorkExperienceService) Update(ctx context.Context, workExperience *entity.WorkExperience) (*entity.WorkExperience, error) {
	_, err := service.WorkExperienceRepository.FindByID(ctx, workExperience.ID)
	if err != nil {
		log.Println("[WorkExperienceService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	err = service.WorkExperienceRepository.Update(ctx, workExperience)
	if err != nil {
		log.Println("[WorkExperienceService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	workExperienceUpdated, err := service.WorkExperienceRepository.FindByID(ctx, workExperience.ID)
	if err != nil {
		log.Println("[WorkExperienceService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return workExperienceUpdated, nil
}

func (service *WorkExperienceService) Delete(ctx context.Context, id int64) error {
	workExperienceCheck, err := service.WorkExperienceRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[WorkExperienceService][FindByID] problem calling repository, err: ", err.Error())
		return err
	}

	err = service.WorkExperienceRepository.Delete(ctx, workExperienceCheck.ID)
	if err != nil {
		log.Println("[WorkExperienceService][Delete] problem calling repository, err: ", err.Error())
		return err
	}

	return nil
}
