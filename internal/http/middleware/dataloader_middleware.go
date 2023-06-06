package middleware

import (
	"context"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/internal/service"
	"github.com/gofiber/fiber/v2"
)

type Loaders struct {
	ListSkillsByProjectIDs        model.ProjectSkillsLoader
	ListSkillsByWorkExperienceIDs model.WorkExperienceSkillsLoader
	ListSkillsByCategoryIDs       model.CategorySkillsLoader
	ListCategoryBySkillIDs        model.SkillsCategoryLoader
}

func DataLoaders(
	skillService service.SkillServiceContract,
	categorySkillService service.CategorySkillServiceContract,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		loaders := Loaders{
			ListSkillsByProjectIDs:        model.FindAllSkillsByProjectIDs(c.Context(), skillService),
			ListSkillsByWorkExperienceIDs: model.FindAllSkillsByWorkExperienceIDs(c.Context(), skillService),
			ListSkillsByCategoryIDs:       model.FindAllSkillsByCategoryIDs(c.Context(), skillService),
			ListCategoryBySkillIDs:        model.FindCategoryBySkillIDs(c.Context(), categorySkillService),
		}

		c.Locals("loaders", &loaders)
		return c.Next()
	}
}

func GetLoaders(ctx context.Context) *Loaders {
	return ctx.Value("loaders").(*Loaders)
}
