package models

import "time"

type HistoricalDataPoint struct {
	Symbol string
	Date   time.Time
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume int
}
