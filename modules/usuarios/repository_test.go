package usuarios

import (
    "regexp"
    "testing"

    "github.com/DATA-DOG/go-sqlmock"
)

func TestUsuarioRepository_ListarUsuarios(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("error creating mock db: %v", err)
    }
    defer db.Close()

    rows := sqlmock.NewRows([]string{"id", "nome", "login", "perfil_id", "ativo"}).
        AddRow(1, "Admin", "admin", 1, true)

    mock.ExpectQuery(regexp.QuoteMeta("SELECT id, nome, login, perfil_id, ativo FROM usuarios")).WillReturnRows(rows)

    repo := NewUsuarioRepository(db)
    res, err := repo.ListarUsuarios()
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if len(res) != 1 || res[0].Login != "admin" {
        t.Fatalf("unexpected result: %+v", res)
    }
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Fatalf("expectations: %v", err)
    }
}
