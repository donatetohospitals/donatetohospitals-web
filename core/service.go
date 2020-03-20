package core

type ServiceConfiguration struct {
	Timeout      int // In seconds
	SupplierRepo SupplierRepository
	StaffRepo    StaffRepository
	ItemRepo     ItemRepository
}

type DonationService interface {
	// Use case signatures here
}
