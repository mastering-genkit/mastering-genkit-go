package tools

// SearchRecipeInput represents the input for recipe search tool
type SearchRecipeInput struct {
	Ingredients []string `json:"ingredients" jsonschema_description:"List of ingredients to search for"`
	Tags        []string `json:"tags,omitempty" jsonschema_description:"Recipe tags to filter by"`
	Category    string   `json:"category,omitempty" jsonschema_description:"Recipe category (Italian, Japanese, etc.)"`
	MaxResults  int      `json:"maxResults,omitempty" jsonschema_description:"Maximum number of results to return"`
}