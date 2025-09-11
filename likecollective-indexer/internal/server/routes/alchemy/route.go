package alchemy

import (
	"net/http"
)

type alchemyHandler struct {
	webhookSigningKey string
}

func (h *alchemyHandler) NewServer() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/webhook", NewAlchemyRequestHandlerMiddleware(h.handleWebhook, h.webhookSigningKey))
	return mux
}

func NewAlchemyHandler(webhookSigningKey string) http.Handler {
	h := &alchemyHandler{
		webhookSigningKey,
	}
	return h.NewServer()
}
