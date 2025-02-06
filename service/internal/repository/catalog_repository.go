package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/nibroos/e-proc-api/service/internal/dtos"
	"github.com/nibroos/e-proc-api/service/internal/models"
	"github.com/nibroos/e-proc-api/service/internal/utils"
	"gorm.io/gorm"
)

type CatalogRepository struct {
	db    *gorm.DB
	sqlDB *sqlx.DB
}

func NewCatalogRepository(db *gorm.DB, sqlDB *sqlx.DB) *CatalogRepository {
	return &CatalogRepository{
		db:    db,
		sqlDB: sqlDB,
	}
}

func (r *CatalogRepository) GetCatalogs(ctx context.Context, filters map[string]string) ([]dtos.CatalogListDTO, int, error) {
	catalogs := []dtos.CatalogListDTO{}
	var total int

	// query parent catalog

	query := `SELECT *
    FROM ( 
        SELECT DISTINCT ON (c.id) 
					c.id, c.catalog_no, c.description, c.remark, c.customer_id, c.created_at, c.updated_at, c.deleted_at,
        ct.name as customer_name,
				i.name as item_name,
        cu.name as created_by_name,
        uu.name as updated_by_name

        FROM catalogs c
        LEFT JOIN users cu ON c.created_by_id = cu.id
        LEFT JOIN users uu ON c.updated_by_id = uu.id
        JOIN customers ct ON c.customer_id = ct.id
				JOIN catalog_details cd ON c.id = cd.catalog_id
				JOIN items i ON cd.item_id = i.id
    ) AS alias WHERE 1=1 AND deleted_at IS NULL`

	countQuery := `SELECT COUNT(*) FROM (
        SELECT DISTINCT ON (c.id) 
					c.id, c.catalog_no, c.description, c.remark, c.customer_id, c.created_at, c.updated_at, c.deleted_at,
        ct.name as customer_name,
				i.name as item_name,
        cu.name as created_by_name,
        uu.name as updated_by_name

        FROM catalogs c

        LEFT JOIN users cu ON c.created_by_id = cu.id
        LEFT JOIN users uu ON c.updated_by_id = uu.id
        JOIN customers ct ON c.customer_id = ct.id
				JOIN catalog_details cd ON c.id = cd.catalog_id
				JOIN items i ON cd.item_id = i.id
    ) AS alias WHERE 1=1 AND deleted_at IS NULL`

	var args []interface{}

	i := 1
	for key, value := range filters {
		switch key {
		case "catalog_no", "description", "remark", "customer_name", "item_name":
			if value != "" {
				query += fmt.Sprintf(" AND %s ILIKE $%d", key, i)
				countQuery += fmt.Sprintf(" AND %s ILIKE $%d", key, i)
				args = append(args, "%"+value+"%")
				i++
			}
		}
	}

	if value, ok := filters["customer_id"]; ok && value != "" {
		query += fmt.Sprintf(" AND customer_id = $%d", i)
		countQuery += fmt.Sprintf(" AND customer_id = $%d", i)
		args = append(args, value)
		i++
	}

	if value, ok := filters["global"]; ok && value != "" {
		query += fmt.Sprintf(" AND (catalog_no ILIKE $%d OR description ILIKE $%d OR customer_name ILIKE $%d OR item_name ILIKE $%d OR remark ILIKE $%d)", i, i+1, i+2, i+3, i+4)
		countQuery += fmt.Sprintf(" AND (catalog_no ILIKE $%d OR description ILIKE $%d OR customer_name ILIKE $%d OR item_name ILIKE $%d OR remark ILIKE $%d)", i, i+1, i+2, i+3, i+4)
		args = append(args, "%"+value+"%", "%"+value+"%", "%"+value+"%", "%"+value+"%", "%"+value+"%")
		i += 5
	}

	countArgs := append([]interface{}{}, args...)

	// Channels for concurrent execution
	countChan := make(chan error)
	selectChan := make(chan error)

	orderColumn := utils.GetStringOrDefault(filters["order_column"], "id")
	orderDirection := utils.GetStringOrDefault(filters["order_direction"], "asc")
	query += fmt.Sprintf(" ORDER BY %s %s", orderColumn, orderDirection)

	// Goroutine for count query
	go func() {
		err := r.sqlDB.GetContext(ctx, &total, countQuery, countArgs...)
		countChan <- err
	}()

	perPage := utils.GetIntOrDefault(filters["per_page"], 10)
	currentPage := utils.GetIntOrDefault(filters["page"], 1)

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, perPage, (currentPage-1)*perPage)

	// Goroutine for select query
	go func() {
		err := r.sqlDB.SelectContext(ctx, &catalogs, query, args...)

		log.Println("catalogs", catalogs)

		if err != nil {
			selectChan <- err
			return
		}

		// query child catalog
		childQuery := `SELECT *
		FROM (
			SELECT DISTINCT ON (cd.id)
				cd.id, cd.catalog_id, cd.item_id, cd.price_buy, cd.price_sell, cd.remark as detail_remark, cd.created_at, cd.updated_at, cd.deleted_at,
			i.name as item_name,
			c.catalog_no as catalog_no,
			c.description as description,
			c.remark as remark,
			c.customer_id as customer_id,
			ct.name as customer_name

			FROM catalog_details cd
			JOIN items i ON cd.item_id = i.id
			JOIN catalogs c ON cd.catalog_id = c.id
			JOIN customers ct ON c.customer_id = ct.id
		) AS alias WHERE 1=1 AND deleted_at IS NULL`

		// where catalog_id in
		catalogIDs := []uint{}
		for _, catalog := range catalogs {
			catalogIDs = append(catalogIDs, catalog.ID)
		}

		childQuery += fmt.Sprintf(" AND catalog_id IN (%s)", utils.JoinUintArray(catalogIDs))

		var childCatalogs []dtos.CatalogChildListDTO
		err = r.sqlDB.SelectContext(ctx, &childCatalogs, childQuery)

		// assign child catalog to parent catalog
		for i, catalog := range catalogs {
			for _, childCatalog := range childCatalogs {
				if catalog.ID == childCatalog.CatalogID {
					catalog.CatalogDetail = append(catalog.CatalogDetail, childCatalog)
				}
			}
			catalogs[i] = catalog
		}

		selectChan <- err
	}()

	// Wait for both goroutines to finish
	countErr := <-countChan
	selectErr := <-selectChan

	if countErr != nil {
		return nil, 0, countErr
	}

	if selectErr != nil {
		return nil, 0, selectErr
	}

	return catalogs, total, nil
}

