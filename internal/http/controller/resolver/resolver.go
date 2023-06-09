package resolver

import (
	"context"
	gql "github.com/arvians-id/go-portfolio/internal/http/controller"
	"github.com/arvians-id/go-portfolio/internal/http/middleware"
	"github.com/arvians-id/go-portfolio/internal/service"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	UserService           service.UserServiceContract
	EducationService      service.EducationServiceContract
	CertificateService    service.CertificateServiceContract
	ContactService        service.ContactServiceContract
	ProjectService        service.ProjectServiceContract
	CategorySkillService  service.CategorySkillServiceContract
	WorkExperienceService service.WorkExperienceServiceContract
	SkillService          service.SkillServiceContract
	ProjectImagService    service.ProjectImageServiceContract
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) CategorySkill() gql.CategorySkillResolver {
	return &categorySkillResolver{r}
}

func (r *Resolver) Project() gql.ProjectResolver {
	return &projectResolver{r}
}

func (r *Resolver) Skill() gql.SkillResolver {
	return &skillResolver{r}
}

func (r *Resolver) WorkExperience() gql.WorkExperienceResolver {
	return &workExperienceResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
type categorySkillResolver struct{ *Resolver }
type projectResolver struct{ *Resolver }
type skillResolver struct{ *Resolver }
type workExperienceResolver struct{ *Resolver }

func GetLoaders(ctx context.Context) *middleware.Loaders {
	return ctx.Value("loaders").(*middleware.Loaders)
}
