package postgres

import (
	"fmt"
	"log"

	"github.com/donatetohospitals/donatetohospitals-web/core"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

type DatabaseConfiguration struct {
	DbName   string
	Port     string
	User     string
	Password string
}

type Supplier struct {
	gorm.Model
	Email       string `gorm:"type:varchar(100);unique_index;not null"`
	Geo         string `gorm:"index:geo"` // state etc. create index with name `loc` for address
	ImageUrl    string `gorm:"size:255"`  // set field size to 255
	Items       []Item `gorm:"foreignkey:SupplierRefer"`
	IsAllocated bool
}

type Item struct {
	gorm.Model
	Name          string
	Count         int
	SupplierRefer uint
}

func newPostgresClient(configuration *DatabaseConfiguration) (*gorm.DB, error) {

	connectionString := fmt.Sprintf(
		"host=127.0.0.1 port=%v user=%v dbname=%v password=%v sslmode=disable",
		configuration.Port,
		configuration.User,
		configuration.DbName,
		configuration.Password,
	)

	createdDb, connectionError := gorm.Open("postgres", connectionString)

	if connectionError != nil {
		log.Fatal(connectionError)
		return nil, errors.Wrap(connectionError, "repository.Postgres")
	}

	// NOTE (daniel): Database connection should be closed on main.go
	return createdDb, nil
}

func NewSupplierRepository(client *gorm.DB) core.SupplierRepository {
	client.AutoMigrate(&core.Supplier{})
	return nil
}

func NewStaffRepository(client *gorm.DB) core.StaffRepository {

	return nil
}

func NewItemRepository(client *gorm.DB) core.ItemRepository {
	client.AutoMigrate(&core.Item{})
	return nil
}
