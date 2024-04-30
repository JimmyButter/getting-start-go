package model

import "time"

type InstPoint struct {
	InstId     string
	Point      int64
	CtVal      float32
	TickSz     float32
	Level      int8
	Status     string
	Scale      int8
	Avg        float32
	AvgVol     float32
	UpdateTime time.Time
	Remark     string
	Tag        string
}
