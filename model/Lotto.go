package model

import "time"

type Lottos struct {
	ID           uint `gorm:"primary_key"`
	Name         string
	Number_lotto int64
	Multiply     int64
	Price        float64
	CreatedAt    time.Time
}
