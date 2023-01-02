package user

type UserObj struct {
	ID                string `json:"id"`
	CollegeID         string `json:"collegeId"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phoneNumber"`
	BirthDate         string `json:"birthDate"`
	ProfileImageUrl   string `json:"profileImageUrl"`
	EncryptedPassword string `json:"encryptedPassword"`
	IsAdmin           bool   `json:"isAdmin"`
}

type UserReward struct {
	UserID    string      `json:"userId"`
	CollegeID string      `json:"collegeId"`
	Rewards   []RewardObj `json:"rewards"`
}

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