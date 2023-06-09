package model

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/service"
	"time"
)

func FindAllSkillsByProjectIDs(ctx context.Context, skillService service.SkillServiceContract) ProjectSkillsLoader {
	return ProjectSkillsLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []int64) ([][]*Skill, []error) {
			skills, err := skillService.FindAllByProjectIDs(ctx, keys)
			if err != nil {
				return nil, []error{err}
			}

			u := make(map[int64][]*Skill, len(skills))
			for _, skill := range skills {
				u[skill.ProjectID] = append(u[skill.ProjectID], &Skill{
					ID:              skill.ID,
					CategorySkillID: skill.CategorySkillID,
					Name:            skill.Name,
					Icon:            skill.Icon,
				})
			}

			result := make([][]*Skill, len(keys))

			for i, key := range keys {
				result[i] = u[key]
			}

			return result, nil
		},
	}
}

func FindAllSkillsByWorkExperienceIDs(ctx context.Context, skillService service.SkillServiceContract) WorkExperienceSkillsLoader {
	return WorkExperienceSkillsLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []int64) ([][]*Skill, []error) {
			skills, err := skillService.FindAllByWorkExperienceIDs(ctx, keys)
			if err != nil {
				return nil, []error{err}
			}

			u := make(map[int64][]*Skill, len(skills))
			for _, skill := range skills {
				u[skill.WorkExperienceID] = append(u[skill.WorkExperienceID], &Skill{
					ID:              skill.ID,
					CategorySkillID: skill.CategorySkillID,
					Name:            skill.Name,
					Icon:            skill.Icon,
				})
			}

			result := make([][]*Skill, len(keys))

			for i, key := range keys {
				result[i] = u[key]
			}

			return result, nil
		},
	}
}

func FindAllSkillsByCategoryIDs(ctx context.Context, skillService service.SkillServiceContract) CategorySkillsLoader {
	return CategorySkillsLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []int64) ([][]*Skill, []error) {
			skills, err := skillService.FindAllByCategorySkillIDs(ctx, keys)
			if err != nil {
				return nil, []error{err}
			}

			u := make(map[int64][]*Skill, len(skills))
			for _, skill := range skills {
				u[skill.CategorySkillID] = append(u[skill.CategorySkillID], &Skill{
					ID:              skill.ID,
					CategorySkillID: skill.CategorySkillID,
					Name:            skill.Name,
					Icon:            skill.Icon,
				})
			}

			result := make([][]*Skill, len(keys))

			for i, key := range keys {
				result[i] = u[key]
			}

			return result, nil
		},
	}
}

func FindCategoryBySkillIDs(ctx context.Context, categoryService service.CategorySkillServiceContract) SkillsCategoryLoader {
	return SkillsCategoryLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []int64) ([]*CategorySkill, []error) {
			categories, err := categoryService.FindAllByIDs(ctx, keys)
			if err != nil {
				return nil, []error{err}
			}

			u := make(map[int64]*CategorySkill, len(categories))
			for _, category := range categories {
				u[category.ID] = &CategorySkill{
					ID:        category.ID,
					Name:      category.Name,
					CreatedAt: category.CreatedAt.String(),
					UpdatedAt: category.UpdatedAt.String(),
				}
			}

			result := make([]*CategorySkill, len(keys))

			for i, key := range keys {
				result[i] = u[key]
			}

			return result, nil
		},
	}
}

func FindAllImagesByProjectIDs(ctx context.Context, projectService service.ProjectServiceContract) ProjectImagesLoader {
	return ProjectImagesLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []int64) ([][]*ProjectImage, []error) {
			images, err := projectService.FindAllImagesByIDs(ctx, keys)
			if err != nil {
				return nil, []error{err}
			}

			u := make(map[int64][]*ProjectImage, len(images))
			for _, image := range images {
				u[image.ProjectID] = append(u[image.ProjectID], &ProjectImage{
					ID:        image.ID,
					ProjectID: image.ProjectID,
					Image:     image.Image,
				})
			}

			result := make([][]*ProjectImage, len(keys))

			for i, key := range keys {
				result[i] = u[key]
			}

			return result, nil
		},
	}
}
