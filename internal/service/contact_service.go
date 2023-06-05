package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"log"
)

type ContactServiceContract interface {
	FindAll(ctx context.Context) ([]*model.Contact, error)
	FindByID(ctx context.Context, id int64) (*model.Contact, error)
	Create(ctx context.Context, request *model.CreateContactRequest) (*model.Contact, error)
	Update(ctx context.Context, request *model.UpdateContactRequest) (*model.Contact, error)
	Delete(ctx context.Context, id int64) error
}

type ContactService struct {
	ContactRepository repository.ContactRepositoryContract
}

func NewContactService(ContactRepository repository.ContactRepositoryContract) ContactServiceContract {
	return &ContactService{
		ContactRepository: ContactRepository,
	}
}

func (service *ContactService) FindAll(ctx context.Context) ([]*model.Contact, error) {
	contacts, err := service.ContactRepository.FindAll(ctx)
	if err != nil {
		log.Println("[ContactService][FindAll] problem calling repository, err: ", err.Error())
		return nil, err
	}

	var results []*model.Contact
	for _, contact := range contacts {
		results = append(results, contact.ToModel())
	}

	return results, nil
}

func (service *ContactService) FindByID(ctx context.Context, id int64) (*model.Contact, error) {
	contact, err := service.ContactRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[ContactService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return contact.ToModel(), nil
}

func (service *ContactService) Create(ctx context.Context, request *model.CreateContactRequest) (*model.Contact, error) {
	contactCreated, err := service.ContactRepository.Create(ctx, request)
	if err != nil {
		log.Println("[ContactService][Create] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return contactCreated.ToModel(), nil
}

func (service *ContactService) Update(ctx context.Context, request *model.UpdateContactRequest) (*model.Contact, error) {
	_, err := service.ContactRepository.FindByID(ctx, request.ID)
	if err != nil {
		log.Println("[ContactService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	err = service.ContactRepository.Update(ctx, request)
	if err != nil {
		log.Println("[ContactService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	contactUpdated, err := service.ContactRepository.FindByID(ctx, request.ID)
	if err != nil {
		log.Println("[ContactService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return contactUpdated.ToModel(), nil
}

func (service *ContactService) Delete(ctx context.Context, id int64) error {
	contactCheck, err := service.ContactRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[ContactService][FindByID] problem calling repository, err: ", err.Error())
		return err
	}

	err = service.ContactRepository.Delete(ctx, contactCheck.ID)
	if err != nil {
		log.Println("[ContactService][Delete] problem calling repository, err: ", err.Error())
		return err
	}

	return nil
}
