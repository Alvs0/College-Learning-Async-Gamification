package session

import (
	"college-learning-asynchronous-gamification/accessor"
	"github.com/google/uuid"
)

func (c *service) UpsertSession(req UpsertSessionReq, res *UpsertSessionRes) error {
	id := uuid.New()
	upsertInput := accessor.SessionDb{
		ID:        id.String(),
		CollegeID: req.Session.CollegeID,
		Name:      req.Session.Name,
		Link:      req.Session.Link,
		ImageUrl:  req.Session.ImageUrl,
	}

	if err := c.accessor.UpsertSession(upsertInput); err != nil {
		return err
	}

	return nil
}
