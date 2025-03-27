package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/GuiCezaF/task-collector-v2/internal/utils"
)

func handler(w http.ResponseWriter, r *http.Request) {
	response := utils.Response{Message: "Server is running", Status: http.StatusOK}
	utils.JSONResponse(w, response, http.StatusOK)
}

func main() {
	port := ":8080"
	var wg sync.WaitGroup

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Printf("🚀 Servidor rodando em http://localhost%s \n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Erro ao iniciar servidor: %v\n", err)
		}
	}()

	<-stop
	fmt.Println("\n🛑 Desligando servidor...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Erro ao desligar servidor: %v\n", err)
	}

	wg.Wait()
	fmt.Println("✅ Servidor encerrado com sucesso!")
}
