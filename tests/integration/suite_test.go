package integration

import (
	"github.com/arvians-id/go-portfolio/cmd/config"
	httpfiber "github.com/arvians-id/go-portfolio/internal/http"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"testing"
)

type Error struct {
	Message    string          `json:"message"`
	Path       []string        `json:"path"`
	Extensions ErrorExtensions `json:"extensions"`
}

type ErrorExtensions struct {
	Code string `json:"code"`
}

type E2ETestSuite struct {
	suite.Suite
	configuration config.Config
	server        *fiber.App
	db            *gorm.DB
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, new(E2ETestSuite))
}

func (suite *E2ETestSuite) SetupSuite() {
	// Init Config
	configuration := config.New("../../.env")

	// Init DB
	db, err := config.NewPostgresSQLGormTest(configuration)
	suite.Require().NoError(err)

	// Init Server
	server, err := httpfiber.NewInitializedRoutes(configuration, nil, db)
	suite.Require().NoError(err)

	// Set Suite Variables
	suite.configuration = configuration
	suite.server = server
	suite.db = db
}

func (suite *E2ETestSuite) TearDownSuite() {
	folderPath := "./assets/images"

	folderList, err := filepath.Glob(filepath.Join(folderPath, "*"))
	suite.Require().NoError(err)

	for _, folder := range folderList {
		fileList, err := filepath.Glob(filepath.Join(folder, "*"))
		suite.Require().NoError(err)

		for _, file := range fileList {
			_ = os.Remove(file)
		}
	}
}

func (suite *E2ETestSuite) TearDownTest() {
	suite.Require().NoError(config.TearDownDBTest(suite.db))
}
