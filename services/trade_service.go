package services

import (
	"context"
	"go-gin-udemy-serverside-dev-2/domain"
	"time"
)

type Clock interface {
	Now() time.Time
}

type systemClock struct{}

func (systemClock) Now() time.Time { return time.Now() }

type ITradeService interface {
	Create(ctx context.Context, t domain.Trade) (domain.TradeID, error)
	GetByID(ctx context.Context, id domain.TradeID) (*domain.Trade, error)
}

type TradeSevice struct {
	repo  domain.TradeRepository
	clock Clock
}

func NewTradeService(r domain.TradeRepository, c Clock) ITradeService {
	if c == nil {
		c = systemClock{}
	}
	return &TradeSevice{repo: r, clock: c}
}

func (s *TradeSevice) Create(ctx context.Context, t domain.Trade) (domain.TradeID, error) {
	if err := t.Validate(); err != nil {
		return 0, err
	}
	t.BookedAt = s.clock.Now()
	return s.repo.Save(ctx, t)
}

func (s *TradeSevice) GetByID(ctx context.Context, id domain.TradeID) (*domain.Trade, error) {
	return s.repo.FindByID(ctx, id)
}