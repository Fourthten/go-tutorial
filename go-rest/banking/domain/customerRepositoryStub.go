package domain

type CustomerRepositoryStub struct {
    customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
    return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
    customers := []Customer {
        {"1001", "Agung", "Jawa Barat", "17519", "1990-01-01", "1"},
        {"1002", "Setia", "Jawa Tengah", "15421", "1990-01-01", "1"},
    }
    return CustomerRepositoryStub{customers}
}