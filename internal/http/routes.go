package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/arvians-id/go-portfolio/cmd/config"
	"github.com/arvians-id/go-portfolio/internal/http/controller"
	"github.com/arvians-id/go-portfolio/internal/http/controller/resolver"
	"github.com/arvians-id/go-portfolio/internal/http/middleware"
	"github.com/arvians-id/go-portfolio/internal/http/response"
	"github.com/arvians-id/go-portfolio/internal/repository"
	"github.com/arvians-id/go-portfolio/internal/service"
	"github.com/blevesearch/bleve"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"net/http"
	"os"
	"time"
)

func NewInitializedRoutes(configuration config.Config, logFile *os.File) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return response.ReturnError(ctx, code, err)
		},
	})
	app.Use(etag.New())
	app.Use(requestid.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-API-KEY",
		AllowMethods:     "POST, DELETE, PUT, PATCH, GET",
		AllowCredentials: true,
	}))
	app.Use(middleware.ExposeFiberContext())
	app.Use(middleware.XApiKey(configuration))

	db, err := config.NewPostgresSQLGorm(configuration)
	if err != nil {
		return nil, err
	}

	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("./database/bleve", mapping)
	if err != nil {
		if err != bleve.ErrorIndexPathExists {
			return nil, err
		}
		index, err = bleve.Open("./database/bleve")
		if err != nil {
			return nil, err
		}
	}

	bleveSearchService := service.NewBleveSearchService(index)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	contactRepository := repository.NewContactRepository(db)
	contactService := service.NewContactService(contactRepository)

	educationRepository := repository.NewEducationRepository(db)
	educationService := service.NewEducationService(educationRepository)

	certificateRepository := repository.NewCertificateRepository(db)
	certificateService := service.NewCertificateService(certificateRepository)

	projectImageRepository := repository.NewProjectImageRepository(db)
	projectImageService := service.NewProjectImageService(projectImageRepository)

	projectRepository := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepository, projectImageRepository, bleveSearchService)

	categorySkillRepository := repository.NewCategorySkillRepository(db)
	categorySkillService := service.NewCategorySkillService(categorySkillRepository)

	workExperienceRepository := repository.NewWorkExperienceRepository(db)
	workExperienceService := service.NewWorkExperienceService(workExperienceRepository)

	skillRepository := repository.NewSkillRepository(db)
	skillService := service.NewSkillService(skillRepository)

	// Set GraphQL Playground
	app.Get("/", func(c *fiber.Ctx) error {
		h := playground.Handler("GraphQL", "/query")
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			h.ServeHTTP(writer, request)
		})(c.Context())

		return nil
	})

	app.Use(middleware.DataLoaders(skillService, categorySkillService, projectImageService))
	app.Post("/query", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			generatedConfig := gql.Config{
				Resolvers: &resolver.Resolver{
					UserService:           userService,
					ContactService:        contactService,
					EducationService:      educationService,
					CertificateService:    certificateService,
					ProjectService:        projectService,
					CategorySkillService:  categorySkillService,
					WorkExperienceService: workExperienceService,
					SkillService:          skillService,
					ProjectImagService:    projectImageService,
				},
			}
			generatedConfig.Directives.IsLoggedIn = middleware.NewJWTMiddlewareGraphQL

			h := handler.NewDefaultServer(gql.NewExecutableSchema(generatedConfig))
			h.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
				oc := graphql.GetOperationContext(ctx)
				_, err := logFile.WriteString(fmt.Sprintf("[%s] | query history: %s %s\n",
					time.Now().Format(time.RFC822),
					oc.OperationName,
					oc.RawQuery,
				))
				if err != nil {
					panic(err)
				}

				return next(ctx)
			})
			h.ServeHTTP(writer, request)
		})(c.Context())

		return nil
	})

	return app, nil
}