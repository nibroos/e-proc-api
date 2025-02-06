package form_requests

import (
	"context"

	"github.com/nibroos/e-proc-api/service/internal/dtos"
	"github.com/thedevsaddam/govalidator"
)

// CatalogStoreRequest handles the validation for the RegisterRequest.
type CatalogStoreRequest struct {
	Validator *govalidator.Validator
}

// NewRegisterStoreRequest creates a new instance of CatalogStoreRequest.
func NewCatalogStoreRequest() *CatalogStoreRequest {
	v := govalidator.New(govalidator.Options{})
	return &CatalogStoreRequest{Validator: v}
}

// Validate validates the RegisterRequest.
func (r *CatalogStoreRequest) Validate(req *dtos.CreateCatalogRequest, ctx context.Context) map[string]string {
	// utils.DD(req)
	rules := govalidator.MapData{
		"customer_id": []string{"required", "exists:customers,id"},
		"catalog_no":  []string{"required", "unique:catalogs,catalog_no"},
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
