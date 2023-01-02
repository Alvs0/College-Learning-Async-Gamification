package session

import "college-learning-asynchronous-gamification/accessor"

func (c *service) ListSession(req ListSessionReq, res *ListSessionRes) error {
	sessionIDs, err := c.accessor.ListSession(accessor.SessionListInput{
		CollegeID: req.CollegeID,
	})
	if err != nil {
		return err
	}

	*res = ListSessionRes{
		SessionIDs: sessionIDs,
	}

	return nil
}
