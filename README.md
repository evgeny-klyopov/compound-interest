# compound-interest

### Install
```bash
go get "github.com/evgeny-klyopov/compound-interest"
```

### Usage
```go
import (
    "fmt"
    "github.com/evgeny-klyopov/compound-interest"
    "time"
)
func main() {
    date, _ := time.Parse("2006-01-02", "2020-12-12")
    
    for monthNumber, row := range compoundInterest.New(compoundInterest.Params{
    	DateStart:                                date,
    	InvestmentTermInYears:                    2,
    	PercentRate:                              10,
    	InitialPayment:                           50000,
    	MonthlyPayment:                           30000,
    	AnnualPercentageIncreaseInMonthlyPayment: 15,
    	AvgPercentDividend:                       4,
        AnnualPercentageIncreaseDividend:         5,
    }).Calculate() {
        fmt.Printf(
            "monthNumber = %v, " +
                "date = %s, " +
                "monthlyPayment = %f, " +
                "AvgPercentDividend = %f, " +
                "Amount = %f, " +
                "Dividend = %f\n",
            monthNumber + 1,
            row.Date.Format("2006 Jan"),
            row.MonthlyPayment,
            row.AvgPercentDividend,
            row.Amount,
            row.MonthlyDividend,
        )
    }       
}
```