func (r *CatalogRepository) GetCatalogByID(ctx context.Context, params *dtos.GetCatalogParams) (*dtos.CatalogDetailDTO, error) {
	var catalog dtos.CatalogDetailDTO
	// deletedAt := params.IsDeleted

	query := `SELECT c.id, c.catalog_no, c.description, c.remark, c.customer_id, c.is_active,
	ct.name as customer_name

	FROM catalogs c
	LEFT JOIN users cu ON c.created_by_id = cu.id
	LEFT JOIN users uu ON c.updated_by_id = uu.id
	JOIN customers ct ON c.customer_id = ct.id
	WHERE 1=1`

	var args []interface{}

	i := 1
	query += " AND c.id = $1"
	args = append(args, params.ID)
	i++

	isDeletedQuery := ` AND c.deleted_at IS NULL`
	if params.IsDeleted != nil && *params.IsDeleted == 1 {
		isDeletedQuery = " AND c.deleted_at IS NOT NULL"
	}

	query += isDeletedQuery

	if err := r.sqlDB.Get(&catalog, query, args...); err != nil {
		return nil, err
	}

	// query child catalog
	childQuery := `SELECT *
	FROM (
		SELECT DISTINCT ON (cd.id)
			cd.id, cd.catalog_id, cd.item_id, cd.price_buy, cd.price_sell, cd.remark as detail_remark, cd.created_at, cd.updated_at, cd.deleted_at,
		i.name as item_name,
		c.catalog_no as catalog_no,
		c.description as description,
		c.remark as remark,
		c.customer_id as customer_id,
		ct.name as customer_name

		FROM catalog_details cd
		JOIN items i ON cd.item_id = i.id
		JOIN catalogs c ON cd.catalog_id = c.id
		JOIN customers ct ON c.customer_id = ct.id
	) AS alias WHERE 1=1 AND deleted_at IS NULL`

	childQuery += " AND catalog_id = $1"
	var childCatalogs []dtos.CatalogChildListDTO
	err := r.sqlDB.SelectContext(ctx, &childCatalogs, childQuery, params.ID)

	if err != nil {
		return nil, err
	}

	// assign child catalog to parent catalog
	catalog.CatalogDetail = append(catalog.CatalogDetail, childCatalogs...)

	return &catalog, nil
}

// BeginTransaction starts a new transaction
func (r *CatalogRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}

func (r *CatalogRepository) CreateCatalog(tx *gorm.DB, catalog *models.Catalog) error {
	if err := tx.Create(catalog).Error; err != nil {
		return err
	}
	return nil
}

func (r *CatalogRepository) UpdateCatalog(tx *gorm.DB, catalog *models.Catalog) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(catalog).Error; err != nil {
			return err
		}
		return nil
	})

}

func (r *CatalogRepository) DeleteCatalog(tx *gorm.DB, id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// if err := tx.Unscoped().Delete(&models.Catalog{}, id).Error; err != nil {
		if err := tx.Delete(&models.Catalog{}, id).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *CatalogRepository) RestoreCatalog(tx *gorm.DB, id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE catalogs SET deleted_at = NULL WHERE id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
}
