package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"log"
)

type WorkExperienceServiceContract interface {
	FindAll(ctx context.Context) ([]*model.WorkExperience, error)
	FindByID(ctx context.Context, id int64) (*model.WorkExperience, error)
	Create(ctx context.Context, request *model.CreateWorkExperienceRequest) (*model.WorkExperience, error)
	Update(ctx context.Context, request *model.UpdateWorkExperienceRequest) (*model.WorkExperience, error)
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
func (service *WorkExperienceService) FindAll(ctx context.Context) ([]*model.WorkExperience, error) {
	workExperiences, err := service.WorkExperienceRepository.FindAll(ctx)
	if err != nil {
		log.Println("[WorkExperienceService][FindAll] problem calling repository, err: ", err.Error())
		return nil, err
	}

	var results []*model.WorkExperience
	for _, workExperience := range workExperiences {
		results = append(results, workExperience.ToModel())
	}

	return results, nil
}

func (service *WorkExperienceService) FindByID(ctx context.Context, id int64) (*model.WorkExperience, error) {
	workExperience, err := service.WorkExperienceRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[WorkExperienceService][FindById] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return workExperience.ToModel(), nil
}

func (service *WorkExperienceService) Create(ctx context.Context, request *model.CreateWorkExperienceRequest) (*model.WorkExperience, error) {
	workExperience, err := service.WorkExperienceRepository.Create(ctx, request)
	if err != nil {
		log.Println("[WorkExperienceService][Create] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return workExperience.ToModel(), nil
}

func (service *WorkExperienceService) Update(ctx context.Context, request *model.UpdateWorkExperienceRequest) (*model.WorkExperience, error) {
	_, err := service.WorkExperienceRepository.FindByID(ctx, request.ID)
	if err != nil {
		log.Println("[WorkExperienceService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	err = service.WorkExperienceRepository.Update(ctx, request)
	if err != nil {
		log.Println("[WorkExperienceService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	workExperienceUpdated, err := service.WorkExperienceRepository.FindByID(ctx, request.ID)
	if err != nil {
		log.Println("[WorkExperienceService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return workExperienceUpdated.ToModel(), nil
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
