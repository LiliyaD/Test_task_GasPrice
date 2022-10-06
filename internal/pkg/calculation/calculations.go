package calculation

import (
	"sync"
	"time"

	"github.com/LiliyaD/Test_task_GasPrice/internal/journal"
	"github.com/LiliyaD/Test_task_GasPrice/internal/pkg/models"
)

func Process() (responce models.ResponceJSON) {
	var wg sync.WaitGroup
	wg.Add(4)

	go CalcWholePeriodPrice(&responce, &wg)

	go CalcMonthlyGas(&responce, &wg)

	go CalcDailyGas(&responce, &wg)

	go CalcFreqHourPriceDistr(&responce, &wg)

	wg.Wait()

	return
}

// Сколько было потрачено gas помесячно
func CalcMonthlyGas(responce *models.ResponceJSON, wg *sync.WaitGroup) {
	i := -1
	year, month := 0, time.Month(0)
	for _, v := range source.Ethereum.Transactions {
		dateTime, err := time.Parse("06-01-02 15:04", v.Time)
		if err != nil {
			journal.LogFatal(err)
		}

		if (year != dateTime.Year() || month != dateTime.Month()) || (year == 0) {
			year, month = dateTime.Year(), dateTime.Month()
			i++
			responce.MonthlyGas = append(responce.MonthlyGas, models.MonthlyGasF{
				YearMonth: v.Time[0:5],
				SpentGas:  0,
			})
		}
		responce.MonthlyGas[i].SpentGas += v.GasValue
	}
	wg.Done()
}

// Средняя цена gas за день
func CalcDailyGas(responce *models.ResponceJSON, wg *sync.WaitGroup) {
	i, daysNumb, sum := -1, 0, float64(0)
	year, month, day := 0, time.Month(0), 0
	len := len(source.Ethereum.Transactions) - 1
	for n, v := range source.Ethereum.Transactions {
		dateTime, err := time.Parse("06-01-02 15:04", v.Time)
		if err != nil {
			journal.LogFatal(err)
		}

		if year != dateTime.Year() || month != dateTime.Month() || day != dateTime.Day() {
			if i >= 0 {
				responce.DailyGasPrice[i].AveragePrice = sum / float64(daysNumb)
			}

			year, month, day = dateTime.Year(), dateTime.Month(), dateTime.Day()
			i++
			daysNumb, sum = 0, float64(0)
			responce.DailyGasPrice = append(responce.DailyGasPrice,
				models.DailyGasF{
					Date:         v.Time[0:8],
					AveragePrice: 0,
				})
		}

		daysNumb++
		sum += v.GasPrice

		if n == len {
			responce.DailyGasPrice[i].AveragePrice = sum / float64(daysNumb)
		}
	}
	wg.Done()
}

// Частотное распределение цены по часам
func CalcFreqHourPriceDistr(responce *models.ResponceJSON, wg *sync.WaitGroup) {
	for _, v := range source.Ethereum.Transactions {
		responce.FreqHourPriceDistr = append(responce.FreqHourPriceDistr,
			models.HourlyGasF{
				DateTime: v.Time,
				Price:    v.GasPrice,
			})
	}
	wg.Done()
}

// Сколько заплатили за весь период
func CalcWholePeriodPrice(responce *models.ResponceJSON, wg *sync.WaitGroup) {
	var sum float64
	for _, v := range source.Ethereum.Transactions {
		sum += v.GasPrice * v.GasValue
	}
	responce.WholePeriodPrice = sum
	wg.Done()
}
