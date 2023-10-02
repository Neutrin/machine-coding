package domains

type LeaderBoard struct {
	Name  string
	Score int64
}

func NewLeaderBoard(user User) LeaderBoard {
	return LeaderBoard{
		Name:  user.Name(),
		Score: user.Score(),
	}
}
