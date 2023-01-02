package college

import "college-learning-asynchronous-gamification/accessor"

type Service interface {
	UpsertCollege(req UpsertCollegeReq, res *UpsertCollegeRes) error
	ListCollege(req ListCollegeReq, res *ListCollegeRes) error
	GetCollegeByID(req GetCollegeByIDReq, res *GetCollegeByIDRes) error
	GetCollegeByIDs(req GetCollegeByIDsReq, res *GetCollegeByIDsRes) error
	DeleteCollege(req DeleteCollegeReq, res *DeleteCollegeRes) error
	GetCollegeByName(req GetCollegeByNameReq, res *GetCollegeByNameRes) error
}

type service struct {
	accessor accessor.Accessor
}

func NewService(accessor accessor.Accessor) Service {
	return &service{
		accessor: accessor,
	}
}

type UpsertCollegeReq struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	ImageUrl string `json:"imageUrl"`
}

type UpsertCollegeRes struct{}

type ListCollegeReq struct{}

type ListCollegeRes struct {
	CollegeIDs []string `json:"collegeIds"`
}

type GetCollegeByIDReq struct {
	CollegeID string `json:"collegeId"`
}

type GetCollegeByIDRes struct {
	College College `json:"college"`
}

type GetCollegeByIDsReq struct {
	CollegeIDs []string `json:"collegeIds"`
}

type GetCollegeByIDsRes struct {
	Colleges []College `json:"colleges"`
}

type DeleteCollegeReq struct {
	ID string `json:"id"`
}

type DeleteCollegeRes struct{}

type GetCollegeByNameReq struct {
	Name string `json:"name"`
}

type GetCollegeByNameRes struct {
	College College `json:"college"`
}
