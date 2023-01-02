package user

import (
	"log"
)

func (s *service) DeleteUser(req DeleteUserReq, res *DeleteUserRes) error {
	err := s.accessor.DeleteUser(req.ID)
	if err != nil {
		log.Println("[DeleteUser] error getting user by email. cause: %v ", err.Error())
	}

	*res = DeleteUserRes{}

	return nil
}
