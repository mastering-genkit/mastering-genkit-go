package schemas

// OrderCreationInput represents the input for order creation
type OrderCreationInput struct {
	CustomerID  string   `json:"customer_id"`
	Items       []string `json:"items"`
	TotalAmount float64  `json:"total_amount"`
}

// OrderResult represents the structured output from order creation
type OrderResult struct {
	OrderID           string  `json:"order_id"`
	Status            string  `json:"status"`
	EstimatedDelivery string  `json:"estimated_delivery"`
	TrackingNumber    string  `json:"tracking_number"`
	TotalAmount       float64 `json:"total_amount"`
	PaymentMethod     string  `json:"payment_method"`
}