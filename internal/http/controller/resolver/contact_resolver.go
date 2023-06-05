package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
)

func (q queryResolver) FindAllContact(ctx context.Context) ([]*model.Contact, error) {
	contacts, err := q.ContactService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func (q queryResolver) FindByIDContact(ctx context.Context, id int64) (*model.Contact, error) {
	contact, err := q.ContactService.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (m mutationResolver) CreateContact(ctx context.Context, input model.CreateContactRequest) (*model.Contact, error) {
	contact, err := m.ContactService.Create(ctx, &input)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (m mutationResolver) UpdateContact(ctx context.Context, input model.UpdateContactRequest) (*model.Contact, error) {
	contact, err := m.ContactService.Update(ctx, &input)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (m mutationResolver) DeleteContact(ctx context.Context, id int64) (bool, error) {
	err := m.ContactService.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
