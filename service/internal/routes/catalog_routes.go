package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/nibroos/e-proc-api/service/internal/controller/rest"
	"github.com/nibroos/e-proc-api/service/internal/repository"
	"github.com/nibroos/e-proc-api/service/internal/service"
	"gorm.io/gorm"
)

func SetupCatalogtRoutes(catalogs fiber.Router, gormDB *gorm.DB, sqlDB *sqlx.DB) {
	catalogRepo := repository.NewCatalogRepository(gormDB, sqlDB)
	catalogService := service.NewCatalogService(catalogRepo)
	catalogController := rest.NewCatalogController(catalogService)

	// prefix /catalogs

	catalogs.Post("/index-catalog", catalogController.GetCatalogs)
	catalogs.Post("/show-catalog", catalogController.GetCatalogByID)
	catalogs.Post("/create-catalog", catalogController.CreateCatalog)
	catalogs.Post("/update-catalog", catalogController.UpdateCatalog)
	catalogs.Post("/delete-catalog", catalogController.DeleteCatalog)
	catalogs.Post("/restore-catalog", catalogController.RestoreCatalog)
}
