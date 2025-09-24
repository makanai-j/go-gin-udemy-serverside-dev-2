package controllers

import (
	"log"
	"net/http"
	"time"
)

func NewMux(th *TradeHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/trades", methodMux(map[string]http.HandlerFunc{
		http.MethodPost: th.Create,
	}))
	mux.HandleFunc("/trades/", methodMux(map[string]http.HandlerFunc{
		http.MethodGet: th.GetByID,
	}))
	return logging(mux)
}

// メソッド分岐の薄いヘルパ
func methodMux(handlers map[string]http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if h, ok := handlers[r.Method]; ok {
			h(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// シンプルなロギングミドルウェア
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}
