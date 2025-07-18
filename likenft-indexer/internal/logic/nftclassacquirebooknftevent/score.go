package nftclassacquirebooknftevent

import (
	"time"
)

type CalculateScoreFn func(
	lastProcessedTime time.Time,
	blockHeight uint64,
	weight float64,
) float64

type CalculateScoreFnFactory interface {
	CalculateScoreFn(
		lastProcessedTime time.Time,
		blockHeight uint64,
		weight float64,
	) float64
}

type calculateScoreFnFactory struct {
	blockHeightContribution float64
	weight0Constant         float64
	weight1Constant         float64
	weightContribution      float64
}

// Calculate score based on the following formula
//
//	score = time + weight * weightContribution
//
// The weight coeff is interpolated by weight0Constant and weight1Constant,
// And finally multiplied by weightContribution to contribute to the final score
//
// Adjust accordingly with respect to the scheduler config such that the score
// can be consumed by the desire frequency.
//
// e.g. with scheduler cron * * * * *
//
//	weight0Constant = 10 // minutes
//	weight1Constant = 1 // minutes
//	weightContribution = 60 // 1 minute
//
// The jobs are expected to have a diff of t + 60 for weight = 1, and a diff of t + 600 for weight = 0
// when the score is re-calculated.
func MakeCalculateScoreFn(
	blockHeightContribution float64,
	weight0Constant float64,
	weight1Constant float64,
	weightContribution float64,
) CalculateScoreFn {
	factory := &calculateScoreFnFactory{
		blockHeightContribution,
		weight0Constant,
		weight1Constant,
		weightContribution,
	}
	return factory.CalculateScoreFn
}

func (f *calculateScoreFnFactory) CalculateScoreFn(
	lastProcessedTime time.Time,
	blockHeight uint64,
	weight float64,
) float64 {
	// ↓ last_processed_time -> ↓ score
	// ↓ block_height -> ↓ score

	// As item of higher weight will be distributed more compatly
	// The slope (rate of change with respect of time delta) should decrease while weight increase
	// i.e. w=0 => high, w=1 => low
	lineOfWeight := func(w float64) float64 {
		return (1-w)*f.weight0Constant + w*f.weight1Constant
	}

	return float64(blockHeight)*f.blockHeightContribution +
		float64(lastProcessedTime.Unix()) +
		lineOfWeight(weight)*f.weightContribution
}
