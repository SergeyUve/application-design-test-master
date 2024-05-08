package handlers

import (
	l "application-design-test-master/logger"
	"application-design-test-master/models"
	o "application-design-test-master/services/order"
	"encoding/json"
	"net/http"
)

type OrderHandler struct {
	OrderService o.OrderService
	Logger       l.Logger
}

func NewOrderHandler(
	orderService o.OrderService,
	logger l.Logger) *OrderHandler {
	return &OrderHandler{
		OrderService: orderService,
		Logger:       logger,
	}
}

func (app *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newOrder models.Order

	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		app.Logger.LogErrorf("json.NewDecoder(): %v", err)

		return
	}

	err = app.OrderService.CreateOrder(newOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		app.Logger.LogErrorf("CreateOrder(): %v", err)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)

	app.Logger.LogInfof("Order successfully created: %v", newOrder)
}
