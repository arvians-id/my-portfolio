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
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

func NewInitializedRoutes(configuration config.Config, logFile *os.File, db *gorm.DB) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			path := ctx.Locals("middleware")
			if path == nil {
				path = "getting an error from fiber's middleware"
			}

			path = fmt.Sprintf("getting an error from fiber's middleware: %s", path)

			gqlError := &gqlerror.Error{
				Message: err.Error(),
				Extensions: map[string]interface{}{
					"path": path,
					"code": code,
				},
			}

			return ctx.Status(http.StatusOK).JSON(gqlError)
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
				if logFile != nil {
					oc := graphql.GetOperationContext(ctx)
					_, err = logFile.WriteString(fmt.Sprintf("[%s] | query history: %s %s\n",
						time.Now().Format(time.RFC822),
						oc.OperationName,
						oc.RawQuery,
					))
					if err != nil {
						panic(err)
					}
				}

				return next(ctx)
			})
			h.ServeHTTP(writer, request)
		})(c.Context())

		return nil
	})

	return app, nil
}
