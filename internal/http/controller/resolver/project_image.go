package resolver

import "context"

func (m mutationResolver) DeleteProjectImage(ctx context.Context, id int64) (bool, error) {
	err := m.ProjectImagService.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
