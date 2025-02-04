package entities

//SRP, Single Responsability Principle (Principio de responsabilidad única)
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func NewProduct(Name string, Price float64, Quantity int) *Product {
	return &Product{Name: Name, Price: Price, Quantity: Quantity}
}
