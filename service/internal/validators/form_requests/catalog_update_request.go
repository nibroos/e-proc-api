package form_requests

import (
	"context"
	"fmt"

	"github.com/nibroos/e-proc-api/service/internal/dtos"
	"github.com/thedevsaddam/govalidator"
)

// CatalogUpdateRequest handles the validation for the RegisterRequest.
type CatalogUpdateRequest struct {
	Validator *govalidator.Validator
}

// NewRegisterUpdateRequest creates a new instance of CatalogUpdateRequest.
func NewCatalogUpdateRequest() *CatalogUpdateRequest {
	v := govalidator.New(govalidator.Options{})
	return &CatalogUpdateRequest{Validator: v}
}

// Validate validates the RegisterRequest.
func (r *CatalogUpdateRequest) Validate(req *dtos.UpdateCatalogRequest, ctx context.Context) map[string]string {
	// utils.DD(req)
	rules := govalidator.MapData{
		"customer_id": []string{"required", "exists:customers,id"},
		"catalog_no":  []string{"required", fmt.Sprintf("unique_ig:catalogs,catalog_no,%d", req.ID)},
		"remark":      []string{},
		"description": []string{},
		"is_active":   []string{"required"},
	}

	opts := govalidator.Options{
		Data:  req,
		Rules: rules,
	}

	v := govalidator.New(opts)
	mappedErrors := v.ValidateStruct()

	if len(mappedErrors) == 0 {
		return nil
	}

	errors := make(map[string]string)
	for field, err := range mappedErrors {
		errors[field] = err[0]
	}
	return errors
}
