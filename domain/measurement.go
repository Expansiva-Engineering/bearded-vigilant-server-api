package domain

import "time"

type Measurement struct {
	ID                  uint64            `sql:"AUTO_INCREMENT; NOT NULL"; json:"id"`
	Value               float64           `sql:"NOT NULL"; json:"value"`
	BinSize             uint16            `sql:"NOT NULL"; json:"binSize"`
	StartTimestamp      time.Time         `sql:"NOT NULL"; json:"startTimestamp"`
	EndTimestamp        *time.Time        `json:"endTimestamp"`
	Latitude            float32           `json:"latitude"`
	Longitude           float32           `json:"longitude"`
	MeasurementSeriesID uint64            `gorm:"column:measurement_series_id"`
	MeasurementSeries   MeasurementSeries `gorm:"foreignkey:MeasurementSeriesId"`
}
