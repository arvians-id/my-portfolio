package service

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"github.com/arvians-id/go-portfolio/util"
	"log"
)

type ContactServiceContract interface {
	FindAll(ctx context.Context) ([]*entity.Contact, error)
	FindByID(ctx context.Context, id int64) (*entity.Contact, error)
	Create(ctx context.Context, contact *entity.Contact) (*entity.Contact, error)
	Update(ctx context.Context, contact *entity.Contact) (*entity.Contact, error)
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

func (service *ContactService) FindAll(ctx context.Context) ([]*entity.Contact, error) {
	contacts, err := service.ContactRepository.FindAll(ctx)
	if err != nil {
		log.Println("[ContactService][FindAll] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return contacts, nil
}

func (service *ContactService) FindByID(ctx context.Context, id int64) (*entity.Contact, error) {
	contact, err := service.ContactRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("[ContactService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return contact, nil
}

func (service *ContactService) Create(ctx context.Context, contact *entity.Contact) (*entity.Contact, error) {
	contactCreated, err := service.ContactRepository.Create(ctx, contact)
	if err != nil {
		log.Println("[ContactService][Create] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return contactCreated, nil
}

func (service *ContactService) Update(ctx context.Context, contact *entity.Contact) (*entity.Contact, error) {
	contactCheck, err := service.ContactRepository.FindByID(ctx, contact.ID)
	if err != nil {
		log.Println("[ContactService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	err = service.ContactRepository.Update(ctx, contact)
	if err != nil {
		log.Println("[ContactService][Update] problem calling repository, err: ", err.Error())
		return nil, err
	}

	if contactCheck.Icon != contact.Icon && contactCheck.Icon != nil {
		path := "images/contact"
		err = util.DeleteFile(path, *contactCheck.Icon)
		if err != nil {
			return nil, err
		}
	}

	contactUpdated, err := service.ContactRepository.FindByID(ctx, contact.ID)
	if err != nil {
		log.Println("[ContactService][FindByID] problem calling repository, err: ", err.Error())
		return nil, err
	}

	return contactUpdated, nil
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

	if contactCheck.Icon != nil {
		path := "images/contact"
		err = util.DeleteFile(path, *contactCheck.Icon)
		if err != nil {
			return err
		}
	}

	return nil
}
