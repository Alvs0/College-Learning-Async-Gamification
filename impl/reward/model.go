package reward

type RewardObj struct {
	ID             string `json:"id"`
	CollegeID      string `json:"collegeId"`
	Name           string `json:"name"`
	ImageUrl       string `json:"imageUrl"`
	Description    string `json:"description"`
	Quantity       int    `json:"quantity"`
	RequiredPoints int    `json:"requiredPoints"`
	MinimalLevel   int    `json:"minimalLevel"`
	IsActive       bool   `json:"isActive"`
}
