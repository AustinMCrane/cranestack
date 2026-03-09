package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AustinMCrane/cranekit/auth"
	"github.com/AustinMCrane/cranestack/internal/db"
)

// subscriptionResponse is the JSON shape returned by subscription endpoints.
type subscriptionResponse struct {
	Status    string `json:"status"`
	ProductID string `json:"product_id"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		// Header already sent; log only.
		_ = err
	}
}

// GetSubscription handles GET /api/subscriptions.
// It returns the authenticated user's current subscription status.
func (h *Handlers) GetSubscription(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.UserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	sub, err := h.repo.GetSubscriptionByUserID(userID)
	if err != nil {
		http.Error(w, "failed to get subscription", http.StatusInternalServerError)
		return
	}

	if sub == nil {
		writeJSON(w, http.StatusOK, subscriptionResponse{Status: "none", ProductID: ""})
		return
	}

	writeJSON(w, http.StatusOK, subscriptionResponse{
		Status:    sub.Status,
		ProductID: sub.ProductID,
	})
}

// CreateSubscription handles POST /api/subscriptions.
// It activates a subscription for the authenticated user.
//
// NOTE: This handler records the subscription directly without payment
// verification. In production, integrate a payment gateway (e.g. Stripe,
// RevenueCat, or StoreKit server-side validation) and only set status =
// "active" after confirming payment.
func (h *Handlers) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.UserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req struct {
		ProductID string `json:"product_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if req.ProductID == "" {
		req.ProductID = "monthly"
	}

	sub := &db.Subscription{
		UserID:    userID,
		ProductID: req.ProductID,
		Status:    "active",
	}
	if err := h.repo.CreateSubscription(sub); err != nil {
		http.Error(w, "failed to create subscription", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, subscriptionResponse{
		Status:    sub.Status,
		ProductID: sub.ProductID,
	})
}
