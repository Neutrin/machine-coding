package domains

type ScoreCalculator interface {
	CalculateScore(ques []Question) int64
}
