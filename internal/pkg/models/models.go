package models

// Source
type SourceJSON struct {
	Ethereum SourceEthereum `json:"ethereum"`
}

type SourceEthereum struct {
	Transactions []SourceTransaction `json:"transactions"`
}

type SourceTransaction struct {
	Time           string  `json:"time"`
	GasPrice       float64 `json:"gasPrice"`
	GasValue       float64 `json:"gasValue"`
	Average        float64 `json:"average"`
	MaxGasPrice    float64 `json:"maxGasPrice"`
	MedianGasPrice float64 `json:"medianGasPrice"`
}

// Response
type ResponceJSON struct {
	MonthlyGas         []MonthlyGasF `json:"MonthlyGas"`         // Сколько было потрачено gas помесячно
	DailyGasPrice      []DailyGasF   `json:"DailyGasPrice"`      // Среднюю цену gas за день
	FreqHourPriceDistr []HourlyGasF  `json:"FreqHourPriceDistr"` // Частотное распределение цены по часам(за весь период).
	WholePeriodPrice   float64       `json:"WholePeriodPrice"`   // Сколько заплатили за весь период
}

type MonthlyGasF struct {
	YearMonth string  `json:"YearMonth"` // год-месяц
	SpentGas  float64 `json:"SpentGas"`  // потраченный газ за месяц
}

type DailyGasF struct {
	Date         string  `json:"YearMonthDay"` // год-месяц-день
	AveragePrice float64 `json:"AveragePrice"` // средняя цена газа за день
}

type HourlyGasF struct {
	DateTime string  `json:"DateTime"` // // год-месяц-день час:00
	Price    float64 `json:"Price"`    // цена газа за этот час
}
