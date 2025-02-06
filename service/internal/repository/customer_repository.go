package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nibroos/e-proc-api/service/internal/dtos"
	"github.com/nibroos/e-proc-api/service/internal/models"
	"github.com/nibroos/e-proc-api/service/internal/utils"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db    *gorm.DB
	sqlDB *sqlx.DB
}

func NewCustomerRepository(db *gorm.DB, sqlDB *sqlx.DB) *CustomerRepository {
	return &CustomerRepository{
		db:    db,
		sqlDB: sqlDB,
	}
}

func (r *CustomerRepository) GetCustomers(ctx context.Context, filters map[string]string) ([]dtos.CustomerListDTO, int, error) {
	educations := []dtos.CustomerListDTO{}
	var total int

	query := `SELECT *
    FROM ( 
        SELECT e.id, e.name, e.email, e.customer_type_id, e.created_at, e.updated_at, e.deleted_at,
        m.name as customer_type_name,
        cu.name as created_by_name,
        uu.name as updated_by_name

        FROM educations e
        JOIN users cu ON e.created_by_id = cu.id
        LEFT JOIN users uu ON e.updated_by_id = uu.id
        JOIN mix_values m ON e.customer_type_id = m.id
    ) AS alias WHERE 1=1 AND deleted_at IS NULL`

	countQuery := `SELECT COUNT(*) FROM (
        SELECT e.id, e.name, e.email, e.customer_type_id, e.created_at, e.updated_at, e.deleted_at,
        m.name as customer_type_name,
        cu.name as created_by_name,
        uu.name as updated_by_name

        FROM educations e
        JOIN users cu ON e.created_by_id = cu.id
        LEFT JOIN users uu ON e.updated_by_id = uu.id
        JOIN mix_values m ON e.customer_type_id = m.id
    ) AS alias WHERE 1=1 AND deleted_at IS NULL`

	var args []interface{}

	i := 1
	for key, value := range filters {
		switch key {
		case "name", "email":
			if value != "" {
				query += fmt.Sprintf(" AND %s ILIKE $%d", key, i)
				countQuery += fmt.Sprintf(" AND %s ILIKE $%d", key, i)
				args = append(args, "%"+value+"%")
				i++
			}
		}
	}

	if value, ok := filters["customer_type_id"]; ok && value != "" {
		query += fmt.Sprintf(" AND customer_type_id = $%d", i)
		countQuery += fmt.Sprintf(" AND customer_type_id = $%d", i)
		args = append(args, value)
		i++
	}

	if value, ok := filters["global"]; ok && value != "" {
		query += fmt.Sprintf(" AND (name ILIKE $%d OR email ILIKE $%d OR customer_type_name ILIKE $%d)", i, i+1, i+2)
		countQuery += fmt.Sprintf(" AND (name ILIKE $%d OR email ILIKE $%d OR customer_type_name ILIKE $%d)", i, i+1, i+2)
		args = append(args, "%"+value+"%", "%"+value+"%", "%"+value+"%")
		i += 3
	}

	countArgs := append([]interface{}{}, args...)

	// Channels for concurrent execution
	countChan := make(chan error)
	selectChan := make(chan error)

	// Goroutine for count query
	go func() {
		err := r.sqlDB.GetContext(ctx, &total, countQuery, countArgs...)
		countChan <- err
	}()

	orderColumn := utils.GetStringOrDefault(filters["order_column"], "id")
	orderDirection := utils.GetStringOrDefault(filters["order_direction"], "asc")
	query += fmt.Sprintf(" ORDER BY %s %s", orderColumn, orderDirection)

	perPage := utils.GetIntOrDefault(filters["per_page"], 10)
	currentPage := utils.GetIntOrDefault(filters["page"], 1)

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, perPage, (currentPage-1)*perPage)

	// Goroutine for select query
	go func() {
		err := r.sqlDB.SelectContext(ctx, &educations, query, args...)
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

	return educations, total, nil
}

func (r *CustomerRepository) GetCustomerByID(ctx context.Context, params *dtos.GetCustomerParams) (*dtos.CustomerDetailDTO, error) {
	var education dtos.CustomerDetailDTO
	// deletedAt := params.IsDeleted

	query := `SELECT e.*,
	m.name as customer_type_name,
	cu.name as created_by_name,
	uu.name as updated_by_name

	FROM educations e
	JOIN users cu ON e.created_by_id = cu.id
	LEFT JOIN users uu ON e.updated_by_id = uu.id
	JOIN mix_values m ON e.customer_type_id = m.id
	WHERE 1=1`

	var args []interface{}

	i := 1
	query += " AND e.id = $1"
	args = append(args, params.ID)
	i++

	isDeletedQuery := ` AND e.deleted_at IS NULL`
	if params.IsDeleted != nil && *params.IsDeleted == 1 {
		isDeletedQuery = " AND e.deleted_at IS NOT NULL"
	}

	query += isDeletedQuery

	if err := r.sqlDB.Get(&education, query, args...); err != nil {
		return nil, err
	}

	return &education, nil
}

// BeginTransaction starts a new transaction
func (r *CustomerRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}

func (r *CustomerRepository) CreateCustomer(tx *gorm.DB, education *models.Customer) error {
	if err := tx.Create(education).Error; err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepository) UpdateCustomer(tx *gorm.DB, education *models.Customer) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(education).Error; err != nil {
			return err
		}
		return nil
	})

}

func (r *CustomerRepository) DeleteCustomer(tx *gorm.DB, id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// if err := tx.Unscoped().Delete(&models.Customer{}, id).Error; err != nil {
		if err := tx.Delete(&models.Customer{}, id).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *CustomerRepository) RestoreCustomer(tx *gorm.DB, id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE educations SET deleted_at = NULL WHERE id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
}
