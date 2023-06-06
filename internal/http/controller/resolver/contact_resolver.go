package resolver

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/entity"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
)

func (q queryResolver) FindAllContact(ctx context.Context) ([]*model.Contact, error) {
	contacts, err := q.ContactService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var results []*model.Contact
	for _, contact := range contacts {
		results = append(results, &model.Contact{
			ID:       contact.ID,
			Platform: contact.Platform,
			URL:      contact.URL,
			Icon:     contact.Icon,
		})
	}

	return results, nil
}

func (q queryResolver) FindByIDContact(ctx context.Context, id int64) (*model.Contact, error) {
	contact, err := q.ContactService.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.Contact{
		ID:       contact.ID,
		Platform: contact.Platform,
		URL:      contact.URL,
		Icon:     contact.Icon,
	}, nil
}

func (m mutationResolver) CreateContact(ctx context.Context, input model.CreateContactRequest) (*model.Contact, error) {
	contact, err := m.ContactService.Create(ctx, &entity.Contact{
		Platform: input.Platform,
		URL:      input.URL,
		Icon:     input.Icon,
	})
	if err != nil {
		return nil, err
	}

	return &model.Contact{
		ID:       contact.ID,
		Platform: contact.Platform,
		URL:      contact.URL,
		Icon:     contact.Icon,
	}, nil
}

func (m mutationResolver) UpdateContact(ctx context.Context, input model.UpdateContactRequest) (*model.Contact, error) {
	contact, err := m.ContactService.Update(ctx, &entity.Contact{
		ID:       input.ID,
		Platform: *input.Platform,
		URL:      *input.URL,
		Icon:     input.Icon,
	})
	if err != nil {
		return nil, err
	}

	return &model.Contact{
		ID:       contact.ID,
		Platform: contact.Platform,
		URL:      contact.URL,
		Icon:     contact.Icon,
	}, nil
}

func (m mutationResolver) DeleteContact(ctx context.Context, id int64) (bool, error) {
	err := m.ContactService.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
