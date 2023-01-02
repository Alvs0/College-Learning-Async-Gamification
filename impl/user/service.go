package user

import "college-learning-asynchronous-gamification/accessor"

type Service interface {
	ListUser(req ListUserReq, res *ListUserRes) error
	ListAndGetUserReward(req ListAndGetUserRewardReq, res *ListAndGetUserRewardRes) error
	GetUserByIDs(req GetUserByIDsReq, res *GetUserByIDsRes) error
	GetUserByEmail(req GetUserByEmailReq, res *GetUserByEmailRes) error
	UpsertUser(req UpsertUserReq, res *UpsertUserRes) error
	DeleteUser(req DeleteUserReq, res *DeleteUserRes) error
	InsertUserPoint(req InsertUserPointReq, res *InsertUserPointRes) error
	UpdateUserPoint(req UpdateUserPointReq, res *UpdateUserPointRes) error
}

type service struct {
	accessor accessor.Accessor
}

func NewService(accessor accessor.Accessor) Service {
	return &service{
		accessor: accessor,
	}
}

type GetUserByEmailReq struct {
	Email string `json:"email"`
}

type GetUserByEmailRes struct {
	Results []UserObj `json:"results"`
}

type ListUserReq struct {
	CollegeID string `json:"collegeId"`
	IsAdmin   bool   `json:"isAdmin"`
}

type ListUserRes struct {
	UserIDs []string `json:"userIds"`
}

type GetUserByIDsReq struct {
	UserIDs []string `json:"userIds"`
}

type GetUserByIDsRes struct {
	Users []UserObj `json:"users"`
}

type UpsertUserReq struct {
	User UserObj `json:"user"`
}

type UpsertUserRes struct{}

type DeleteUserReq struct {
	ID string `json:"id"`
}

type DeleteUserRes struct{}

type InsertUserPointReq struct {
	UserID    string `json:"userId"`
	CollegeID string `json:"collegeId"`
	Point     int    `json:"point"`
}

type InsertUserPointRes struct{}

type ListAndGetUserRewardReq struct {
	UserID    string `json:"userId"`
	CollegeID string `json:"collegeId"`
}

type ListAndGetUserRewardRes struct {
	UserReward UserReward `json:"userReward"`
}

type UpdateUserPointReq struct {
	UserID    string `json:"userId"`
	CollegeID string `json:"collegeId"`
	Point     int    `json:"point"`
}

type UpdateUserPointRes struct{}
