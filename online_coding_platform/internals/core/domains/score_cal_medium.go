package domains

type ScoreCalMedium struct {
	offset int64
}

func NewScoreCalMedium() *ScoreCalMedium {
	return &ScoreCalMedium{30}
}

func (cal *ScoreCalMedium) CalculateScore(ques []Question) int64 {
	curSum := int64(0)
	for _, curQues := range ques {
		curSum += curQues.points
	}
	return int64(curSum - cal.offset)
}
