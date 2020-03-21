package core

import "github.com/jinzhu/gorm"

type ServiceConfiguration struct {
	Timeout      int // In seconds
	SupplierRepo SupplierRepository
	StaffRepo    StaffRepository
	ItemRepo     ItemRepository
	DB           *gorm.DB // NOTE(daniel): Temporary utility for http handlers to use
}

type DonationService interface {
	// Use case signatures here
	GetDatabase() *gorm.DB // Returns the database being used. This is a temporary workaround
}
