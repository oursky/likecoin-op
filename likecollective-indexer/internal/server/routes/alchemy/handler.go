package alchemy

import (
	"fmt"
	"net/http"
)

func (h *alchemyHandler) handleWebhook(w http.ResponseWriter, r *http.Request, event *AlchemyWebhookEvent) {
	fmt.Println(event)
}
