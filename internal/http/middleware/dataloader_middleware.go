package middleware

import (
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/internal/service"
	"github.com/gofiber/fiber/v2"
)

type Loaders struct {
	ListSkillsByProjectIDs        model.ProjectSkillsLoader
	ListSkillsByWorkExperienceIDs model.WorkExperienceSkillsLoader
	ListSkillsByCategoryIDs       model.CategorySkillsLoader
	ListCategoryBySkillIDs        model.SkillsCategoryLoader
	ListImagesByProjectIDs        model.ProjectImagesLoader
}

func DataLoaders(
	skillService service.SkillServiceContract,
	categorySkillService service.CategorySkillServiceContract,
	projectImageService service.ProjectImageServiceContract,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		loaders := Loaders{
			ListSkillsByProjectIDs:        model.FindAllSkillsByProjectIDs(c.Context(), skillService),
			ListSkillsByWorkExperienceIDs: model.FindAllSkillsByWorkExperienceIDs(c.Context(), skillService),
			ListSkillsByCategoryIDs:       model.FindAllSkillsByCategoryIDs(c.Context(), skillService),
			ListCategoryBySkillIDs:        model.FindCategoryBySkillIDs(c.Context(), categorySkillService),
			ListImagesByProjectIDs:        model.FindAllImagesByProjectIDs(c.Context(), projectImageService),
		}

		c.Locals("loaders", &loaders)
		return c.Next()
	}
}
