package user

import (
	"log"
)

func (s *service) GetUserByIDs(req GetUserByIDsReq, res *GetUserByIDsRes) error {
	if req.UserIDs == nil {
		return nil
	}

	userDbs, err := s.accessor.GetUsers(req.UserIDs)
	if err != nil {
		log.Println("[GetUserByIDs] error getting user by email. cause: %v ", err.Error())
	}

	var userObjs []UserObj
	for _, userDb := range userDbs {
		userObjs = append(userObjs, UserObj{
			ID:                userDb.ID,
			CollegeID:         userDb.CollegeID,
			Name:              userDb.Name,
			Email:             userDb.Email,
			PhoneNumber:       userDb.PhoneNumber,
			BirthDate:         userDb.BirthDate,
			ProfileImageUrl:   userDb.ProfileImageUrl,
			IsAdmin:           userDb.IsAdmin,
		})
	}

	*res = GetUserByIDsRes{Users: userObjs}

	return nil
}
