# ERP Opity - Sistema ERP Modular Multiempresa

## Visão Geral

O ERP Opity é um sistema de gestão empresarial completo, modular e extensível desenvolvido em Go usando o framework Hexya. Inspirado nas melhores práticas do TOTVS, Odoo e SAP, o sistema foi projetado para atender inicialmente o varejo de supermercados, mas com arquitetura modular que permite adaptação a qualquer segmento.

## Características Principais

### 🏢 Multiempresa
- Suporte completo a múltiplas empresas
- Isolamento de dados por empresa
- Configurações independentes por organização

### 👥 Multiusuário
- Sistema de permissões granular
- Controle de acesso por módulo e função
- Auditoria completa de ações dos usuários

### 🔧 Modular
- Arquitetura baseada em módulos independentes
- Instalação/desinstalação de funcionalidades
- Extensibilidade através de plugins

### 🛒 Foco em Varejo
- Gestão de estoque em tempo real
- Controle de caixa e PDV
- Gestão de fornecedores e compras
- Relatórios específicos do varejo

## Módulos Disponíveis

### Core (Base)
- Gestão de usuários e permissões
- Configurações do sistema
- Multiempresa e multilíngue

### Retail (Varejo)
- Gestão de produtos e categorias
- Controle de estoque
- Gestão de caixa e PDV
- Relatórios de vendas

### Stock (Estoque)
- Controle de inventário
- Movimentações de estoque
- Gestão de armazéns
- Rastreabilidade de produtos

### Sale (Vendas)
- Gestão de clientes
- Pedidos de venda
- Faturas e recebimentos
- Comissões de vendedores

### Purchase (Compras)
- Gestão de fornecedores
- Pedidos de compra
- Recebimento de mercadorias
- Controle de pagamentos

### Account (Financeiro)
- Plano de contas
- Lançamentos contábeis
- Relatórios financeiros
- Conciliação bancária

### HR (Recursos Humanos)
- Gestão de funcionários
- Controle de ponto
- Folha de pagamento
- Benefícios

## Arquitetura

```
erp-opity/
├── cmd/                    # Executáveis principais
├── internal/              # Código interno da aplicação
│   ├── core/             # Módulo base
│   ├── retail/           # Módulo varejo
│   ├── stock/            # Módulo estoque
│   ├── sale/             # Módulo vendas
│   ├── purchase/         # Módulo compras
│   ├── account/          # Módulo financeiro
│   ├── hr/               # Módulo RH
│   └── shared/           # Código compartilhado
├── pkg/                  # Bibliotecas públicas
├── configs/              # Configurações
├── migrations/           # Migrações de banco
├── docs/                 # Documentação
└── tests/                # Testes
```

## Tecnologias

- **Linguagem**: Go 1.21+
- **Framework**: Hexya
- **Banco de Dados**: PostgreSQL
- **Interface**: Desktop (nativo) + API REST
- **Arquitetura**: Modular, Event-driven

## Instalação

### Pré-requisitos
- Go 1.21 ou superior
- PostgreSQL 13+
- Git

### Passos

1. Clone o repositório:
```bash
git clone https://github.com/seu-usuario/erp-opity.git
cd erp-opity
```

2. Instale as dependências:
```bash
go mod download
```

3. Configure o banco de dados (o arquivo `configs/config.yaml` contém credenciais de exemplo, altere-as antes de usar em produção):
```bash
# Edite configs/config.yaml com suas configurações
```

4. Execute as migrações:
```bash
go run cmd/migrate/main.go
```

5. Inicie o servidor:
```bash
go run cmd/server/main.go
```

6. Inicie a aplicação desktop:
```bash
go run cmd/desktop/main.go
```

7. Inicie a interface web (opcional):
```bash
cd ui
npm install
npm run dev
```

## Configuração

O sistema utiliza arquivos YAML para configuração. Principais seções:

- **Database**: Configurações do PostgreSQL
- **Server**: Configurações do servidor HTTP
- **Modules**: Módulos ativos
- **Security**: Configurações de segurança
- **Logging**: Configurações de log

## Desenvolvimento

### Estrutura de um Módulo

```go
package retail

import (
    "github.com/hexya-erp/hexya"
)

func init() {
    hexya.RegisterModule(&hexya.Module{
        Name:        "retail",
        Description: "Módulo de Varejo",
        Version:     "1.0.0",
        Dependencies: []string{"base", "stock"},
        Models:      []interface{}{
            &Product{},
            &Category{},
            &POS{},
        },
    })
}
```

### Criando um Novo Modelo

```go
type Product struct {
    hexya.Model
    Name        string  `json:"name"`
    Code        string  `json:"code"`
    Price       float64 `json:"price"`
    CategoryID  int64   `json:"category_id"`
    Category    *Category `json:"category"`
}

func (p *Product) TableName() string {
    return "retail_products"
}
```

## API REST

O sistema expõe uma API REST completa para integração:

- **Autenticação**: JWT
- **Formato**: JSON
- **Versionamento**: v1
- **Documentação**: Swagger/OpenAPI

### Exemplos de Endpoints

```
GET    /api/v1/products
POST   /api/v1/products
GET    /api/v1/products/{id}
PUT    /api/v1/products/{id}
DELETE /api/v1/products/{id}
```

## Testes

```bash
# Executar todos os testes
go test ./...

# Executar testes com cobertura
go test -cover ./...

# Executar testes de um módulo específico
go test ./internal/retail/...
```

## Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature
3. Commit suas mudanças
4. Push para a branch
5. Abra um Pull Request

## Licença

Este projeto está licenciado sob a MIT License - veja o arquivo [LICENSE](LICENSE) para detalhes.

## Suporte

- **Documentação**: [docs/](docs/)
- **Issues**: [GitHub Issues](https://github.com/seu-usuario/erp-opity/issues)
- **Email**: suporte@erp-opity.com

## Roadmap

### Versão 1.0 (Atual)
- ✅ Módulos core implementados
- ✅ Sistema multiempresa
- ✅ API REST básica
- ✅ Interface desktop

### Versão 1.1 (Próxima)
- 🔄 Interface web
- 🔄 App mobile
- 🔄 Integração com marketplaces
- 🔄 Business Intelligence

### Versão 2.0 (Futuro)
- 📋 IA para previsão de demanda
- 📋 Blockchain para rastreabilidade
- 📋 IoT para automação
- 📋 Cloud nativo 