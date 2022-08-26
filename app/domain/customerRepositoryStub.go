package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Jorge", City: "São Paulo", Zipcode: "06720450", DateofBirth: "1985-04-12", Status: "1"},
		{Id: "1002", Name: "Eliete", City: "Pernambuco", Zipcode: "12345678", DateofBirth: "1971-10-28", Status: "1"},
	}
	return CustomerRepositoryStub{
		customers: customers,
	}
}
