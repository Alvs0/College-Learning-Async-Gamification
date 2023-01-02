package user

import (
	"college-learning-asynchronous-gamification/accessor"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) UpsertUser(req UpsertUserReq, res *UpsertUserRes) error {
	id := uuid.New()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.User.EncryptedPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	upsertInput := accessor.UserUpsertInput{
		ID:              id.String(),
		CollegeID:       req.User.CollegeID,
		Name:            req.User.Name,
		Email:           req.User.Email,
		PhoneNumber:     req.User.PhoneNumber,
		BirthDate:       req.User.BirthDate,
		ProfileImageUrl: req.User.ProfileImageUrl,
		Password:        string(hashedPassword),
		IsAdmin:         req.User.IsAdmin,
	}

	if err := s.accessor.UpsertUser(upsertInput); err != nil {
		return err
	}

	*res = UpsertUserRes{}

	return nil
}
