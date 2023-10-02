package domains

import (
	custom_enums "github.com/online_coding_platform/internals/core/domains/enums"
	custom_errors "github.com/online_coding_platform/internals/core/errors"
)

func ScoreCalculatorFactory(level int64) (ScoreCalculator, error) {
	var (
		cal ScoreCalculator
		err error
	)
	switch level {
	case custom_enums.Low:
		cal = NewScoreCalLow()
	case custom_enums.Medium:
		cal = NewScoreCalMedium()
	case custom_enums.High:
		cal = NewScoreCalHigh()
	default:
		err = custom_errors.ScoreCalNotDefined
	}
	return cal, err
}
