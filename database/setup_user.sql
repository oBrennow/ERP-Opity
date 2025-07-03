-- ERP Opity - Setup de Usuário e Banco

-- Criar usuário específico para o ERP
CREATE USER erp_user WITH PASSWORD 'erp_password_2024';

-- Criar banco de dados
CREATE DATABASE erp_opity OWNER erp_user;

-- Conceder privilégios ao usuário
GRANT ALL PRIVILEGES ON DATABASE erp_opity TO erp_user;

-- Conectar ao banco erp_opity
\c erp_opity;

-- Conceder privilégios de schema
GRANT ALL ON SCHEMA public TO erp_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO erp_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO erp_user;
GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA public TO erp_user;

-- Configurar para futuras tabelas
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO erp_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON SEQUENCES TO erp_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON FUNCTIONS TO erp_user; 