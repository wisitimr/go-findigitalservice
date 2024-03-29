package repository

type Repository struct {
	User           UserRepository
	Account        AccountRepository
	ForwardAccount ForwardAccountRepository
	Supplier       SupplierRepository
	Customer       CustomerRepository
	Document       DocumentRepository
	PaymentMethod  PaymentMethodRepository
	Product        ProductRepository
	Company        CompanyRepository
	Daybook        DaybookRepository
	DaybookDetail  DaybookDetailRepository
	Role           RoleRepository
	Material       MaterialRepository
}
