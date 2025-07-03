-- ERP Opity - Estrutura do Banco de Dados (PostgreSQL)

-- Tabela: perfis
CREATE TABLE perfis (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(50) NOT NULL UNIQUE,
    permissoes JSONB NOT NULL DEFAULT '{}'
);

-- Tabela: usuarios
CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    login VARCHAR(50) NOT NULL UNIQUE,
    senha_hash VARCHAR(255) NOT NULL,
    perfil_id INTEGER NOT NULL REFERENCES perfis(id),
    ativo BOOLEAN NOT NULL DEFAULT TRUE
);

-- Tabela: clientes
CREATE TABLE clientes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(150) NOT NULL,
    tipo_pessoa VARCHAR(10) NOT NULL CHECK (tipo_pessoa IN ('Física', 'Jurídica')),
    cpf_cnpj VARCHAR(20) NOT NULL UNIQUE,
    telefone VARCHAR(20),
    email VARCHAR(100),
    endereco_logradouro VARCHAR(100),
    endereco_numero VARCHAR(10),
    endereco_complemento VARCHAR(50),
    endereco_bairro VARCHAR(50),
    endereco_cidade VARCHAR(50),
    endereco_uf CHAR(2),
    endereco_cep VARCHAR(10),
    criado_em TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    atualizado_em TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Tabela: categorias_produtos
CREATE TABLE categorias_produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL UNIQUE,
    descricao TEXT
);

-- Tabela: produtos
CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    codigo_barras VARCHAR(50) UNIQUE,
    nome VARCHAR(100) NOT NULL,
    descricao TEXT,
    categoria_id INTEGER REFERENCES categorias_produtos(id),
    unidade_medida VARCHAR(10) NOT NULL DEFAULT 'UN',
    estoque_atual NUMERIC(12,3) NOT NULL DEFAULT 0,
    preco_custo NUMERIC(12,2) NOT NULL DEFAULT 0,
    preco_venda NUMERIC(12,2) NOT NULL DEFAULT 0,
    ativo BOOLEAN NOT NULL DEFAULT TRUE
);

-- Tabela: vendas
CREATE TABLE vendas (
    id SERIAL PRIMARY KEY,
    cliente_id INTEGER REFERENCES clientes(id),
    usuario_id INTEGER REFERENCES usuarios(id),
    data_venda TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(20) NOT NULL CHECK (status IN ('Aberta', 'Paga', 'Cancelada')),
    total_bruto NUMERIC(12,2) NOT NULL DEFAULT 0,
    desconto NUMERIC(12,2) NOT NULL DEFAULT 0,
    total_liquido NUMERIC(12,2) NOT NULL DEFAULT 0
);

-- Tabela: itens_venda
CREATE TABLE itens_venda (
    id SERIAL PRIMARY KEY,
    venda_id INTEGER NOT NULL REFERENCES vendas(id) ON DELETE CASCADE,
    produto_id INTEGER NOT NULL REFERENCES produtos(id),
    quantidade NUMERIC(12,3) NOT NULL,
    preco_unitario NUMERIC(12,2) NOT NULL,
    subtotal NUMERIC(12,2) NOT NULL
);

-- Tabela: pagamentos
CREATE TABLE pagamentos (
    id SERIAL PRIMARY KEY,
    venda_id INTEGER NOT NULL REFERENCES vendas(id) ON DELETE CASCADE,
    forma_pagamento VARCHAR(20) NOT NULL CHECK (forma_pagamento IN ('Dinheiro', 'Cartão', 'PIX', 'Ticket')),
    valor NUMERIC(12,2) NOT NULL,
    data_pagamento TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Índices para performance
CREATE INDEX idx_produtos_categoria ON produtos(categoria_id);
CREATE INDEX idx_vendas_cliente ON vendas(cliente_id);
CREATE INDEX idx_vendas_usuario ON vendas(usuario_id);
CREATE INDEX idx_itens_venda_venda ON itens_venda(venda_id);
CREATE INDEX idx_itens_venda_produto ON itens_venda(produto_id);
CREATE INDEX idx_pagamentos_venda ON pagamentos(venda_id); 