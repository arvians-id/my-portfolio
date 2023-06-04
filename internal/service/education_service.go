package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"log"
)

type EducationServiceContract interface {
	FindAll(ctx context.Context) ([]*model.Education, error)
	FindById(ctx context.Context, id int64) (*model.Education, error)
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

	var educationsModel []*model.Education
	for _, education := range educations {
		educationsModel = append(educationsModel, (*model.Education)(education))
	}

	return educationsModel, nil
}

func (service *EducationService) FindById(ctx context.Context, id int64) (*model.Education, error) {
	education, err := service.EducationRepository.FindById(ctx, id)
	if err != nil {
		log.Println("[EducationService][FindById] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return (*model.Education)(education), nil
}

func (service *EducationService) Create(ctx context.Context, request *model.CreateEducationRequest) (*model.Education, error) {
	var education entity.Education
	education.Institution = request.Institution
	education.Degree = request.Degree
	education.FieldOfStudy = request.FieldOfStudy
	education.Grade = request.Grade
	education.Description = request.Description
	education.StartDate = request.StartDate
	education.EndDate = request.EndDate

	educationCreated, err := service.EducationRepository.Create(ctx, &education)
	if err != nil {
		log.Println("[EducationService][Create] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return (*model.Education)(educationCreated), nil
}

func (service *EducationService) Update(ctx context.Context, request *model.UpdateEducationRequest) (*model.Education, error) {
	educationCheck, err := service.EducationRepository.FindById(ctx, request.ID)
	if err != nil {
		log.Println("[EducationService][FindById] problem calling repository, err: ", err.Error())
		return nil, err
	}

	educationCheck.ID = request.ID
	educationCheck.Institution = *request.Institution
	educationCheck.Degree = *request.Degree
	educationCheck.FieldOfStudy = *request.FieldOfStudy
	educationCheck.Grade = *request.Grade
	educationCheck.Description = request.Description
	educationCheck.StartDate = *request.StartDate
	educationCheck.EndDate = request.EndDate

	err = service.EducationRepository.Update(ctx, educationCheck)
	if err != nil {
		log.Println("[EducationService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return (*model.Education)(educationCheck), nil
}

func (service *EducationService) Delete(ctx context.Context, id int64) error {
	educationCheck, err := service.EducationRepository.FindById(ctx, id)
	if err != nil {
		log.Println("[EducationService][FindById] problem calling repository, err: ", err.Error())
		return err
	}

	err = service.EducationRepository.Delete(ctx, educationCheck.ID)
	if err != nil {
		log.Println("[EducationService][Delete] problem calling repository, err: ", err.Error())
		return err
	}

	return nil
}
