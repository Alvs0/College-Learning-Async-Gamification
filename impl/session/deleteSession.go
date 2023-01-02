package session

func (c *service) DeleteSession(req DeleteSessionReq, res *DeleteSessionRes) error {
	if err := c.accessor.DeleteSession(req.SessionID); err != nil {
		return err
	}

	return nil
}
