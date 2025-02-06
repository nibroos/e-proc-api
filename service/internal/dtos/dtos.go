package dtos

import (
	"time"
)

type GetUsersRequest struct {
	Global         string  `json:"global"`
	Username       string  `json:"username"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	PerPage        *string `json:"per_page" default:"10"`         // Default per_page to 10
	Page           *string `json:"page" default:"1"`              // Default page to 1
	OrderColumn    string  `json:"order_column" default:"id"`     // Default order column to "id"
	OrderDirection string  `json:"order_direction" default:"asc"` // Default order direction to "asc"
}

type CreateUserRequest struct {
	Name     string   `json:"name"`
	Username *string  `json:"username"`
	Email    string   `json:"email"`
	Address  *string  `json:"address"`
	Password string   `json:"password"`
	RoleIDs  []uint32 `json:"role_ids"`
}

type UpdateUserRequest struct {
	ID       uint     `json:"id"`
	Username *string  `json:"username"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Address  *string  `json:"address"`
	Password *string  `json:"password"`
	RoleIDs  []uint32 `json:"role_ids"`
}

type GetUserByIDParams struct {
	ID        uint `json:"id"`
	IsDeleted *int
}

type GetUserByIDRequest struct {
	ID uint `json:"id"`
}

type DeleteUserRequest struct {
	ID uint `json:"id"`
}

type UserListDTO struct {
	ID       int     `json:"id"`
	Username *string `json:"username"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
}

type UserDetailDTO struct {
	ID           uint     `json:"id"`
	Name         string   `json:"name"`
	Username     *string  `json:"username"`
	Email        string   `json:"email"`
	Address      *string  `json:"address"`
	Password     *string  `json:"password"`
	CustomerName *string  `json:"customer_name" db:"customer_name"`
	Roles        []string `json:"roles"`
	Permissions  []string `json:"permissions"`
	CreatedAt    *string  `json:"created_at"`
}
type GetUsersResult struct {
	Users []UserListDTO
	Total int
	Err   error
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RegisterRequest struct {
	Name           string  `json:"name"`
	Username       *string `json:"username"`
	Email          string  `json:"email"`
	Address        *string `json:"address"`
	Password       string  `json:"password"`
	CustomerTypeID uint    `json:"customer_type_id"`
}

type CreateIdentifierRequest struct {
	UserID           uint   `json:"user_id"`
	TypeIdentifierID uint   `json:"type_identifier_id"`
	RefNum           string `json:"ref_num"`
	Status           uint   `json:"status"`
}

type UpdateIdentifierRequest struct {
	ID               uint   `json:"id"`
	UserID           uint   `json:"user_id"`
	TypeIdentifierID *uint  `json:"type_identifier_id"`
	RefNum           string `json:"ref_num"`
	Status           uint   `json:"status"`
}

type GetIdentifierByIDRequest struct {
	ID uint `json:"id"`
}

type GetIdentifierParams struct {
	ID        uint
	UserID    uint
	IsDeleted *int
}

func NewGetIdentifierParams(id uint) *GetIdentifierParams {
	defaultIsDeleted := 0
	return &GetIdentifierParams{
		ID:        id,
		IsDeleted: &defaultIsDeleted,
	}
}

type DeleteIdentifierRequest struct {
	ID uint `json:"id"`
}

type IdentifierListDTO struct {
	ID                 int     `json:"id" db:"id"`
	UserID             uint    `json:"user_id" db:"user_id"`
	UserName           string  `json:"user_name" db:"user_name"`
	TypeIdentifierID   uint    `json:"type_identifier_id" db:"type_identifier_id"`
	TypeIdentifierName string  `json:"type_identifier_name" db:"type_identifier_name"`
	RefNum             string  `json:"ref_num" db:"ref_num"`
	Status             uint    `json:"status" db:"status"`
	CreatedAt          *string `json:"created_at" db:"created_at"`
	UpdatedAt          *string `json:"updated_at" db:"updated_at"`
}

type IdentifierDetailDTO struct {
	ID                 uint       `json:"id" db:"id"`
	UserID             uint       `json:"user_id" db:"user_id"`
	UserName           string     `json:"user_name" db:"user_name"`
	TypeIdentifierID   uint       `json:"type_identifier_id" db:"type_identifier_id"`
	TypeIdentifierName string     `json:"type_identifier_name" db:"type_identifier_name"`
	RefNum             string     `json:"ref_num" db:"ref_num"`
	Status             uint       `json:"status" db:"status"`
	CreatedAt          *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at" db:"deleted_at"`
}
type ListIdentifiersResult struct {
	Identifiers []IdentifierListDTO
	Total       int
	Err         error
}

