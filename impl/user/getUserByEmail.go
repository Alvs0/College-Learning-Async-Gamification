package user

import (
	"log"
)

func (s *service) GetUserByEmail(req GetUserByEmailReq, res *GetUserByEmailRes) error {
	userDbs, err := s.accessor.GetUserByEmail(req.Email)
	if err != nil {
		log.Println("[GetUserByEmail] error getting user by email. cause: %v ", err.Error())
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
			EncryptedPassword: userDb.Password,
			IsAdmin:           userDb.IsAdmin,
		})
	}

	*res = GetUserByEmailRes{Results: userObjs}

	return nil
}
