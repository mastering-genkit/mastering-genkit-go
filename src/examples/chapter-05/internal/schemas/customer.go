package schemas

// CustomerValidationInput represents the input for customer validation
type CustomerValidationInput struct {
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Age         int     `json:"age"`
	CreditLimit float64 `json:"credit_limit"`
}

// CustomerValidationResult represents the structured output from customer validation
type CustomerValidationResult struct {
	CustomerID   string   `json:"customer_id"`
	IsValid      bool     `json:"is_valid"`
	CreditScore  int      `json:"credit_score"`
	AccountTier  string   `json:"account_tier"`
	Restrictions []string `json:"restrictions,omitempty"`
}