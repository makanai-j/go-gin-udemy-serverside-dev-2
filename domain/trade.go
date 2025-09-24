package domain

import (
	"errors"
	"time"
)

var (
	ErrInvalid = errors.New("invalid parameter")
	ErrNotFound = errors.New("trade not found")
)

type TradeID int64

type Price struct {
	value int64
}

func NewPrice(v int64) (Price, error) {
	if v <= 0 {
		return Price{}, ErrInvalid
	}

	return Price{value: v}, nil
}

type Trade struct {
	ID       TradeID
	Symbol   string
	Price    Price
	Quantity int
	BookedAt time.Time
}

func (t Trade) Validate() error {
	if t.Symbol == "" || t.Quantity <= 0 || t.Price.value <= 0 {
		return ErrInvalid
	}
	return nil
}

func (t Trade) PnL(current Price) int64 {
	return (current.value - t.Price.value) * int64(t.Quantity)
}