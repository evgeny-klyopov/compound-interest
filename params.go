package compoundInterest

import "time"

type Params struct {
	DateStart                                time.Time `json:"date_start"`
	InvestmentTermInYears                    float64   `json:"investment_term_in_years"`
	PercentRate                              float64   `json:"percent_rate"`
	InitialPayment                           float64   `json:"initial_payment"`
	MonthlyPayment                           float64   `json:"monthly_payment"`
	AnnualPercentageIncreaseInMonthlyPayment float64   `json:"annual_percentage_increase_in_monthly_payment"`
	AvgPercentDividend                       float64   `json:"avg_percent_dividend"`
	AnnualPercentageIncreaseDividend         float64   `json:"annual_percentage_increase_dividend"`
}
