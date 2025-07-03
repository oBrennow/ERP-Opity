@echo off
echo ========================================
echo ERP Opity - Setup do Banco de Dados
echo ========================================

echo.
echo 1. Criando usuario erp_user...
psql -U postgres -c "CREATE USER erp_user WITH PASSWORD 'erp_password_2024';"

echo.
echo 2. Criando banco de dados erp_opity...
psql -U postgres -c "CREATE DATABASE erp_opity OWNER erp_user;"

echo.
echo 3. Concedendo privilegios...
psql -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE erp_opity TO erp_user;"

echo.
echo 4. Criando tabelas no banco erp_opity...
psql -U erp_user -d erp_opity -f database/schema.sql

echo.
echo ========================================
echo Setup concluido!
echo ========================================
echo.
echo Credenciais do banco:
echo Usuario: erp_user
echo Senha: erp_password_2024
echo Banco: erp_opity
echo.
pause 