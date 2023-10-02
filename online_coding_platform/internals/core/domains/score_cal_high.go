package domains

type ScoreCalHigh struct {
	offset int64
}

func NewScoreCalHigh() *ScoreCalHigh {
	return &ScoreCalHigh{0}
}

func (cal *ScoreCalHigh) CalculateScore(ques []Question) int64 {
	curSum := int64(0)
	for _, curQues := range ques {
		curSum += curQues.points
	}
	return int64(curSum - cal.offset)
}
