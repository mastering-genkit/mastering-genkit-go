package schemas

// InvoiceData represents structured invoice information
type InvoiceData struct {
	InvoiceNumber string  `json:"invoice_number"`
	Amount        float64 `json:"amount"`
	DueDate       string  `json:"due_date"`
	CustomerName  string  `json:"customer_name"`
	Items         []Item  `json:"items"`
}

// Item represents a line item in an invoice
type Item struct {
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	Total       float64 `json:"total"`
}