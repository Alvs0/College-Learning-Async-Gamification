package college

import (
	"github.com/google/uuid"
	"college-learning-asynchronous-gamification/accessor"
)

func (c *service) UpsertCollege(req UpsertCollegeReq, res *UpsertCollegeRes) error {
	id := uuid.New()
	upsertInput := accessor.CollegeDb{
		ID:      id.String(),
		Name:    req.Name,
		Address: req.Address,
		IconURL: req.ImageUrl,
	}

	if err := c.accessor.UpsertCollege(upsertInput); err != nil {
		return err
	}

	return nil
}
