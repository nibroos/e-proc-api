package rest

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nibroos/e-proc-api/service/internal/dtos"
	"github.com/nibroos/e-proc-api/service/internal/middleware"
	"github.com/nibroos/e-proc-api/service/internal/models"
	"github.com/nibroos/e-proc-api/service/internal/service"
	"github.com/nibroos/e-proc-api/service/internal/utils"
	"github.com/nibroos/e-proc-api/service/internal/validators/form_requests"
)

type CatalogController struct {
	service *service.CatalogService
}

func NewCatalogController(service *service.CatalogService) *CatalogController {
	return &CatalogController{service: service}
}

func (c *CatalogController) GetCatalogs(ctx *fiber.Ctx) error {
	filters, ok := ctx.Locals("filters").(map[string]string)
	if !ok {
		return utils.SendResponse(ctx, utils.WrapResponse(nil, nil, "Invalid filters", http.StatusBadRequest), http.StatusBadRequest)
	}

	catalogs, total, err := c.service.GetCatalogs(ctx.Context(), filters)
	if err != nil {
		return utils.SendResponse(ctx, utils.WrapResponse(nil, nil, err.Error(), http.StatusInternalServerError), http.StatusInternalServerError)
	}

	paginationMeta := utils.CreatePaginationMeta(filters, total)

	return utils.GetResponse(ctx, catalogs, paginationMeta, "Catalogs fetched successfully", http.StatusOK, nil, nil)
}
func (c *CatalogController) CreateCatalog(ctx *fiber.Ctx) error {
	var req dtos.CreateCatalogRequest

	// Use the utility function to parse the request body
	if err := utils.BodyParserWithNull(ctx, &req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"errors": err.Error(), "message": "Invalid request", "status": http.StatusBadRequest})
	}

	// Validate the request
	reqValidator := form_requests.NewCatalogStoreRequest().Validate(&req, ctx.Context())
	if reqValidator != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"errors": reqValidator, "message": "Validation failed", "status": http.StatusBadRequest})
	}

	// Extract user ID from JWT
	claims, err := middleware.GetAuthUser(ctx)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Unauthorized", http.StatusUnauthorized, err.Error(), nil)
	}
	userID := uint(claims["user_id"].(float64))

	// utils.DD(map[string]interface{}{
	// 	// "attachmentUrls":     req.AttachmentUrls,
	// 	"attachmentUrlsJSON": string(attachmentUrlsJSON),
	// })

	createdAt := time.Now()

	catalog := models.Catalog{
		CustomerID:  req.CustomerID,
		CatalogNo:   req.CatalogNo,
		Remark:      req.Remark,
		Description: req.Description,
		IsActive:    req.IsActive,
		CreatedByID: &userID,
		CreatedAt:   &createdAt,
	}

	createdCatalog, err := c.service.CreateCatalog(ctx.Context(), &catalog)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Failed to create catalog", http.StatusInternalServerError, err.Error(), nil)
	}

	params := &dtos.GetCatalogParams{ID: createdCatalog.ID}
	getCatalog, err := c.service.GetCatalogByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Catalog not found", http.StatusNotFound, err.Error(), nil)
	}

	filters := ctx.Locals("filters").(map[string]string)
	paginationMeta := utils.CreatePaginationMeta(filters, 1)

	return utils.GetResponse(ctx, []interface{}{getCatalog}, paginationMeta, "Catalog created successfully", http.StatusCreated, nil, nil)
}

func (c *CatalogController) GetCatalogByID(ctx *fiber.Ctx) error {
	var req dtos.GetCatalogByIDRequest

	if err := ctx.BodyParser(&req); err != nil {
		return utils.GetResponse(ctx, nil, nil, "Catalog not found", http.StatusBadRequest, err.Error(), nil)
	}

	if req.ID == 0 {
		return utils.GetResponse(ctx, nil, nil, "Catalog not found", http.StatusBadRequest, "ID is required", nil)
	}

	params := &dtos.GetCatalogParams{ID: req.ID}
	catalog, err := c.service.GetCatalogByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Catalog not found", http.StatusNotFound, err.Error(), nil)
	}

	catalogArray := []interface{}{catalog}

	filters := ctx.Locals("filters").(map[string]string)
	paginationMeta := utils.CreatePaginationMeta(filters, 1)

	return utils.GetResponse(ctx, catalogArray, paginationMeta, "Catalog fetched successfully", http.StatusOK, nil, nil)
}