type CreateContactRequest struct {
	TypeContactID uint   `json:"type_contact_id"`
	UserID        uint   `json:"user_id"`
	RefNum        string `json:"ref_num"`
	Status        uint   `json:"status"`
}

type UpdateContactRequest struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	TypeContactID *uint  `json:"type_contact_id"`
	RefNum        string `json:"ref_num"`
	Status        uint   `json:"status"`
}

type GetContactByIDRequest struct {
	ID uint `json:"id"`
}

type GetContactParams struct {
	ID        uint
	UserID    uint
	IsDeleted *int
}

func NewGetContactParams(id uint) *GetContactParams {
	defaultIsDeleted := 0
	return &GetContactParams{
		ID:        id,
		IsDeleted: &defaultIsDeleted,
	}
}

type DeleteContactRequest struct {
	ID uint `json:"id"`
}

type ContactListDTO struct {
	ID              int     `json:"id" db:"id"`
	UserID          uint    `json:"user_id" db:"user_id"`
	UserName        string  `json:"user_name" db:"user_name"`
	TypeContactID   uint    `json:"type_contact_id" db:"type_contact_id"`
	TypeContactName string  `json:"type_contact_name" db:"type_contact_name"`
	RefNum          string  `json:"ref_num" db:"ref_num"`
	Status          uint    `json:"status" db:"status"`
	CreatedAt       *string `json:"created_at" db:"created_at"`
	UpdatedAt       *string `json:"updated_at" db:"updated_at"`
}

type ContactDetailDTO struct {
	ID              uint       `json:"id" db:"id"`
	UserID          uint       `json:"user_id" db:"user_id"`
	UserName        string     `json:"user_name" db:"user_name"`
	TypeContactID   uint       `json:"type_contact_id" db:"type_contact_id"`
	TypeContactName string     `json:"type_contact_name" db:"type_contact_name"`
	RefNum          string     `json:"ref_num" db:"ref_num"`
	Status          uint       `json:"status" db:"status"`
	CreatedAt       *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at" db:"deleted_at"`
}
type ListContactsResult struct {
	Contacts []ContactListDTO
	Total    int
	Err      error
}

type CreateAddressRequest struct {
	TypeAddressID uint   `json:"type_address_id"`
	UserID        uint   `json:"user_id"`
	RefNum        string `json:"ref_num"`
	Status        uint   `json:"status"`
}

type UpdateAddressRequest struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	TypeAddressID *uint  `json:"type_address_id"`
	RefNum        string `json:"ref_num"`
	Status        uint   `json:"status"`
}

type GetAddressByIDRequest struct {
	ID uint `json:"id"`
}

type GetAddressParams struct {
	ID        uint
	UserID    uint
	IsDeleted *int
}

func NewGetAddressParams(id uint) *GetAddressParams {
	defaultIsDeleted := 0
	return &GetAddressParams{
		ID:        id,
		IsDeleted: &defaultIsDeleted,
	}
}

type DeleteAddressRequest struct {
	ID uint `json:"id"`
}

type AddressListDTO struct {
	ID              int     `json:"id" db:"id"`
	UserID          uint    `json:"user_id" db:"user_id"`
	UserName        string  `json:"user_name" db:"user_name"`
	TypeAddressID   uint    `json:"type_address_id" db:"type_address_id"`
	TypeAddressName string  `json:"type_address_name" db:"type_address_name"`
	RefNum          string  `json:"ref_num" db:"ref_num"`
	Status          uint    `json:"status" db:"status"`
	CreatedAt       *string `json:"created_at" db:"created_at"`
	UpdatedAt       *string `json:"updated_at" db:"updated_at"`
}

