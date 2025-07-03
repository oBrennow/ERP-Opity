-- ERP Opity - Dados Iniciais (Seeds)

-- Inserir perfis padrão
INSERT INTO perfis (nome, permissoes) VALUES
('Administrador', '{"todos": ["read", "write", "delete"]}'),
('Gerente', '{"clientes": ["read", "write"], "produtos": ["read", "write"], "vendas": ["read", "write"]}'),
('Vendedor', '{"clientes": ["read"], "produtos": ["read"], "vendas": ["read", "write"]}'),
('Fiscal', '{"vendas": ["read"], "relatorios": ["read"]}');

-- Inserir usuário administrador padrão
-- Senha: admin123 (hash bcrypt)
INSERT INTO usuarios (nome, login, senha_hash, perfil_id, ativo) VALUES
('Administrador Sistema', 'admin', '$2a$12$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewdBPj4J/HS.iK8.', 1, true);

-- Inserir categorias de produtos padrão
INSERT INTO categorias_produtos (nome, descricao) VALUES
('Alimentos', 'Produtos alimentícios em geral'),
('Bebidas', 'Bebidas e refrigerantes'),
('Limpeza', 'Produtos de limpeza'),
('Higiene', 'Produtos de higiene pessoal'),
('Eletrônicos', 'Produtos eletrônicos'),
('Vestuário', 'Roupas e acessórios');

-- Inserir produtos de exemplo
INSERT INTO produtos (codigo_barras, nome, descricao, categoria_id, unidade_medida, estoque_atual, preco_custo, preco_venda, ativo) VALUES
('7891234567890', 'Arroz Integral 1kg', 'Arroz integral tipo 1, pacote de 1kg', 1, 'UN', 50.000, 4.50, 6.90, true),
('7891234567891', 'Feijão Preto 1kg', 'Feijão preto tipo 1, pacote de 1kg', 1, 'UN', 30.000, 3.20, 5.50, true),
('7891234567892', 'Coca-Cola 2L', 'Refrigerante Coca-Cola, garrafa de 2L', 2, 'UN', 20.000, 3.80, 6.50, true),
('7891234567893', 'Detergente Líquido 500ml', 'Detergente líquido para louças, frasco de 500ml', 3, 'UN', 25.000, 2.50, 4.20, true),
('7891234567894', 'Sabonete Líquido 300ml', 'Sabonete líquido para mãos, frasco de 300ml', 4, 'UN', 15.000, 3.50, 5.80, true);

-- Inserir clientes de exemplo
INSERT INTO clientes (nome, tipo_pessoa, cpf_cnpj, telefone, email, endereco_logradouro, endereco_numero, endereco_bairro, endereco_cidade, endereco_uf, endereco_cep) VALUES
('João Silva', 'Física', '123.456.789-00', '(11) 99999-9999', 'joao@email.com', 'Rua das Flores', '123', 'Centro', 'São Paulo', 'SP', '01234-567'),
('Maria Santos', 'Física', '987.654.321-00', '(11) 88888-8888', 'maria@email.com', 'Av. Paulista', '456', 'Bela Vista', 'São Paulo', 'SP', '01310-100'),
('Empresa ABC Ltda', 'Jurídica', '12.345.678/0001-90', '(11) 77777-7777', 'contato@empresaabc.com', 'Rua Augusta', '789', 'Consolação', 'São Paulo', 'SP', '01212-001'); 