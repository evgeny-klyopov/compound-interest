package compoundInterest

import "time"

type Prediction struct {
	Date               time.Time
	Amount             float64
	MonthlyDividend    float64
	MonthlyPayment     float64
	AvgPercentDividend float64
}
