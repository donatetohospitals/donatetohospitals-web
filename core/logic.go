package core

import (
	"time"

	"github.com/jinzhu/gorm"
)

// TODO (daniel): move these entities to their own package eventually
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

// TODO (daniel): Split this further into its own package
type donationService struct {
	timeout      time.Duration
	supplierRepo SupplierRepository
	staffRepo    StaffRepository
	itemRepo     ItemRepository
}

func NewDonationService(sc ServiceConfiguration) DonationService {
	// Accept an interface, return a concrete type
	return donationService{
		supplierRepo: sc.SupplierRepo,
		staffRepo:    sc.StaffRepo,
		itemRepo:     sc.ItemRepo,
	}
}
