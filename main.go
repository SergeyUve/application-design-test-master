package main

import (
	"application-design-test-master/handlers"
	l "application-design-test-master/logger"
	b "application-design-test-master/repositories/booking"
	o "application-design-test-master/services/order"
	"errors"
	"net/http"
	"os"
)

func main() {
	reposytory := b.NewReposytory()
	orderService := o.NewOrderService(reposytory)
	logger := l.NewLogger()
	orderHandler := handlers.NewOrderHandler(orderService, logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/orders", orderHandler.Create)

	logger.LogInfo("Server listening on localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		logger.LogInfo("Server closed")
	} else if err != nil {
		logger.LogErrorf("Server failed: %s", err)
		os.Exit(1)
	}
}
