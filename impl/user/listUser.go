package user

import (
	"college-learning-asynchronous-gamification/accessor"
	"log"
)

func (s *service) ListUser(req ListUserReq, res *ListUserRes) error {
	userIds, err := s.accessor.ListUser(accessor.UserListInput{
		CollegeID: req.CollegeID,
		IsAdmin:   req.IsAdmin,
	})
	if err != nil {
		log.Println("[ListUser] error getting user by email. cause: %v ", err.Error())
	}

	*res = ListUserRes{
		UserIDs: userIds,
	}

	return nil
}
