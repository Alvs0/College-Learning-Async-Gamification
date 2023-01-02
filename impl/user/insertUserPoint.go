package user

import (
	"college-learning-asynchronous-gamification/accessor"
)

func (s *service) InsertUserPoint(req InsertUserPointReq, res *InsertUserPointRes) error {
	upsertInput := accessor.UserPointDb{
		UserID:    req.UserID,
		CollegeID: req.CollegeID,
		Point:     req.Point,
	}

	if err := s.accessor.InsertUserPoint(upsertInput); err != nil {
		return err
	}

	*res = InsertUserPointRes{}

	return nil
}
