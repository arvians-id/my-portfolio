package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"log"
)

type ContactServiceContract interface {
	FindAll(ctx context.Context) ([]*model.Contact, error)
	FindById(ctx context.Context, id int64) (*model.Contact, error)
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

	var contactsModel []*model.Contact
	for _, contact := range contacts {
		contactsModel = append(contactsModel, (*model.Contact)(contact))
	}

	return contactsModel, nil
}

func (service *ContactService) FindById(ctx context.Context, id int64) (*model.Contact, error) {
	contact, err := service.ContactRepository.FindById(ctx, id)
	if err != nil {
		log.Println("[ContactService][FindById] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return (*model.Contact)(contact), nil
}

func (service *ContactService) Create(ctx context.Context, request *model.CreateContactRequest) (*model.Contact, error) {
	var contact entity.Contact
	contact.Platform = request.Platform
	contact.URL = request.URL
	contact.Icon = request.Icon

	contactCreated, err := service.ContactRepository.Create(ctx, &contact)
	if err != nil {
		log.Println("[ContactService][Create] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return (*model.Contact)(contactCreated), nil
}

func (service *ContactService) Update(ctx context.Context, request *model.UpdateContactRequest) (*model.Contact, error) {
	contactCheck, err := service.ContactRepository.FindById(ctx, request.ID)
	if err != nil {
		log.Println("[ContactService][FindById] problem calling repository, err: ", err.Error())
		return nil, err
	}

	contactCheck.Platform = *request.Platform
	contactCheck.URL = *request.URL
	contactCheck.Icon = request.Icon

	err = service.ContactRepository.Update(ctx, contactCheck)
	if err != nil {
		log.Println("[ContactService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return (*model.Contact)(contactCheck), nil
}

func (service *ContactService) Delete(ctx context.Context, id int64) error {
	contactCheck, err := service.ContactRepository.FindById(ctx, id)
	if err != nil {
		log.Println("[ContactService][FindById] problem calling repository, err: ", err.Error())
		return err
	}

	err = service.ContactRepository.Delete(ctx, contactCheck.ID)
	if err != nil {
		log.Println("[ContactService][Delete] problem calling repository, err: ", err.Error())
		return err
	}

	return nil
}