// update catalog
func (c *CatalogController) UpdateCatalog(ctx *fiber.Ctx) error {
	var req dtos.UpdateCatalogRequest

	if err := utils.BodyParserWithNull(ctx, &req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"errors": err.Error(), "message": "Invalid request", "status": http.StatusBadRequest})
	}

	// Validate the request
	reqValidator := form_requests.NewCatalogUpdateRequest().Validate(&req, ctx.Context())
	if reqValidator != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"errors": reqValidator, "message": "Validation failed", "status": http.StatusBadRequest})
	}

	params := &dtos.GetCatalogParams{ID: req.ID}
	// Fetch the existing catalog to get the current data
	existingCatalog, err := c.service.GetCatalogByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Catalog not found", http.StatusNotFound, err.Error(), nil)
	}

	// Extract user ID from JWT
	claims, err := middleware.GetAuthUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"errors": err.Error(), "message": "Unauthorized", "status": fiber.StatusUnauthorized})
	}
	userID := uint(claims["user_id"].(float64))

	catalog := models.Catalog{
		ID:          req.ID,
		CustomerID:  req.CustomerID,
		CatalogNo:   req.CatalogNo,
		Remark:      req.Remark,
		Description: req.Description,
		IsActive:    req.IsActive,
		CreatedByID: &existingCatalog.CreatedByID,
		UpdatedByID: &userID,
		CreatedAt:   existingCatalog.CreatedAt,
	}

	updatedCatalog, err := c.service.UpdateCatalog(ctx.Context(), &catalog)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Failed to update catalog", http.StatusInternalServerError, err.Error(), nil)
	}

	params = &dtos.GetCatalogParams{ID: updatedCatalog.ID}
	getCatalog, err := c.service.GetCatalogByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Catalog not found", http.StatusNotFound, err.Error(), nil)
	}

	filters := ctx.Locals("filters").(map[string]string)
	paginationMeta := utils.CreatePaginationMeta(filters, 1)

	return utils.GetResponse(ctx, []interface{}{getCatalog}, paginationMeta, "Catalog updated successfully", http.StatusOK, nil, nil)
}

// delete catalog
func (c *CatalogController) DeleteCatalog(ctx *fiber.Ctx) error {
	var req dtos.DeleteCatalogRequest

	if err := ctx.BodyParser(&req); err != nil {
		return utils.GetResponse(ctx, nil, nil, "Catalog not found", http.StatusBadRequest, err.Error(), nil)
	}

	if req.ID == 0 {
		return utils.GetResponse(ctx, nil, nil, "Catalog not found", http.StatusBadRequest, "ID is required", nil)
	}

	params := &dtos.GetCatalogParams{ID: req.ID}
	// GET catalog by ID
	_, err := c.service.GetCatalogByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Catalog not found", http.StatusNotFound, err.Error(), nil)
	}

	err = c.service.DeleteCatalog(ctx.Context(), req.ID)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Failed to delete catalog", http.StatusInternalServerError, err.Error(), nil)
	}

	return utils.GetResponse(ctx, nil, nil, "Catalog deleted successfully", http.StatusOK, nil, nil)
}

// restore catalog
func (c *CatalogController) RestoreCatalog(ctx *fiber.Ctx) error {
	var req dtos.DeleteCatalogRequest

	if err := ctx.BodyParser(&req); err != nil {
		return utils.GetResponse(ctx, nil, nil, "Catalog not found", http.StatusBadRequest, err.Error(), nil)
	}

	if req.ID == 0 {
		return utils.GetResponse(ctx, nil, nil, "Catalog not found", http.StatusBadRequest, "ID is required", nil)
	}

	isDeleted := 1
	params := &dtos.GetCatalogParams{ID: req.ID, IsDeleted: &isDeleted}
	// GET catalog by ID
	_, err := c.service.GetCatalogByID(ctx.Context(), params)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Catalog not found", http.StatusNotFound, err.Error(), nil)
	}

	err = c.service.RestoreCatalog(ctx.Context(), req.ID)
	if err != nil {
		return utils.GetResponse(ctx, nil, nil, "Failed to restore catalog", http.StatusInternalServerError, err.Error(), nil)
	}

	return utils.GetResponse(ctx, nil, nil, "Catalog restored successfully", http.StatusOK, nil, nil)
}
