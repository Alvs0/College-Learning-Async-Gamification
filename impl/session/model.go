package session

type SessionObj struct {
	ID        string `json:"id"`
	CollegeID string `json:"collegeId"`
	Name      string `json:"name"`
	Link      string `json:"link"`
	ImageUrl  string `json:"imageUrl"`
}
