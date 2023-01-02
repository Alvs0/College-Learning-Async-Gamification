package user

import "college-learning-asynchronous-gamification/accessor"

func (s *service) UpdateUserPoint(req UpdateUserPointReq, res *UpdateUserPointRes) error {
	err := s.accessor.UpdateUserPoint(accessor.UserPointDb{
		UserID:    req.UserID,
		CollegeID: req.CollegeID,
		Point:     req.Point,
	})
	if err != nil {
		return err
	}

	return nil
}
