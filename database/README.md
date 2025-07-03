# Banco de Dados ERP Opity

Este diretório contém todos os scripts SQL para configuração e manutenção do banco de dados PostgreSQL do ERP Opity.

## Arquivos

- `schema.sql` - Estrutura das tabelas do banco
- `seed_data.sql` - Dados iniciais (perfis, usuários, categorias, produtos, clientes)
- `setup_user.sql` - Configuração de usuário e privilégios
- `create_database.bat` - Script automatizado para Windows

## Configuração Manual

### 1. Criar Usuário e Banco

```bash
# Conectar como postgres
psql -U postgres

# Criar usuário
CREATE USER erp_user WITH PASSWORD 'erp_password_2024';

# Criar banco
CREATE DATABASE erp_opity OWNER erp_user;

# Conceder privilégios
GRANT ALL PRIVILEGES ON DATABASE erp_opity TO erp_user;

# Sair
\q
```

### 2. Criar Tabelas

```bash
# Conectar ao banco erp_opity com usuário erp_user
psql -U erp_user -d erp_opity -f database/schema.sql
```

### 3. Inserir Dados Iniciais

```bash
# Inserir dados de exemplo
psql -U erp_user -d erp_opity -f database/seed_data.sql
```

## Configuração Automatizada (Windows)

Execute o arquivo `create_database.bat` que irá:

1. Criar o usuário `erp_user`
2. Criar o banco `erp_opity`
3. Conceder privilégios
4. Criar todas as tabelas
5. Inserir dados iniciais

## Credenciais

- **Usuário**: `erp_user`
- **Senha**: `erp_password_2024`
- **Banco**: `erp_opity`
- **Host**: `localhost`
- **Porta**: `5432`

## Estrutura das Tabelas

### Tabelas Principais

1. **perfis** - Perfis de usuário e permissões
2. **usuarios** - Usuários do sistema
3. **clientes** - Cadastro de clientes
4. **categorias_produtos** - Categorias de produtos
5. **produtos** - Cadastro de produtos
6. **vendas** - Cabeçalho das vendas
7. **itens_venda** - Itens de cada venda
8. **pagamentos** - Formas de pagamento das vendas

### Relacionamentos

- `usuarios.perfil_id` → `perfis.id`
- `produtos.categoria_id` → `categorias_produtos.id`
- `vendas.cliente_id` → `clientes.id`
- `vendas.usuario_id` → `usuarios.id`
- `itens_venda.venda_id` → `vendas.id`
- `itens_venda.produto_id` → `produtos.id`
- `pagamentos.venda_id` → `vendas.id`

## Dados Iniciais

### Perfis Criados
- **Administrador** - Acesso total ao sistema
- **Gerente** - Acesso a clientes, produtos e vendas
- **Vendedor** - Acesso limitado a vendas
- **Fiscal** - Acesso apenas a relatórios

### Usuário Padrão
- **Login**: `admin`
- **Senha**: `admin123`
- **Perfil**: Administrador

### Categorias de Produtos
- Alimentos
- Bebidas
- Limpeza
- Higiene
- Eletrônicos
- Vestuário

### Produtos de Exemplo
- Arroz Integral 1kg
- Feijão Preto 1kg
- Coca-Cola 2L
- Detergente Líquido 500ml
- Sabonete Líquido 300ml

### Clientes de Exemplo
- João Silva (Pessoa Física)
- Maria Santos (Pessoa Física)
- Empresa ABC Ltda (Pessoa Jurídica)

## Backup e Restore

### Backup
```bash
pg_dump -U erp_user -d erp_opity > backup_erp_opity.sql
```

### Restore
```bash
psql -U erp_user -d erp_opity < backup_erp_opity.sql
```

## Manutenção

### Verificar Conexão
```bash
psql -U erp_user -d erp_opity -c "SELECT version();"
```

### Listar Tabelas
```bash
psql -U erp_user -d erp_opity -c "\dt"
```

### Verificar Tamanho do Banco
```bash
psql -U erp_user -d erp_opity -c "SELECT pg_size_pretty(pg_database_size('erp_opity'));"
``` 