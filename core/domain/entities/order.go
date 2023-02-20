package entities

type Order struct {
	ID            string
	InvoiceNumber string
	SubtotalPrice float64
	TotalPrice    float64
}

func NewOrder(id, invoiceNumber string, subtotalPrice, totalPrice float64) *Order {
	return &Order{
		ID:            id,
		InvoiceNumber: invoiceNumber,
		SubtotalPrice: subtotalPrice,
		TotalPrice:    totalPrice,
	}
}
