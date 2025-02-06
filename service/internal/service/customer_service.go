package service

import (
	"context"
	"fmt"

	"github.com/nibroos/e-proc-api/service/internal/dtos"
	"github.com/nibroos/e-proc-api/service/internal/models"
	"github.com/nibroos/e-proc-api/service/internal/repository"
)

type CustomerService struct {
	repo *repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (c *CustomerService) GetCustomers(ctx context.Context, filters map[string]string) ([]dtos.CustomerListDTO, int, error) {

	resultChan := make(chan dtos.GetCustomersResult, 1)

	go func() {
		customers, total, err := c.repo.GetCustomers(ctx, filters)
		resultChan <- dtos.GetCustomersResult{Customers: customers, Total: total, Err: err}
	}()

	select {
	case res := <-resultChan:
		return res.Customers, res.Total, res.Err
	case <-ctx.Done():
		return nil, 0, ctx.Err()
	}
}
func (c *CustomerService) CreateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	// Transaction handling
	tx := c.repo.BeginTransaction()
	if tx == nil {
		return nil, fmt.Errorf("failed to begin transaction")
	}
	if err := tx.Error; err != nil {
		return nil, err
	}

	// Create customer
	if err := c.repo.CreateCustomer(tx, customer); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *CustomerService) GetCustomerByID(ctx context.Context, params *dtos.GetCustomerParams) (*dtos.CustomerDetailDTO, error) {
	customerChan := make(chan *dtos.CustomerDetailDTO, 1)
	errChan := make(chan error, 1)

	go func() {
		customer, err := c.repo.GetCustomerByID(ctx, params)
		if err != nil {
			errChan <- err
			return
		}
		customerChan <- customer
	}()

	select {
	case customer := <-customerChan:
		return customer, nil
	case err := <-errChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (c *CustomerService) UpdateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	// Transaction handling
	tx := c.repo.BeginTransaction()
	if err := tx.Error; err != nil {
		return nil, err
	}

	// Update customer
	if err := c.repo.UpdateCustomer(tx, customer); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *CustomerService) DeleteCustomer(ctx context.Context, id uint) error {
	// Transaction handling
	tx := c.repo.BeginTransaction()
	if err := tx.Error; err != nil {
		return err
	}

	// Delete customer
	if err := c.repo.DeleteCustomer(tx, id); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (c *CustomerService) RestoreCustomer(ctx context.Context, id uint) error {
	// Transaction handling
	tx := c.repo.BeginTransaction()
	if err := tx.Error; err != nil {
		return err
	}

	// Restore customer
	if err := c.repo.RestoreCustomer(tx, id); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
