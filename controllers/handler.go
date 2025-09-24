package controllers

import (
	"encoding/json"
	"errors"
	"go-gin-udemy-serverside-dev-2/domain"
	"go-gin-udemy-serverside-dev-2/dto"
	"go-gin-udemy-serverside-dev-2/services"
	"net/http"
	"strconv"
	"strings"
)

type TradeHandler struct {
	svc services.ITradeService
}

func NewTradeHandler(svc services.ITradeService) *TradeHandler {
	return &TradeHandler{svc: svc}
}

// POST /trades
func (h *TradeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.TradeCreateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	t, err := req.ToDomain()
	if err != nil {
		status := http.StatusBadRequest
		if errors.Is(err, domain.ErrInvalid) {
			status = http.StatusBadRequest
		}
		writeError(w, status, err.Error())
		return
	}
	id, err := h.svc.Create(r.Context(), t)
	if err != nil {
		writeDomainError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]any{"id": id})
}

// GET /trades/{id}
func (h *TradeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/trades/")
	n, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || n <= 0 {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	t, err := h.svc.GetByID(r.Context(), domain.TradeID(n))
	if err != nil {
		writeDomainError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(dto.TradeRes{
		ID:       int64(t.ID),
		Symbol:   t.Symbol,
		Price:    t.Price // ← Getter を用意して使う
		Quantity: t.Quantity,
		BookedAt: t.BookedAt.UTC().Format(timeLayout),
	})
}

func writeDomainError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalid):
		writeError(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, domain.ErrNotFound):
		writeError(w, http.StatusNotFound, err.Error())
	default:
		writeError(w, http.StatusInternalServerError, "internal error")
	}
}

func writeError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
