package clientes

import (
    "regexp"
    "testing"
    "time"

    "github.com/DATA-DOG/go-sqlmock"
)

func TestClienteRepository_Listar(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("error creating mock db: %v", err)
    }
    defer db.Close()

    rows := sqlmock.NewRows([]string{"id", "nome", "tipo_pessoa", "cpf_cnpj", "telefone", "email", "endereco_logradouro", "endereco_numero", "endereco_complemento", "endereco_bairro", "endereco_cidade", "endereco_uf", "endereco_cep", "criado_em", "atualizado_em"}).
        AddRow(1, "Teste", "Física", "123", "", "", "Rua", "10", "", "Centro", "SP", "SP", "00000-000", time.Now(), time.Now())

    mock.ExpectQuery(regexp.QuoteMeta("SELECT id, nome, tipo_pessoa, cpf_cnpj, telefone, email, endereco_logradouro, endereco_numero, endereco_complemento, endereco_bairro, endereco_cidade, endereco_uf, endereco_cep, criado_em, atualizado_em FROM clientes ORDER BY id DESC")).WillReturnRows(rows)

    repo := NewClienteRepository(db)
    res, err := repo.Listar()
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if len(res) != 1 || res[0].Nome != "Teste" {
        t.Fatalf("unexpected result: %+v", res)
    }
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Fatalf("expectations: %v", err)
    }
}

func TestClienteRepository_BuscarPorNome(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("error creating mock db: %v", err)
    }
    defer db.Close()

    rows := sqlmock.NewRows([]string{"id", "nome", "tipo_pessoa", "cpf_cnpj", "telefone", "email", "endereco_logradouro", "endereco_numero", "endereco_complemento", "endereco_bairro", "endereco_cidade", "endereco_uf", "endereco_cep", "criado_em", "atualizado_em"}).
        AddRow(2, "Fulano", "Física", "456", "", "", "Rua", "20", "", "Centro", "SP", "SP", "11111-111", time.Now(), time.Now())

    mock.ExpectQuery(regexp.QuoteMeta("SELECT id, nome, tipo_pessoa, cpf_cnpj, telefone, email, endereco_logradouro, endereco_numero, endereco_complemento, endereco_bairro, endereco_cidade, endereco_uf, endereco_cep, criado_em, atualizado_em FROM clientes WHERE nome ILIKE $1 ORDER BY nome")).
        WithArgs("%ful%").
        WillReturnRows(rows)

    repo := NewClienteRepository(db)
    res, err := repo.BuscarPorNome("ful")
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if len(res) != 1 || res[0].ID != 2 {
        t.Fatalf("unexpected result: %+v", res)
    }
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Fatalf("expectations: %v", err)
    }
}
