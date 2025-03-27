package routes

import (
	"net/http"

	"github.com/GuiCezaF/task-collector-v2/internal/handlers"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/status", handlers.HandlerStatus)

	mux.Handle("/api/v2/", http.StripPrefix("/api/v2", apiMux))

	return mux
}
