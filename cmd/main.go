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

	"github.com/GuiCezaF/task-collector-v2/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
	}

	port := os.Getenv("PORT")
	var wg sync.WaitGroup

	mux := routes.RegisterRoutes()

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("🚀 Servidor rodando em http://localhost%s/api/v2/ \n", port)
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
