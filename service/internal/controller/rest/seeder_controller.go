package rest

import (
	"database/sql"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/nibroos/e-proc-api/service/internal/utils"
)

type SeederController struct {
	db *sql.DB
}

func NewSeederController(db *sql.DB) *SeederController {
	return &SeederController{db: db}
}

func (c *SeederController) RunSeeders(ctx *fiber.Ctx) error {
	// List of seed files to execute

	seedFiles := []string{
		"20240916203801_create_users_seeder.sql",
		"20240916203915_create_groups_seeder.sql",
		"20240916203802_create_roles_values_seeder.sql",
		"20240916203953_create_permissions_values_seeder.sql",
		"20240916204023_create_user_roles_seeder.sql",
		"20240916204053_create_role_permissions_seeder.sql",
		"20241105045641_create_mix_values_identifier_seeder.sql",
		"20241105045650_create_mix_values_contact_seeder.sql",
		"20241105045700_create_mix_values_address_seeder.sql",
		"20250206045649_create_mix_values_customer_types_seeder.sql",
		"20250206045650_create_mix_values_item_groups_seeder.sql",
		"20250206045651_create_mix_values_item_sub_groups_seeder.sql",
		"20250206045652_create_mix_values_item_units_seeder.sql",
		"20250206045653_create_mix_values_currencies_seeder.sql",
		"20250206045655_create_customers_seeder.sql",
		"20250206045657_create_items_seeder.sql",
		"20250206045658_create_catalogs_seeder.sql",
		"20250206045659_create_catalog_details_seeder.sql",
	}

	// Get the seed files directory from the environment variable
	seedDir := os.Getenv("SEEDER_DIR")
	if seedDir == "" {
		seedDir = "internal/database/seeders" // Default directory
	}

	// Prepend the directory path to each seed file
	for i, file := range seedFiles {
		seedFiles[i] = filepath.Join(seedDir, file)
	}

	err := utils.ExecuteSeeders(c.db, seedFiles)
	if err != nil {
		return utils.JSONError(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(fiber.Map{
		"message": "Seeders executed successfully",
	})
}
