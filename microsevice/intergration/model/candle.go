package model

type Candle struct {
	Id      int64
	InstId  string
	Bar     string
	O       float32
	H       float32
	L       float32
	C       float32
	Vol     int64
	Confirm string
	Ts      int64
}
