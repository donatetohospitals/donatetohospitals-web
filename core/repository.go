package core

type DatabaseConfiguration struct {
	DbName   string
	Port     string
	User     string
	Password string
}

// TODO (daniel): Move repos to their own packages, preferably alongside the entities
type StaffRepository interface {
	GetAll()
	GetStaffById()
}

type SupplierRepository interface {
	GetAll()
	GetSupplierById()
}

type ItemRepository interface {
	GetAll()
	GetItemById()
}
