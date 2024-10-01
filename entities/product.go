package entities

type Product struct {
	SKU   string
	Name  string
	Qty   int
	Price float64
}

type Order struct {
	OrderID int
	SKU     string
	Qty     int
}

type ProductOrder struct {
	SKU string
	Qty int
}
