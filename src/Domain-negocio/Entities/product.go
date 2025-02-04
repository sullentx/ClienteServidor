package entities

//SRP, Single Responsability Principle (Principio de responsabilidad única)
type Product struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Quantity     int    `json:"quantity"`
	CodigoBarras string `json:"codigo_barras"`
}

func NewProduct(Name string, Price float64, Quantity int, CodigoBarras string) *Product {
	return &Product{Name: Name, Quantity: Quantity, CodigoBarras: CodigoBarras}
}
