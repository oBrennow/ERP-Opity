package main

import (
	"erp-opity/internal/account"
	"erp-opity/internal/core"
	"erp-opity/internal/hr"
	"erp-opity/internal/purchase"
	"erp-opity/internal/retail"
	"erp-opity/internal/sale"
	"erp-opity/internal/stock"
	"erp-opity/pkg/config"
	"erp-opity/pkg/database"
	"erp-opity/pkg/logger"
	"os"
)

func main() {
	// Inicializar logger
	logger := logger.New()
	logger.Info("Iniciando migrações do ERP Opity...")

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

	// Inicializar módulos
	modules := []core.Module{
		core.New(),
		retail.New(),
		stock.New(),
		sale.New(),
		purchase.New(),
		account.New(),
		hr.New(),
	}

	// Registrar módulos
	for _, module := range modules {
		if err := module.Register(); err != nil {
			logger.Fatal("Erro ao registrar módulo:", err)
		}
	}

	// Executar migrações
	logger.Info("Executando migrações...")
	if err := database.Migrate(db, modules); err != nil {
		logger.Fatal("Erro ao executar migrações:", err)
	}

	logger.Info("Migrações executadas com sucesso!")

	// Se solicitado, criar dados iniciais
	if len(os.Args) > 1 && os.Args[1] == "--seed" {
		logger.Info("Criando dados iniciais...")
		if err := database.Seed(db, modules); err != nil {
			logger.Fatal("Erro ao criar dados iniciais:", err)
		}
		logger.Info("Dados iniciais criados com sucesso!")
	}
}
