package service

import (
	"context"

	"github.com/nibroos/e-proc-api/service/internal/dtos"
	"github.com/nibroos/e-proc-api/service/internal/models"
	"github.com/nibroos/e-proc-api/service/internal/repository"
)

type CatalogService struct {
	repo *repository.CatalogRepository
}

func NewCatalogService(repo *repository.CatalogRepository) *CatalogService {
	return &CatalogService{repo: repo}
}

func (s *CatalogService) GetCatalogs(ctx context.Context, filters map[string]string) ([]dtos.CatalogListDTO, int, error) {

	resultChan := make(chan dtos.GetCatalogsResult, 1)

	go func() {
		catalogs, total, err := s.repo.GetCatalogs(ctx, filters)
		resultChan <- dtos.GetCatalogsResult{Catalogs: catalogs, Total: total, Err: err}
	}()

	select {
	case res := <-resultChan:
		return res.Catalogs, res.Total, res.Err
	case <-ctx.Done():
		return nil, 0, ctx.Err()
	}
}

func (s *CatalogService) CreateCatalog(ctx context.Context, catalog *models.Catalog) (*models.Catalog, error) {
	// Transaction handling
	tx := s.repo.BeginTransaction()
	if err := tx.Error; err != nil {
		return nil, err
	}

	// Create catalog
	if err := s.repo.CreateCatalog(tx, catalog); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return catalog, nil
}

func (s *CatalogService) GetCatalogByID(ctx context.Context, params *dtos.GetCatalogParams) (*dtos.CatalogDetailDTO, error) {
	catalogChan := make(chan *dtos.CatalogDetailDTO, 1)
	errChan := make(chan error, 1)

	go func() {
		catalog, err := s.repo.GetCatalogByID(ctx, params)
		if err != nil {
			errChan <- err
			return
		}
		catalogChan <- catalog
	}()

	select {
	case catalog := <-catalogChan:
		return catalog, nil
	case err := <-errChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (s *CatalogService) UpdateCatalog(ctx context.Context, catalog *models.Catalog) (*models.Catalog, error) {
	// Transaction handling
	tx := s.repo.BeginTransaction()
	if err := tx.Error; err != nil {
		return nil, err
	}

	// Update catalog
	if err := s.repo.UpdateCatalog(tx, catalog); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return catalog, nil
}

func (s *CatalogService) DeleteCatalog(ctx context.Context, id uint) error {
	// Transaction handling
	tx := s.repo.BeginTransaction()
	if err := tx.Error; err != nil {
		return err
	}

	// Delete catalog
	if err := s.repo.DeleteCatalog(tx, id); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (s *CatalogService) RestoreCatalog(ctx context.Context, id uint) error {
	// Transaction handling
	tx := s.repo.BeginTransaction()
	if err := tx.Error; err != nil {
		return err
	}

	// Restore catalog
	if err := s.repo.RestoreCatalog(tx, id); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
