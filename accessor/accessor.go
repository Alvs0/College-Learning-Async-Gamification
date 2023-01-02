package accessor

type Accessor interface {
	User
	Reward
	College
	Session
	UserPoint
	UserReward
}

type accessor struct {
	sqlAdapter SqlAdapter
}

func NewAccessor(sqlAdapter SqlAdapter) Accessor {
	return &accessor{
		sqlAdapter: sqlAdapter,
	}
}
