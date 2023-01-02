package session

import "college-learning-asynchronous-gamification/accessor"

type Service interface {
	UpsertSession(req UpsertSessionReq, res *UpsertSessionRes) error
	DeleteSession(req DeleteSessionReq, res *DeleteSessionRes) error
	ListSession(req ListSessionReq, res *ListSessionRes) error
	GetSessionByIDs(req GetSessionByIDsReq, res *GetSessionByIDsRes) error
}

type service struct {
	accessor accessor.Accessor
}

func NewService(accessor accessor.Accessor) Service {
	return &service{
		accessor: accessor,
	}
}

type ListSessionReq struct {
	CollegeID string `json:"collegeId"`
}

type ListSessionRes struct {
	SessionIDs []string `json:"sessionIds"`
}

type GetSessionByIDsReq struct {
	SessionIDs []string `json:"sessionIds"`
}

type GetSessionByIDsRes struct {
	Sessions []SessionObj `json:"sessions"`
}

type UpsertSessionReq struct {
	Session SessionObj `json:"session"`
}

type UpsertSessionRes struct {}

type DeleteSessionReq struct {
	SessionID string `json:"sessionId"`
}

type DeleteSessionRes struct {}
