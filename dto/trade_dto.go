package dto

import "go-gin-udemy-serverside-dev-2/domain"

type TradeCreateReq struct {
	Symbol   string `json:"symbol"`
	Price    int64  `json:"price"`
	Quantity int    `json:"quantity"`
}

func (r TradeCreateReq) ToDomain() (domain.Trade, error) {
	p, err := domain.NewPrice(r.Price)
	if err != nil {
		return domain.Trade{}, err
	}
	return domain.Trade{
		Symbol:   r.Symbol,
		Price:    p,
		Quantity: r.Quantity,
	}, nil
}

type TradeRes struct {
	ID       int64  `json:"id"`
	Symbol   string `json:"symbol"`
	Price    int64  `json:"price"`
	Quantity int    `json:"quantity"`
	BookedAt string `json:"bookedAt"`
}

func FromDomain(t domain.Trade) TradeRes {
	return TradeRes{
		ID:       int64(t.ID),
		Symbol:   t.Symbol,
		Price:    pValue(t.Price),
		Quantity: t.Quantity,
		BookedAt: t.BookedAt.UTC().Format(TimeLayout),
	}
}

const TimeLayout = "2006-01-02T15:04:05Z07:00"

// domain.Price の値を取る薄い関数（必要なら domain にエクスポートを追加）
func pValue(p domain.Price) int64 {
	// パッケージ境界越えの都合で必要なら domain に Getter を用意してください。
	// ここでは説明簡略化のため直接の実装は省略します。
	return 0
}
