package domain

import (
	"context"
	"sync"
)

type ITradeRepository interface {
	Save(ctx context.Context, t Trade) (TradeID, error)
	FindByID(ctx context.Context, id TradeID) (*Trade, error)
}

type TradeRepoInMem struct {
	mu    sync.RWMutex
	seq   int64
	store map[TradeID]Trade
}



func NewTradeRepoInMem() ITradeRepository {
	return &TradeRepoInMem{store: make(map[TradeID]Trade)}
}

func (r *TradeRepoInMem) FindByID(ctx context.Context, id TradeID) (*Trade, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if v, ok := r.store[id]; ok {
		c := v // コピー
		return &c, nil
	}
	return nil, ErrNotFound
}

func (r *TradeRepoInMem) Save(ctx context.Context, t Trade) (TradeID, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if t.ID == 0 {
		r.seq++
		t.ID = TradeID(r.seq)
	}
	r.store[t.ID] = t
	return t.ID, nil
}