package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"erp-opity/modules/clientes"
	"erp-opity/modules/usuarios"
	"erp-opity/pkg/config"
	"erp-opity/pkg/database"
	"erp-opity/pkg/logger"
	"erp-opity/pkg/server"

	"github.com/gorilla/mux"
)

func main() {
	// Inicializar logger
	logger := logger.New()
	logger.Info("Iniciando ERP Opity...")

	// Carregar configurações
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		logger.Fatal("Erro ao carregar configurações:", err)
	}

	// Conectar ao banco de dados
	db, err := database.Connect(cfg.Database)
	if err != nil {
		logger.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	// Configurar roteador
	router := mux.NewRouter()

	// Middleware
        router.Use(server.LoggingMiddleware(logger))
        router.Use(server.CORSMiddleware)
        router.Use(server.AuthMiddleware(cfg.Security.JWTSecret))
        router.Use(server.RateLimitMiddleware(100, time.Minute))

	// Health check
	router.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","timestamp":"` + time.Now().Format(time.RFC3339) + `"}`))
	}).Methods("GET")

	// --- Integração do módulo de usuários ---
	usuarioRepo := usuarios.NewUsuarioRepository(db.DB)
	usuarioHandler := usuarios.NewUsuarioHandler(usuarioRepo)
	router.HandleFunc("/api/v1/usuarios", usuarioHandler.ListarUsuarios).Methods("GET")
	// --- Fim integração usuários ---

	// --- Integração do módulo de clientes ---
	clienteRepo := clientes.NewClienteRepository(db.DB)
	clienteHandler := clientes.NewClienteHandler(clienteRepo)
	router.HandleFunc("/api/v1/clientes", clienteHandler.Criar).Methods("POST")
	router.HandleFunc("/api/v1/clientes/{id}", clienteHandler.Editar).Methods("PUT")
	router.HandleFunc("/api/v1/clientes", clienteHandler.Listar).Methods("GET")
	router.HandleFunc("/api/v1/clientes/search", clienteHandler.Pesquisar).Methods("GET")
	// --- Fim integração clientes ---

	// Configurar servidor HTTP
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Canal para receber sinais de interrupção
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Iniciar servidor em goroutine
	go func() {
		logger.Infof("Servidor iniciado na porta %d", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Erro ao iniciar servidor:", err)
		}
	}()

	// Aguardar sinal de interrupção
	<-done
	logger.Info("Recebido sinal de interrupção, encerrando servidor...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Erro durante shutdown:", err)
	}

	logger.Info("Servidor encerrado com sucesso")
}
