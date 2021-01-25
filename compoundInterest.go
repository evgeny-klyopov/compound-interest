package compoundInterest

import (
	"math"
	"time"
)

const numberOfMonthsInYear float64 = 12

type compoundInterest struct {
	params            Params
	countMonthDeposit int
	percentMonth      float64
	coefficient       float64
}

type Predictor interface {
	Calculate() []Prediction

	getDate(month int) time.Time
	getAmount(month int) float64
	getPercent(amount float64, percent float64) float64
	getPrediction(month int) Prediction
	updateParamsWithInflation()
}

func New(params Params) Predictor {
	var p Predictor

	percentMonth := (params.PercentRate / 100) / numberOfMonthsInYear

	p = &compoundInterest{
		params:            params,
		countMonthDeposit: int(params.InvestmentTermInYears * numberOfMonthsInYear),
		percentMonth:      percentMonth,
		coefficient:       1 + percentMonth,
	}

	return p
}

func (c *compoundInterest) Calculate() []Prediction {
	result := make([]Prediction, c.countMonthDeposit, c.countMonthDeposit)
	for month := 1; month <= c.countMonthDeposit; month++ {
		result[month-1] = c.getPrediction(month)

		if month%int(numberOfMonthsInYear) == 0 {
			c.updateParamsWithInflation()
		}
	}

	return result
}

func (c *compoundInterest) getDate(month int) time.Time {
	return c.params.DateStart.AddDate(0, month, 0)
}

func (c *compoundInterest) getAmount(month int) float64 {
	return c.params.InitialPayment*
		math.Pow(c.coefficient, float64(month)) +
		(c.params.MonthlyPayment*math.Pow(c.coefficient, float64(month))*
			c.coefficient-c.params.MonthlyPayment*c.coefficient)/
			c.percentMonth
}

func (c *compoundInterest) getPercent(amount float64, percent float64) float64 {
	return (amount / 100) * percent
}

func (c *compoundInterest) getPrediction(month int) Prediction {
	var prediction Prediction
	prediction.Date = c.getDate(month)
	prediction.Amount = c.getAmount(month)
	prediction.MonthlyDividend = c.getPercent(prediction.Amount, c.params.AvgPercentDividend) / numberOfMonthsInYear
	prediction.MonthlyPayment = c.params.MonthlyPayment
	prediction.AvgPercentDividend = c.params.AvgPercentDividend

	return prediction
}

func (c *compoundInterest) updateParamsWithInflation() {
	c.params.MonthlyPayment = c.getPercent(c.params.MonthlyPayment, c.params.AnnualPercentageIncreaseInMonthlyPayment) +
		c.params.MonthlyPayment
	c.params.AvgPercentDividend = c.getPercent(c.params.AvgPercentDividend, c.params.AnnualPercentageIncreaseDividend) +
		c.params.AvgPercentDividend
}
