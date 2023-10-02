package domains

type ScoreCalLow struct {
	lowOffSet int64
}

func NewScoreCalLow() *ScoreCalLow {
	return &ScoreCalLow{50}
}

func (cal *ScoreCalLow) CalculateScore(ques []Question) int64 {
	curSum := int64(0)
	for _, curQues := range ques {
		curSum += curQues.points
	}
	return int64(curSum - cal.lowOffSet)
}
