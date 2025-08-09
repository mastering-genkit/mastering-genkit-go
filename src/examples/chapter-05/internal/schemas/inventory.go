package schemas

// InventoryCheckInput represents the input for inventory checking
type InventoryCheckInput struct {
	Items []InventoryItem `json:"items"`
}

// InventoryItem represents a single item to check
type InventoryItem struct {
	SKU      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

// InventoryCheckResult represents the structured output from inventory check
type InventoryCheckResult struct {
	Available     bool           `json:"available"`
	ItemsInStock  map[string]int `json:"items_in_stock"`
	ReservedUntil string         `json:"reserved_until"`
	Warehouse     string         `json:"warehouse"`
}