type AddressDetailDTO struct {
	ID              uint       `json:"id" db:"id"`
	UserID          uint       `json:"user_id" db:"user_id"`
	UserName        string     `json:"user_name" db:"user_name"`
	TypeAddressID   uint       `json:"type_address_id" db:"type_address_id"`
	TypeAddressName string     `json:"type_address_name" db:"type_address_name"`
	RefNum          string     `json:"ref_num" db:"ref_num"`
	Status          uint       `json:"status" db:"status"`
	CreatedAt       *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at" db:"deleted_at"`
}
type ListAddressesResult struct {
	Addresses []AddressListDTO
	Total     int
	Err       error
}

// type Scheduler struct {
// 	ID          uint       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
// 	Name        string     `json:"name" gorm:"column:name"`
// 	Description string     `json:"description" gorm:"column:description"`
// 	Cron        string     `json:"cron" gorm:"column:cron"`
// 	Payload     string     `json:"payload" gorm:"column:payload"`
// 	Status      string     `json:"status" gorm:"column:status"`
// 	StartAt     time.Time  `json:"start_at" gorm:"column:start_at"`
// 	EndAt       *time.Time `json:"end_at" gorm:"column:end_at"`
// }

type SchedulerListDTO struct {
	ID          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Cron        string  `json:"cron" db:"cron"`
	Payload     string  `json:"payload" db:"payload"`
	Status      string  `json:"status" db:"status"`
	StartAt     *string `json:"start_at" db:"start_at"`
	EndAt       *string `json:"end_at" db:"end_at"`
}

type CreateCatalogRequest struct {
	CustomerID  uint   `json:"customer_id"`
	CatalogNo   string `json:"catalog_no"`
	Description string `json:"description"`
	Remark      string `json:"remark"`
	IsActive    int8   `json:"is_active"`
}

type UpdateCatalogRequest struct {
	ID          uint   `json:"id"`
	CustomerID  uint   `json:"customer_id"`
	CatalogNo   string `json:"catalog_no"`
	Description string `json:"description"`
	Remark      string `json:"remark"`
	IsActive    int8   `json:"is_active"`
}

type GetCatalogByIDRequest struct {
	ID uint `json:"id"`
}

type GetCatalogParams struct {
	ID        uint
	IsDeleted *int
}

func NewGetCatalogParams(id uint) *GetCatalogParams {
	defaultIsDeleted := 0
	return &GetCatalogParams{
		ID:        id,
		IsDeleted: &defaultIsDeleted,
	}
}

type DeleteCatalogRequest struct {
	ID uint `json:"id"`
}

type CatalogListDTO struct {
	ID            uint                  `json:"id" db:"id"`
	CatalogNo     string                `json:"catalog_no" db:"catalog_no"`
	Description   string                `json:"description" db:"description"`
	Remark        string                `json:"remark" db:"remark"`
	IsActive      int8                  `json:"is_active"`
	CustomerID    uint                  `json:"customer_id" db:"customer_id"`
	CustomerName  string                `json:"customer_name" db:"customer_name"`
	ItemName      string                `json:"item_name" db:"item_name"`
	CreatedByName *string               `json:"created_by_name" db:"created_by_name"`
	UpdatedByName *string               `json:"updated_by_name" db:"updated_by_name"`
	CreatedAt     *string               `json:"created_at" db:"created_at"`
	UpdatedAt     *string               `json:"updated_at" db:"updated_at"`
	DeleteAt      *string               `json:"deleted_at" db:"deleted_at"`
	CatalogDetail []CatalogChildListDTO `json:"catalog_details"`
}

type CatalogChildListDTO struct {
	ID           uint    `json:"id" db:"id"`
	CatalogID    uint    `json:"catalog_id" db:"catalog_id"`
	ItemID       uint    `json:"item_id" db:"item_id"`
	PriceBuy     int     `json:"price_buy" db:"price_buy"`
	PriceSell    int     `json:"price_sell" db:"price_sell"`
	Remark       string  `json:"remark" db:"remark"`
	DetailRemark string  `json:"detail_remark" db:"detail_remark"`
	IsActive     int8    `json:"is_active" db:"is_active"`
	ItemName     string  `json:"item_name" db:"item_name"`
	CatalogNo    string  `json:"catalog_no" db:"catalog_no"`
	Description  string  `json:"description" db:"description"`
	CustomerID   uint    `json:"customer_id" db:"customer_id"`
	CustomerName string  `json:"customer_name" db:"customer_name"`
	CreatedAt    *string `json:"created_at" db:"created_at"`
	UpdatedAt    *string `json:"updated_at" db:"updated_at"`
	DeletedAt    *string `json:"deleted_at" db:"deleted_at"`
}

type CatalogDetailDTO struct {
	ID            uint                  `json:"id" db:"id"`
	CatalogNo     string                `json:"catalog_no" db:"catalog_no"`
	Description   string                `json:"description" db:"description"`
	Remark        string                `json:"remark" db:"remark"`
	IsActive      int8                  `json:"is_active" db:"is_active"`
	CustomerID    uint                  `json:"customer_id" db:"customer_id"`
	CustomerName  string                `json:"customer_name" db:"customer_name"`
	CreatedByID   uint                  `json:"created_by_id" db:"created_by_id"`
	CreatedAt     *time.Time            `json:"created_at" db:"created_at"`
	CatalogDetail []CatalogChildListDTO `json:"catalog_details"`
}
type GetCatalogsResult struct {
	Catalogs []CatalogListDTO
	Total    int
	Err      error
}

type CreateCustomerRequest struct {
	CustomerID uint   `json:"customer_id"`
	Name       uint   `json:"name"`
	Email      string `json:"email"`
	IsActive   int8   `json:"is_active"`
}

type UpdateCustomerRequest struct {
	ID         uint   `json:"id"`
	CustomerID uint   `json:"customer_id"`
	Name       uint   `json:"name"`
	Email      string `json:"email"`
	IsActive   int8   `json:"is_active"`
}

type GetCustomerByIDRequest struct {
	ID uint `json:"id"`
}

type GetCustomerParams struct {
	ID        uint
	IsDeleted *int
}

func NewGetCustomerParams(id uint) *GetCustomerParams {
	defaultIsDeleted := 0
	return &GetCustomerParams{
		ID:        id,
		IsDeleted: &defaultIsDeleted,
	}
}

type DeleteCustomerRequest struct {
	ID uint `json:"id"`
}

type CustomerListDTO struct {
	ID            uint    `json:"id" db:"id"`
	Name          uint    `json:"name" db:"name"`
	Email         string  `json:"email" db:"email"`
	IsActive      int8    `json:"is_active"`
	CustomerID    uint    `json:"customer_id" db:"customer_id"`
	CustomerName  string  `json:"customer_name" db:"customer_name"`
	CreatedByName *string `json:"created_by_name" db:"created_by_name"`
	UpdatedByName *string `json:"updated_by_name" db:"updated_by_name"`
	CreatedAt     *string `json:"created_at" db:"created_at"`
	UpdatedAt     *string `json:"updated_at" db:"updated_at"`
	DeleteAt      *string `json:"deleted_at" db:"deleted_at"`
}

type CustomerDetailDTO struct {
	ID            uint       `json:"id" db:"id"`
	Name          uint       `json:"name" db:"name"`
	Email         string     `json:"email" db:"email"`
	IsActive      int8       `json:"is_active"`
	CustomerID    uint       `json:"customer_id" db:"customer_id"`
	CreatedByID   uint       `json:"created_by_id" db:"created_by_id"`
	UpdatedByID   *uint      `json:"updated_by_id" db:"updated_by_id"`
	CreatedByName *string    `json:"created_by_name" db:"created_by_name"`
	UpdatedByName *string    `json:"updated_by_name" db:"updated_by_name"`
	CreatedAt     *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at" db:"deleted_at"`
}
type GetCustomersResult struct {
	Customers []CustomerListDTO
	Total     int
	Err       error
}
